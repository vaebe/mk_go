package enum

import (
	"github.com/gin-gonic/gin"
	"mk/service/enum"
)

// LoadEnumRouter 加载枚举路由
func LoadEnumRouter(r *gin.RouterGroup) {
	routes := r.Group("enum")
	{
		routes.POST("/save", enum.Save)
		routes.DELETE("/delete", enum.Delete)
		routes.GET("/details", enum.Details)
		routes.GET("/getEnumsByType", enum.GetEnumsByType)
		routes.GET("/getAllEnums", enum.GetAllEnums)
		routes.POST("/getEnumsList", enum.GetEnumsList)
	}
}
