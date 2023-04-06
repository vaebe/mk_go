package commentInfo

import (
	"github.com/gin-gonic/gin"
	"mk/global"
	"mk/models/commentInfo"
	"mk/utils"
)

// Save
//
//	@Summary		保存评论信息
//	@Description	保存评论信息
//	@Tags			commentInfo评论
//	@Accept			json
//	@Produce		json
//	@Param			param	body		commentInfo.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/commentInfo/save [post]
func Save(ctx *gin.Context) {
	saveForm := commentInfo.SaveForm{}

	if err := ctx.ShouldBind(&saveForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	res := global.DB.Create(&saveForm)
	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}
	utils.ResponseResultsSuccess(ctx, "保存成功！")
}
