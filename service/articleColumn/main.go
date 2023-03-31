package articleColumn

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mk/global"
	"mk/models"
	"mk/models/articleColumn"
	"mk/utils"
)

// Save
//
//	@Summary		保存专栏
//	@Description	保存专栏
//	@Tags			articleColumn专栏
//	@Accept			json
//	@Produce		json
//	@Param			param	body		articleColumn.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/articleColumn/save [post]
func Save(ctx *gin.Context) {
	saveForm := articleColumn.SaveForm{}

	if err := ctx.ShouldBind(&saveForm); err != nil {
		zap.S().Info("专栏保存信息:", &saveForm)
		utils.HandleValidatorError(ctx, err)
		return
	}

	savaInfo := articleColumn.ArticleColumn{
		Name:         saveForm.Name,
		Introduction: saveForm.Introduction,
		CoverImg:     saveForm.CoverImg,
		Status:       "1", // 新建专栏需要审核
	}

	if savaInfo.CoverImg == "" {
		savaInfo.CoverImg = "http://rrajr4lp6.bkt.clouddn.com/mk/default/default_article_column.jpg"
	}

	// id 不存在新增
	if saveForm.ID == 0 {
		global.DB.Create(&savaInfo)
		utils.ResponseResultsSuccess(ctx, map[string]any{"id": savaInfo.ID})
	} else {
		userId, _ := ctx.Get("userId")
		res := global.DB.Model(&articleColumn.ArticleColumn{}).Where("id = ? AND user_id = ?", saveForm.ID, userId).Updates(&savaInfo)
		if res.Error != nil {
			utils.ResponseResultsError(ctx, res.Error.Error())
			return
		}

		utils.ResponseResultsSuccess(ctx, map[string]any{"id": saveForm.ID})
	}
}

// Delete
//
//	@Summary		根据id删除专栏
//	@Description	根据id删除专栏
//	@Tags			articleColumn专栏
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"专栏id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/articleColumn/delete [delete]
func Delete(ctx *gin.Context) {
	columnId := ctx.Query("id")

	if columnId == "" {
		utils.ResponseResultsError(ctx, "专栏id不能为空！")
		return
	}

	userId, _ := ctx.Get("userId")
	res := global.DB.Where("id = ? AND user_id = ?", columnId, userId).Delete(&articleColumn.ArticleColumn{})

	total := res.RowsAffected
	if total == 0 {
		utils.ResponseResultsError(ctx, "需要删除的数据不存在！")
		return
	}

	utils.ResponseResultsSuccess(ctx, "删除成功！")
}

// Details
//
//	@Summary		获取专栏详情
//	@Description	获取专栏详情
//	@Tags			articleColumn专栏
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"专栏id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/articleColumn/details [get]
func Details(ctx *gin.Context) {
	enumsId := ctx.Query("id")

	if enumsId == "" {
		utils.ResponseResultsError(ctx, "专栏id不能为空！")
		return
	}

	columnInfo := articleColumn.ArticleColumn{}
	res := global.DB.Model(&articleColumn.ArticleColumn{}).Where("id = ?", enumsId).First(&columnInfo)

	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, columnInfo)
}

// List
//
//	@Summary		获取专栏列表
//	@Description	获取专栏列表
//	@Tags			articleColumn专栏
//	@Accept			json
//	@Produce		json
//	@Param			param	body		articleColumn.ListForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/articleColumn/getList [post]
func List(ctx *gin.Context) {
	listForm := articleColumn.ListForm{}

	if err := ctx.ShouldBind(&listForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	userId, _ := ctx.Get("userId")
	var articleColumnList []articleColumn.ArticleColumn
	res := global.DB.Where("user_id = ? ", userId)

	if listForm.Name != "" {
		res.Where("name LIKE ?", "%"+listForm.Name+"%")
	}

	if listForm.Status != "" {
		res.Where("status = ?", listForm.Status)
	}

	res.Find(&articleColumnList)

	// 存在错误
	if res.Error != nil {
		zap.S().Info(res.Error)
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	// 获取总数
	total := int32(res.RowsAffected)

	// 分页
	res.Scopes(utils.Paginate(listForm.PageNo, listForm.PageSize)).Find(&articleColumnList)

	utils.ResponseResultsSuccess(ctx, &models.PagingData{
		PageSize: listForm.PageSize,
		PageNo:   listForm.PageNo,
		Total:    total,
		List:     articleColumnList,
	})
}

// Review
//
//	@Summary		文章专栏审核
//	@Description	文章专栏审核
//	@Tags			articleColumn专栏
//	@Accept			json
//	@Produce		json
//	@Param			param	body		articleColumn.ReviewForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/articleColumn/review [post]
func Review(ctx *gin.Context) {
	reviewForm := articleColumn.ReviewForm{}
	if err := ctx.ShouldBind(&reviewForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	authorityId, _ := ctx.Get("authorityId")
	if authorityId == 1 {
		utils.ResponseResultsError(ctx, "您没有审核权限！")
		return
	}

	res := global.DB.Where("id = ?", reviewForm.ID).Updates(articleColumn.ArticleColumn{
		Status: reviewForm.Status,
	})

	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
	} else {
		utils.ResponseResultsSuccess(ctx, "审核成功！")
	}
}
