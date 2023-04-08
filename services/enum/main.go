package enum

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mk/global"
	"mk/models"
	"mk/models/enum"
	"mk/utils"
)

// Save
//
//	@Summary		增加、编辑
//	@Description	增加、编辑
//	@Tags			enum枚举
//	@Accept			json
//	@Produce		json
//	@Param			param	body		enum.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/enum/save [post]
func Save(ctx *gin.Context) {
	saveForm := enum.SaveForm{}
	if err := ctx.ShouldBind(&saveForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	saveInfo := enum.Enum{
		Value:    saveForm.Value,
		Name:     saveForm.Name,
		TypeCode: saveForm.TypeCode,
		TypeName: saveForm.TypeName,
	}

	// id 不存在新增
	if saveForm.ID == 0 {
		global.DB.Create(&saveInfo)
		utils.ResponseResultsSuccess(ctx, map[string]any{"id": saveInfo.ID})
	} else {
		res := global.DB.Model(&enum.Enum{}).Where("id = ?", saveForm.ID).Updates(&saveInfo)
		if res.Error != nil {
			utils.ResponseResultsError(ctx, res.Error.Error())
			return
		}

		utils.ResponseResultsSuccess(ctx, map[string]any{"id": saveInfo.ID})
	}
}

// Delete todo 考虑增加类型 如系统则不能被删除
//
//	@Summary		根据id删除指定枚举
//	@Description	根据id删除指定枚举
//	@Tags			enum枚举
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"枚举id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/enum/delete [delete]
func Delete(ctx *gin.Context) {
	enumsId := ctx.Query("id")

	if enumsId == "" {
		utils.ResponseResultsError(ctx, "枚举id不能为空！")
		return
	}

	res := global.DB.Where("id = ?", enumsId).Delete(&enum.Enum{})

	total := res.RowsAffected
	if total == 0 {
		utils.ResponseResultsError(ctx, "需要删除的数据不存在！")
		return
	}

	utils.ResponseResultsSuccess(ctx, "删除成功！")
}

// Details
//
//	@Summary		获取枚举详情
//	@Description	获取枚举详情
//	@Tags			enum枚举
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"枚举id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/enum/details [get]
func Details(ctx *gin.Context) {
	enumsId := ctx.Query("id")

	if enumsId == "" {
		utils.ResponseResultsError(ctx, "枚举id不能为空！")
		return
	}

	enumInfo := enum.SaveForm{}
	res := global.DB.Model(&enum.Enum{}).Where("id = ?", enumsId).First(&enumInfo)

	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, enumInfo)
}

// GetEnumsByType
//
//	@Summary		根据分类查询枚举
//	@Description	根据分类查询枚举
//	@Tags			enum枚举
//	@Accept			json
//	@Produce		json
//	@Param			type	query		string	true	"枚举类型code"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/enum/getEnumsByType [get]
func GetEnumsByType(ctx *gin.Context) {
	typeCode := ctx.Query("type")

	if typeCode == "" {
		utils.ResponseResultsError(ctx, "枚举类型code不能为空！")
		return
	}

	var enumsList []enum.SaveForm
	res := global.DB.Model(&enum.Enum{}).Where("type_code", typeCode).Find(&enumsList)

	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, enumsList)
}

// GetAllEnums
//
//	@Summary		获取全部数据
//	@Description	获取全部数据
//	@Tags			enum枚举
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/enum/getAllEnums [get]
func GetAllEnums(ctx *gin.Context) {
	var enumsList []enum.SaveForm
	res := global.DB.Model(&enum.Enum{}).Find(&enumsList)
	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	enums := make(map[string][]enum.SaveForm)

	for _, v := range enumsList {
		enums[v.TypeCode] = append(enums[v.TypeCode], v)
	}

	utils.ResponseResultsSuccess(ctx, enums)
}

// GetEnumsList
//
//	@Summary			分页获取枚举列表
//	@Description	分页获取枚举列表
//	@Tags				enum枚举
//	@Accept			json
//	@Produce		json
//	@Param			param	body		enum.ListForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Router			/enum/getEnumsList [post]
func GetEnumsList(ctx *gin.Context) {
	listForm := enum.ListForm{}

	if err := ctx.ShouldBind(&listForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	var enums []enum.Enum
	res := global.DB.Where("name LIKE ? AND type_name LIKE ?", "%"+listForm.Name+"%", "%"+listForm.TypeName+"%").Find(&enums)

	// 存在错误
	if res.Error != nil {
		zap.S().Info(res.Error)
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	// 获取总数
	total := int32(res.RowsAffected)

	// 分页
	res.Scopes(utils.Paginate(listForm.PageNo, listForm.PageSize)).Find(&enums)

	utils.ResponseResultsSuccess(ctx, &models.PagingData{
		PageSize: listForm.PageSize,
		PageNo:   listForm.PageNo,
		Total:    total,
		List:     enums,
	})
}
