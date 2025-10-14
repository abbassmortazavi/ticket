package utils

import "github.com/spf13/viper"

type Config struct {
	Host  string `mapstructure:"APP_HOST"`
	Port  string `mapstructure:"APP_PORT"`
	Debug string `mapstructure:"APP_DEBUG"`
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
