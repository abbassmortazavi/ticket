package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Set() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	err = viper.Unmarshal(&configurations)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
