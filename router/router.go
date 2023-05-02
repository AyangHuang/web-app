package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"web-app/common/logger"
	"web-app/controller"
	_ "web-app/docs"
)

func Setup(mode string) *gin.Engine {
	switch mode {
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	case gin.TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	e := gin.New()

	v1 := e.Group("/api/v1")

	// 日志和 recovery 中间件
	v1.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 设置 header 中间件
	v1.Use(controller.CrosHandler())

	registerSwagger(v1)
	registerExample(v1)

	return e
}

func registerSwagger(group *gin.RouterGroup) {
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func registerExample(group *gin.RouterGroup) {
	group.POST("/example", controller.ExampleHandler)
}

