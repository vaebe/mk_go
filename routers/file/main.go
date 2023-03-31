package file

import (
	"github.com/gin-gonic/gin"
	"mk/service/file"
)

// LoadFileRouter 加载文件操作路由
func LoadFileRouter(r *gin.RouterGroup) {
	routes := r.Group("file")
	{
		routes.POST("/upload", file.Upload)
	}
}
