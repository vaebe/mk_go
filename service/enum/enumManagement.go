package enum

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mk/global"
	"mk/models/enum"
	"mk/utils"
)

// Save
// @Summary     增加、编辑
// @Description  增加、编辑
// @Tags         enum枚举
// @Accept       json
// @Produce      json
// @Param 				param body    enum.EnumsForm true  "请求对象"
// @Success      200  {object}  utils.ResponseResultInfo
// @Failure      500  {object}  utils.EmptyInfo
// @Router       /enum/save [post]
func Save(ctx *gin.Context) {
	enumForm := enum.EnumsForm{}
	if err := ctx.ShouldBind(&enumForm); err != nil {
		zap.S().Info("枚举保存信息:", &enumForm)
		utils.HandleValidatorError(ctx, err)
		return
	}

	res := global.DB.Create(&enum.Enum{
		Value:    enumForm.Value,
		Name:     enumForm.Name,
		TypeCode: enumForm.TypeCode,
		TypeName: enumForm.TypeName,
	})

	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, enumForm)
}

// Delete
// @Summary      根据id删除指定枚举
// @Description  根据id删除指定枚举
// @Tags         enum枚举
// @Accept       json
// @Produce      json
// @Param        id   query      int  true  "枚举id"
// @Success      200  {object}  utils.ResponseResultInfo
// @Failure      500  {object}  utils.EmptyInfo
// @Router       /enum/delete [delete]
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
// @Summary     获取枚举详情
// @Description  获取枚举详情
// @Tags         enum枚举
// @Accept       json
// @Produce      json
// @Param        id   query      int  true  "枚举id"
// @Success      200  {object}  utils.ResponseResultInfo
// @Failure      500  {object}  utils.EmptyInfo
// @Router       /enum/getEnumDetails [get]
func Details(ctx *gin.Context) {

}

// GetEnumsByType 根据分类查询 单个 多个
func GetEnumsByType(ctx *gin.Context) {

}

// GetAllEnums
// @Summary     获取全部数据
// @Description  获取全部数据
// @Tags         enum枚举
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.ResponseResultInfo
// @Failure      500  {object}  utils.EmptyInfo
// @Router       /enum/getAllEnums [get]
func GetAllEnums(ctx *gin.Context) {

}
