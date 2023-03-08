package user

import (
	"github.com/gin-gonic/gin"
	"mk/service/user"
)

func LoadUserRouter(r *gin.RouterGroup) {
	userRoutes := r.Group("user")
	{
		userRoutes.POST("/login", user.Login)
		userRoutes.POST("/register", user.Register)
		userRoutes.POST("/sendVerificationCode", user.SendVerificationCode)
		userRoutes.POST("/getUserList", user.GetUserList)
	}
}
