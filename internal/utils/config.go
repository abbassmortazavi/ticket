package utils

import "github.com/spf13/viper"

type Config struct {
	Host           string `mapstructure:"DB_HOST"`
	Port           string `mapstructure:"DB_PORT"`
	Debug          string `mapstructure:"APP_DEBUG"`
	Username       string `mapstructure:"DB_USERNAME"`
	Password       string `mapstructure:"DB_PASSWORD"`
	Name           string `mapstructure:"DB_NAME"`
	MaxIdle        int    `mapstructure:"DB_MAX_IDLE"`
	MaxConn        int    `mapstructure:"DB_MAX_CONN"`
	MaxIdleTimeout string `mapstructure:"DB_MAX_IDLE_TIMEOUT"`
	AppPort        string `mapstructure:"APP_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil

}
