package enum

import (
	"github.com/gin-gonic/gin"
	"mk/service/enum"
)

// LoadEnumRouter 加载枚举路由
func LoadEnumRouter(r *gin.RouterGroup) {
	userRoutes := r.Group("enum")
	{
		userRoutes.POST("/save", enum.Save)
		userRoutes.DELETE("/delete", enum.Delete)
		userRoutes.GET("/details", enum.Details)
		userRoutes.GET("/getEnumsByType", enum.GetEnumsByType)
		userRoutes.GET("/getAllEnums", enum.GetAllEnums)
		userRoutes.POST("/getEnumsList", enum.GetEnumsList)
	}
}
