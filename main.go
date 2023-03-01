package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mk/initialize"
	"mk/routers"
)

func main() {
	// 初始化日志
	initialize.InitLogger()

	// 初始化配置
	initialize.InitConfig()

	r := gin.Default()

	// 路由分组
	baseRouter := r.Group("/mk")
	{
		// user 路由
		user.LoadUserRouter(baseRouter)
	}

	err := r.Run(":3234")
	if err != nil {
		fmt.Println(err)
		return
	}
}
