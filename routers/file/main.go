package file

import (
	"github.com/gin-gonic/gin"
	"mk/service/file"
)

// LoadFileRouter 加载文件操作路由
func LoadFileRouter(r *gin.RouterGroup) {
	userRoutes := r.Group("file")
	{
		userRoutes.POST("/upload", file.Upload)
	}
}
