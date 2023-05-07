package articleColumn

import (
	"github.com/gin-gonic/gin"
	"mk/models/articleAssociatedInfo"
	"mk/services/articleColumnServices"
	"mk/utils"
)

// GetAssociatedArticlesList
//
//	@Summary			获取专栏关联文章列表
//	@Description	获取专栏关联文章列表
//	@Tags					articleColumn专栏
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"专栏id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/articleColumn/getAssociatedArticlesList [get]
func GetAssociatedArticlesList(ctx *gin.Context) {
	columnId := ctx.Query("id")

	if columnId == "" {
		utils.ResponseResultsError(ctx, "专栏id不能为空！")
		return
	}

	articles, err := articleColumnServices.GetAssociatedArticlesList(columnId)

	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, articles)
}

// ListArticlesThatCanBeIncluded
//
//	@Summary			获取可以被收录的文章列表
//	@Description	获取可以被收录的文章列表
//	@Tags					articleColumn专栏
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"专栏id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/articleColumn/listArticlesThatCanBeIncluded [get]
func ListArticlesThatCanBeIncluded(ctx *gin.Context) {
	columnId := ctx.Query("id")

	if columnId == "" {
		utils.ResponseResultsError(ctx, "专栏id不能为空！")
		return
	}

	articles, err := articleColumnServices.ListArticlesThatCanBeIncluded(columnId)

	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, articles)
}

// DeleteAssociatedArticle
//
//	@Summary			删除关联文章
//	@Description	删除关联文章
//	@Tags				articleColumn专栏
//	@Accept			json
//	@Produce		json
//	@Param			param	body		articleAssociatedInfo.ArticlesAssociatedColumnsForm	true	"请求对象"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/articleColumn/deleteAssociatedArticle [post]
func DeleteAssociatedArticle(ctx *gin.Context) {
	delForm := articleAssociatedInfo.ArticlesAssociatedColumnsForm{}
	if err := ctx.ShouldBind(&delForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	userId, _ := ctx.Get("userId")
	err := articleColumnServices.DeleteAssociatedArticle(delForm, userId.(int32))

	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "删除成功！")
}

// AddAssociatedArticle
//
//	@Summary		添加关联文章
//	@Description	添加关联文章
//	@Tags				articleColumn专栏
//	@Accept			json
//	@Produce		json
//	@Param			param	body		articleAssociatedInfo.ArticlesAssociatedColumnsForm	true	"请求对象"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/articleColumn/addAssociatedArticle [post]
func AddAssociatedArticle(ctx *gin.Context) {
	saveForm := articleAssociatedInfo.ArticlesAssociatedColumnsForm{}

	if err := ctx.ShouldBind(&saveForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	userId, _ := ctx.Get("userId")

	err := articleColumnServices.AddAssociatedArticle(saveForm, userId.(int32))

	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "关联成功！")
}
