package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // path to look for the config file in

	// 读取配置文件
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 序列化到当前包变量中
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

// Conf 兜底配置，即默认配置。后续配置文件会写入覆盖
var Conf GlobalConf = GlobalConf{
	MySQLConfig: &MySQLConfig{
		MaxOpenConns: 100,
		MaxIdleConns: 50,
	},
	RedisConfig: &RedisConfig{
		DB:           0,
		PoolSize:     100,
		MaxIdleConns: 50,
	},
}

type GlobalConf struct {
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*AppConfig   `mapstructure:"app"`
	*LogConfig   `mapstructure:"log"`
}
