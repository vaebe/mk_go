package middlewares

import (
	"github.com/gin-gonic/gin"
	"mk/utils"
)

// IsAdmin 验证登陆用户是否是admin
func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取jwt验证后设置的用户信息
		claims, _ := ctx.Get("claims")

		// AuthorityId != 2 表示非管理员
		if claims != 2 {
			utils.ResponseResultsError(ctx, "用户无权限！")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
