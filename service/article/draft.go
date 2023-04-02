package article

import (
	"github.com/gin-gonic/gin"
)

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
	SaveArticle(ctx, "1")
}
