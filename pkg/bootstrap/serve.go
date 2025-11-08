package bootstrap

import (
	"ticket/pkg/auth"
	"ticket/pkg/config"
	"ticket/pkg/database"
	"ticket/pkg/logger"
	"ticket/pkg/routing"

	"github.com/spf13/viper"
)

func Serve() {
	config.Set()
	database.Connect()
	_ = auth.NewJwtAuthenticator(viper.GetString("JWT_SECRET"))
	logger.Init()
	routing.Init()
	routing.RegisterRoutes()
	routing.Run()
}
