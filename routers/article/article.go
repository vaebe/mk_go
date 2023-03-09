package article

import (
	"github.com/gin-gonic/gin"
	"mk/service/article"
)

// LoadArticleRouter 加载文章路由
func LoadArticleRouter(r *gin.RouterGroup) {
	userRoutes := r.Group("article")
	{
		userRoutes.POST("/save", article.Save)
		userRoutes.POST("/saveDraft", article.Draft)
		userRoutes.POST("/getArticleList", article.GetArticleList)
		userRoutes.POST("/getUserArticleList", article.GetUserArticleList)
		userRoutes.GET("/getArticleDetails", article.Details)
	}
}
