package articleColumn

import (
	"github.com/gin-gonic/gin"
	"mk/global"
	"mk/models/article"
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
	var articlesAssociatedColumns []article.ArticlesAssociatedColumns
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
