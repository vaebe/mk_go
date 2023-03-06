package user

import (
	"github.com/gin-gonic/gin"
	"mk/service/user"
)

func LoadUserRouter(r *gin.RouterGroup) {
	userRoutes := r.Group("user")
	{
		userRoutes.GET("/login", user.Login)
		userRoutes.POST("/registerUser", user.Register)
		userRoutes.POST("/sendVerificationCode", user.SendVerificationCode)
	}
}
