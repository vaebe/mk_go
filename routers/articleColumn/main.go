package articleColumn

import (
	"github.com/gin-gonic/gin"
	"mk/service/articleColumn"
)

// LoadArticleColumnRouter 加载文章专栏路由
func LoadArticleColumnRouter(r *gin.RouterGroup) {
	routes := r.Group("articleColumn")
	{
		routes.POST("/save", articleColumn.Save)
		routes.DELETE("/delete", articleColumn.Delete)
		routes.GET("/details", articleColumn.Details)
		routes.POST("/getList", articleColumn.List)
		routes.POST("/review", articleColumn.Review)
	}
}
