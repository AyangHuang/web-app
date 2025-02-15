package config

type AppConfig struct {
	Mode    string `mapstructure:"mode"`
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
	BaseURL string `mapstructure:"baseURL"`
}
