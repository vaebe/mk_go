package article

import (
	"github.com/gin-gonic/gin"
	"mk/service/article"
)

// LoadArticleRouter 加载文章路由
func LoadArticleRouter(r *gin.RouterGroup) {
	routes := r.Group("article")
	{
		routes.POST("/save", article.Save)
		routes.POST("/saveDraft", article.Draft)
		routes.POST("/getArticleList", article.GetArticleList)
		routes.POST("/getUserArticleList", article.GetUserArticleList)
		routes.GET("/getArticleDetails", article.Details)
		routes.POST("/review", article.Review)
	}
}
