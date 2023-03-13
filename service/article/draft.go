package article

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mk/global"
	"mk/models/article"
	"mk/utils"
)

// Draft
//
//	@Summary		保存草稿
//	@Description	保存草稿
//	@Tags			article文章
//	@Accept			json
//	@Produce		json
//	@Param			param	body		article.SaveDraftForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/article/saveDraft [post]
func Draft(ctx *gin.Context) {
	saveDraftForm := article.SaveDraftForm{}

	if err := ctx.ShouldBind(&saveDraftForm); err != nil {
		zap.S().Info("草稿保存信息:", &saveDraftForm)
		utils.HandleValidatorError(ctx, err)
		return
	}

	global.DB.Create(&article.Article{
		UserId:           saveDraftForm.UserId,
		Classify:         saveDraftForm.Classify,
		CollectionColumn: saveDraftForm.CollectionColumn,
		Content:          saveDraftForm.Content,
		CoverImg:         saveDraftForm.CoverImg,
		Tags:             saveDraftForm.Tags,
		Title:            saveDraftForm.Title,
		Summary:          saveDraftForm.Summary,
		Status:           "1",
	})

	utils.ResponseResultsSuccess(ctx, "保存成功！")
}
