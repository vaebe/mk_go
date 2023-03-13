package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mk/initialize"
	middlewares "mk/middleware"
	"mk/routers/article"
	"mk/routers/articleColumn"
	"mk/routers/enum"
	"mk/routers/user"
)

// @contact.name				API Support
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	// 路由白名单
	routerWhiteList := []string{"/mk/user/login", "/swagger/index.html", "/favicon.ico"}

	r := gin.Default()
	r.Use(middlewares.Cors(), middlewares.JWTAuth(routerWhiteList))
	baseRouter := r.Group("/mk")
	{
		user.LoadUserRouter(baseRouter)
		article.LoadArticleRouter(baseRouter)
		enum.LoadEnumRouter(baseRouter)
		articleColumn.LoadArticleColumnRouter(baseRouter)
	}

	serviceAddress := fmt.Sprintf("%s:%d", "127.0.0.1", 53105)

	// 初始化配置
	initialize.InitConfig()

	// 初始化swagger
	initialize.InitSwagger(r, serviceAddress)

	err := r.Run(serviceAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
}
