package config

// RedisConfig 缓存配置
type RedisConfig struct {
	Address            string `mapstructure:"address"`
	Password           string `mapstructure:"password"`
	DB                 int `mapstructure:"db"`
	PoolSize           int `mapstructure:"pool_size"`
	MaxIdleConns       int `mapstructure:"max_idle_conns"`
}

