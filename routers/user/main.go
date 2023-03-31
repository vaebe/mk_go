package user

import (
	"github.com/gin-gonic/gin"
	"mk/service/user"
)

func LoadUserRouter(r *gin.RouterGroup) {
	routes := r.Group("user")
	{
		routes.POST("/login", user.Login)
		routes.POST("/register", user.Register)
		routes.POST("/sendVerificationCode", user.SendVerificationCode)
		routes.POST("/getUserList", user.GetUserList)
		routes.GET("/getUserDetails", user.Details)
	}
}
