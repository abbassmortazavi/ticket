package bootstrap

import (
	"ticket/pkg/config"
	"ticket/pkg/database"
	"ticket/pkg/logger"
	"ticket/pkg/routing"
)

func Serve() {
	config.Set()
	database.Connect()
	logger.Init()
	routing.Init()
	routing.RegisterRoutes()
	routing.Run()
}
