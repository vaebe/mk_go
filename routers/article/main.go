package article

import (
	"github.com/gin-gonic/gin"
	"mk/controllers/article"
)

// LoadRouter 加载文章路由
func LoadRouter(r *gin.RouterGroup) {
	routes := r.Group("article")
	{
		routes.POST("/save", article.Save)
		routes.POST("/saveDraft", article.Draft)
		routes.POST("/getArticleList", article.GetArticleList)
		routes.GET("/getArticleDetails", article.Details)
		routes.POST("/review", article.Review)
		routes.DELETE("/delete", article.Delete)
	}
}
