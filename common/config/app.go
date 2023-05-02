package config

type AppConfig struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Address string `mapstructure:"address"`
	Version string `mapstructure:"version"`
}
