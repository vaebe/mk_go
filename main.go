package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mk/initialize"
	middlewares "mk/middleware"
	"mk/routers/article"
	"mk/routers/articleColumn"
	"mk/routers/commentInfo"
	"mk/routers/enum"
	"mk/routers/file"
	"mk/routers/user"
)

// @contact.name				API Support
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	// 路由白名单
	routerWhiteList := []string{"/mk/user/login", "/mk/article/getArticleList",
		"/mk/article/getArticleDetails", "/mk/user/register",
		"/swagger/index.html", "/favicon.ico"}

	r := gin.Default()
	r.Use(middlewares.Cors(), middlewares.JWTAuth(routerWhiteList))
	baseRouter := r.Group("/mk")
	{
		user.LoadRouter(baseRouter)
		article.LoadRouter(baseRouter)
		enum.LoadRouter(baseRouter)
		articleColumn.LoadRouter(baseRouter)
		file.LoadRouter(baseRouter)
		commentInfo.LoadRouter(baseRouter)
	}

	port := 53105
	serviceAddress := fmt.Sprintf("%s:%d", "0.0.0.0", port)

	// 初始化配置
	initialize.InitConfig()

	// 初始化swagger
	initialize.InitSwagger(r, port)

	err := r.Run(serviceAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
}
