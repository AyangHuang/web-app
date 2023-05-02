package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web-app/common/config"
	"web-app/common/logger"
	"web-app/dao/mysql"
	"web-app/dao/redis"
	"web-app/router"
)

func main() {
	// 1. 加载配置
	config.Init()

	// 2. 初始化日志
	logger.Init(config.Conf.LogConfig)

	// 3. 初始化数据库连接（包括 mysql 和 redis）
	mysql.Init(config.Conf.MySQLConfig)
	redis.Init(config.Conf.RedisConfig)
	defer redis.Close()

	// 4. 注册路由
	e := router.Setup(gin.DebugMode)

	// 5. 启动服务和优雅关机
	srv := &http.Server{
		Addr:    ":8080",
		Handler: e,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// 注册监听 chan，收到信号操作系统会调用 go 库函数往 quit 里面发送一个 os.Signal
	// syscall.SIGINT kill -2 pid
	signal.Notify(quit, syscall.SIGINT)

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}
	zap.L().Info("Server exiting!")
}
