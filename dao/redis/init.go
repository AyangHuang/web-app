package redis

import (
	"context"

	"aichat/common/config"

	"github.com/redis/go-redis/v9"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// Init 初始化连接
func Init(conf *config.RedisConfig) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     conf.Address,
		Password: conf.Password, // no password set
		DB:       conf.DB,       // use default DB
		PoolSize: conf.PoolSize,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return
}

func Close() {
	_ = rdb.Close()
}
