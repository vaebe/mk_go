package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"mk/docs"
)

// InitSwagger 初始化swagger
func InitSwagger(r *gin.Engine, serviceAddress string) {
	docs.SwaggerInfo.Title = "MK API"
	docs.SwaggerInfo.Description = "MK API server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = serviceAddress
	docs.SwaggerInfo.BasePath = "/mk"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
}
