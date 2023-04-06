package commentInfo

import (
	"github.com/gin-gonic/gin"
	"mk/service/commentInfo"
)

// LoadRouter 加载评论信息路由
func LoadRouter(r *gin.RouterGroup) {
	routes := r.Group("commentInfo")
	{
		routes.POST("/save", commentInfo.Save)
	}
}
