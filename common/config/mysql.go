package config

// MySQLConfig 数据库配置
type MySQLConfig struct {
	DSN string `mapstructure:"dsn"`
	MaxOpenConns int `mapstructure:"max_open_conns"`
	MaxIdleConns int `mapstructure:"max_idle_conns"`
}

