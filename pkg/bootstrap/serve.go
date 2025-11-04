package bootstrap

import (
	"ticket/pkg/config"
	"ticket/pkg/database"
	"ticket/pkg/logger"
)

func Serve() {
	config.Set()
	database.Connect()
	logger.Init()
}
