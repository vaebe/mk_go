package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mk/initialize"
	"mk/routers"
)

func main() {
	// 初始化配置
	initialize.InitConfig()

	r := gin.Default()
	// 路由分组
	baseRouter := r.Group("/mk")
	{
		user.LoadUserRouter(baseRouter)
	}

	//Port, _ := utils.GetFreePort()
	serviceAddress := fmt.Sprintf("%s:%d", "127.0.0.1", 53105)

	err := r.Run(serviceAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
}
