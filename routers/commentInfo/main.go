package commentInfo

import (
	"github.com/gin-gonic/gin"
	"mk/services/commentInfo"
)

// LoadRouter 加载评论信息路由
func LoadRouter(r *gin.RouterGroup) {
	routes := r.Group("commentInfo")
	{
		routes.POST("/save", commentInfo.Save)
		routes.GET("/getListById", commentInfo.GetListById)
		routes.GET("/delete", commentInfo.Delete)
	}
}
