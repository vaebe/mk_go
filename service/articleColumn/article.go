package articleColumn

import (
	"github.com/gin-gonic/gin"
	"mk/global"
	"mk/models/article"
	"mk/models/articleAssociatedInfo"
	"mk/models/articleColumn"
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

	// 查询专栏关联文章
	var articleIds []int32
	var articlesAssociatedColumns []articleAssociatedInfo.ArticlesAssociatedColumns
	res := global.DB.Select("article_id").Where("column_id = ?", columnId).Find(&articlesAssociatedColumns)

	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	// 获取文章id数组
	for _, v := range articlesAssociatedColumns {
		articleIds = append(articleIds, v.ArticleId)
	}

	// 根据文章id 查询返回
	var articles []article.Article
	res = global.DB.Where("id IN ?", articleIds).Find(&articles)

	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
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
	searchForm := articleAssociatedInfo.ArticlesAssociatedColumnsForm{}
	if err := ctx.ShouldBind(&searchForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	userId, _ := ctx.Get("userId")
	res := global.DB.Where("column_id = ? AND article_id = ? AND user_id = ?",
		searchForm.ColumnId, searchForm.ArticleId, userId).Unscoped().Delete(&articleAssociatedInfo.ArticlesAssociatedColumns{})

	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	if res.RowsAffected == 0 {
		utils.ResponseResultsError(ctx, "需要删除的数据不存在！")
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

	// 校验是否是当前用户的专栏
	columnInfo := articleColumn.ArticleColumn{}
	res := global.DB.Where("id = ?", saveForm.ColumnId).First(&columnInfo)
	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	if columnInfo.UserId != userId {
		utils.ResponseResultsError(ctx, "非本用户的专栏无法关联！")
		return
	}

	if columnInfo.Status != "3" {
		utils.ResponseResultsError(ctx, "不是发布状态的专栏无法关联！")
		return
	}

	// 校验是否是当前用户的文章
	articleInfo := article.Article{}
	res = global.DB.Where("id = ?", saveForm.ArticleId).First(&articleInfo)
	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	if articleInfo.UserId != userId {
		utils.ResponseResultsError(ctx, "非本用户的文章无法关联！")
		return
	}

	if articleInfo.Status != "4" {
		utils.ResponseResultsError(ctx, "不是发布状态的文章不能关联！")
		return
	}

	// 验证通过创建关联信息
	res = global.DB.Create(&articleAssociatedInfo.ArticlesAssociatedColumns{
		UserId:    articleInfo.UserId,
		ArticleId: saveForm.ArticleId,
		ColumnId:  saveForm.ColumnId,
	})

	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "关联成功！")
}
