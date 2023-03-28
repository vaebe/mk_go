package articleColumn

import (
	"github.com/gin-gonic/gin"
	"mk/service/articleColumn"
)

// LoadArticleColumnRouter 加载文章专栏路由
func LoadArticleColumnRouter(r *gin.RouterGroup) {
	userRoutes := r.Group("articleColumn")
	{
		userRoutes.POST("/save", articleColumn.Save)
		userRoutes.DELETE("/delete", articleColumn.Delete)
		userRoutes.GET("/details", articleColumn.Details)
		userRoutes.POST("/getAllArticleColumnList", articleColumn.AllArticleColumnList)
		userRoutes.POST("/review", articleColumn.Review)
	}
}
