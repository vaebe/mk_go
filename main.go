package main

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"mk/global"
	"mk/initialize"
	middlewares "mk/middleware"
	"mk/routers/article"
	"mk/routers/articleColumn"
	"mk/routers/commentInfo"
	"mk/routers/enum"
	"mk/routers/file"
	"mk/routers/user"
	"time"
)

// @contact.name				API Support
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	// 初始化配置
	initialize.InitConfig()

	r := gin.Default()
	// 替换 gin logger 为 zap
	r.Use(ginzap.Ginzap(global.Logger, time.RFC3339, true), ginzap.RecoveryWithZap(global.Logger, true))

	// 添加跨域中间件
	r.Use(middlewares.Cors())

	// 路由白名单
	routerWhiteList := []string{
		"/mk/user/login",
		"/mk/user/register",
		"/mk/user/details",
		"/mk/user/getVerificationCode",
		"/mk/enum/getAllEnums",
		"/mk/article/getArticleList",
		"/mk/article/getUserArticleList",
		"/mk/article/getArticleDetails",
		"/mk/commentInfo/getListById",
		"/swagger/index.html",
		"/favicon.ico",
	}
	// 添加jwt中间件
	r.Use(middlewares.JWTAuth(routerWhiteList))

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
	serviceAddress := fmt.Sprintf("%s:%d", "127.0.0.1", port)

	// 初始化swagger
	initialize.InitSwagger(r, serviceAddress)

	err := r.Run(serviceAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
}
