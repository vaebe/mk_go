package file

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Upload
//
//	@Summary		文件上传
//	@Description	文件上传
//	@Tags			file文件
//	@Accept			json
//	@Produce		json
//	@Param			param	formData  file	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/file/upload [post]
func Upload(ctx *gin.Context) {
	// 单文件
	file, _ := ctx.FormFile("file")
	zap.S().Info(file.Filename)
}
