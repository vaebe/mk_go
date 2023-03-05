package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mk/initialize"
	"mk/routers"
)

func main() {
	// 初始化日志
	initialize.InitLogger()

	// 初始化配置
	initialize.InitConfig()

	// 初始化mysql
	initialize.InitMysql()

	// 初始化redis
	initialize.InitRedis()

	// 初始化表单验证翻译
	err := initialize.InitTrans("zh")
	if err != nil {
		zap.S().Error("InitTrans:", err.Error())
	}

	// 初始化自定义表单验证规则
	initialize.CustomValidators()

	r := gin.Default()

	// 路由分组
	baseRouter := r.Group("/mk")
	{
		// user 路由
		user.LoadUserRouter(baseRouter)
	}

	//Port, _ := utils.GetFreePort()
	serviceAddress := fmt.Sprintf("%s:%d", "127.0.0.1", 53105)

	err = r.Run(serviceAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
}
