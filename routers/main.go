package routers

import (
	"github.com/gin-gonic/gin"
	"mk/routers/article"
	"mk/routers/articleColumn"
	"mk/routers/commentInfo"
	"mk/routers/enum"
	"mk/routers/file"
	"mk/routers/user"
)

// GetRouterWhiteList 获取路由白名单
func GetRouterWhiteList() []string {
	return []string{
		"/mk/user/login",
		"/mk/user/register",
		"/mk/user/details",
		"/mk/user/getVerificationCode",
		"/mk/enum/getAllEnums",
		"/mk/article/getArticleList",
		"/mk/article/getArticleDetails",
		"/mk/articleColumn/getList",
		"/mk/articleColumn/details",
		"/mk/articleColumn/getAssociatedArticlesList",
		"/mk/commentInfo/getListById",
		"/swagger/index.html",
		"/favicon.ico",
	}
}

// LoadAllRouter 加载全部路由
func LoadAllRouter(r *gin.Engine) {
	baseRouter := r.Group("/mk")
	{
		user.LoadRouter(baseRouter)
		article.LoadRouter(baseRouter)
		enum.LoadRouter(baseRouter)
		articleColumn.LoadRouter(baseRouter)
		file.LoadRouter(baseRouter)
		commentInfo.LoadRouter(baseRouter)
	}
}
