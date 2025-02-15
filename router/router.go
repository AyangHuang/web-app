package router

import (
	"aichat/common/logger"
	"aichat/controller"
	_ "aichat/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Setup title 和 version 必须填写，不然不符合 swagger 标准，不能导入 postman
// swag init -g ./router/router.go 指定搜索该文件
// 访问：$baseURL/swagger/index.html
// @title           AIChat
// @version         0.0.0
// @host 127.0.0.1:8081
// @BasePath /aichat
func Setup(mode string, baseURL string) *gin.Engine {
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

	baseUrl := e.Group(baseURL)

	// 日志和 recovery 中间件
	baseUrl.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 设置 header 中间件
	baseUrl.Use(controller.CorsHandler())

	registerSwagger(baseUrl)
	registerExample(baseUrl)

	return e
}

func registerSwagger(group *gin.RouterGroup) {
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func registerExample(group *gin.RouterGroup) {
	group.POST("/example", controller.ExampleHandler)
}
