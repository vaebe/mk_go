package user

import (
	"github.com/gin-gonic/gin"
	middlewares "mk/middleware"
	"mk/service/user"
)

// LoadRouter 加载用户信息路由
func LoadRouter(r *gin.RouterGroup) {
	routes := r.Group("user")
	{
		routes.POST("/login", user.Login)
		routes.POST("/register", user.Register)
		routes.POST("/sendVerificationCode", user.SendVerificationCode)
		routes.GET("/details", user.Details)
		// 非管理员不能获取用户列表
		routes.POST("/getUserList", middlewares.IsAdmin(), user.GetUserList)
	}
}
