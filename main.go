package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mk/initialize"
	"mk/routers/article"
	"mk/routers/user"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	r := gin.Default()
	baseRouter := r.Group("/mk")
	{
		user.LoadUserRouter(baseRouter)
		article.LoadArticleRouter(baseRouter)
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
