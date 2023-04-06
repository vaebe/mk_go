package articleColumn

import (
	"github.com/gin-gonic/gin"
	"mk/service/articleColumn"
)

// LoadRouter 加载文章专栏路由
func LoadRouter(r *gin.RouterGroup) {
	routes := r.Group("articleColumn")
	{
		routes.POST("/save", articleColumn.Save)
		routes.DELETE("/delete", articleColumn.Delete)
		routes.GET("/details", articleColumn.Details)
		routes.POST("/getList", articleColumn.List)
		routes.POST("/review", articleColumn.Review)
		routes.GET("/getAssociatedArticlesList", articleColumn.GetAssociatedArticlesList)
		routes.POST("/deleteAssociatedArticle", articleColumn.DeleteAssociatedArticle)
		routes.POST("/addAssociatedArticle", articleColumn.AddAssociatedArticle)
		routes.GET("/listArticlesThatCanBeIncluded", articleColumn.ListArticlesThatCanBeIncluded)
	}
}
