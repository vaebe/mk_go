package article

import (
	"github.com/gin-gonic/gin"
	"mk/models"
	"mk/models/article"
	"mk/services/articleServices"
	"mk/utils"
)

// saveArticle 保存文章
func saveArticle(ctx *gin.Context, status string) {
	saveForm := article.SaveForm{}

	if err := ctx.ShouldBind(&saveForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	// 验证是否是当前用户的文章
	userId, _ := ctx.Get("userId")
	if saveForm.UserId != userId {
		utils.ResponseResultsError(ctx, "非当前用户文章不能保存！")
		return
	}

	id, err := articleServices.CreateAndUpdate(saveForm, status, userId.(int32))
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, map[string]any{"id": id})
}

// Draft
//
//	@Summary		保存草稿
//	@Description	保存草稿
//	@Tags			article文章
//	@Accept			json
//	@Produce		json
//	@Param			param	body		article.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/article/saveDraft [post]
func Draft(ctx *gin.Context) {
	saveArticle(ctx, "1")
}

// Save
//
//	@Summary		保存文章
//	@Description	保存文章
//	@Tags			article文章
//	@Accept			json
//	@Produce		json
//	@Param			param	body		article.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/article/save [post]
func Save(ctx *gin.Context) {
	saveArticle(ctx, "2")
}

// GetArticleList
//
//	@Summary		获取文章列表
//	@Description	获取文章列表
//	@Tags			article文章
//	@Accept			json
//	@Produce		json
//	@Param			param	body		article.ListForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Router			/article/getArticleList [post]
func GetArticleList(ctx *gin.Context) {
	listForm := article.ListForm{}

	if err := ctx.ShouldBind(&listForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	list, total, err := articleServices.GetArticleList(listForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, &models.PagingData{
		PageSize: listForm.PageSize,
		PageNo:   listForm.PageNo,
		Total:    total,
		List:     list,
	})
}

// Details
//
//	@Summary		获取文章详情
//	@Description	获取文章详情
//	@Tags			article文章
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"文章id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Router			/article/getArticleDetails [get]
func Details(ctx *gin.Context) {
	articleId := ctx.Query("id")

	if articleId == "" {
		utils.ResponseResultsError(ctx, "文章id不能为空！")
		return
	}

	details, err := articleServices.Details(articleId)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, details)
}

// Delete
//
//	@Summary		文章删除
//	@Description	文章删除
//	@Tags			article文章
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"文章id"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/article/delete [delete]
func Delete(ctx *gin.Context) {
	articleId := ctx.Query("id")

	if articleId == "" {
		utils.ResponseResultsError(ctx, "文章id不能为空！")
		return
	}

	userId, _ := ctx.Get("userId")
	err := articleServices.Delete(articleId, userId.(int32))

	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "删除成功！")
}

// Review
//
//	@Summary		文章审核
//	@Description	文章审核
//	@Tags			article文章
//	@Accept			json
//	@Produce		json
//	@Param			param	body		article.ReviewForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/article/review [post]
func Review(ctx *gin.Context) {
	reviewForm := article.ReviewForm{}
	if err := ctx.ShouldBind(&reviewForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	authorityId, _ := ctx.Get("authorityId")
	if authorityId == 1 {
		utils.ResponseResultsError(ctx, "您没有审核权限！")
		return
	}

	err := articleServices.Review(reviewForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "审核成功！")
}
