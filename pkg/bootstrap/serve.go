package bootstrap

import (
	"ticket/internal/modules/auth/services"
	"ticket/pkg/auth"
	"ticket/pkg/config"
	"ticket/pkg/database"
	"ticket/pkg/logger"
	"ticket/pkg/middlewares"
	"ticket/pkg/routing"

	"github.com/spf13/viper"
)

func Serve() {

	config.Set()
	// Initialize RabbitMQ with error handling
	/*if err := rabbitmq.Init(); err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}*/
	database.Connect()
	//authentication
	jwtAuth := auth.NewJwtAuthenticator(viper.GetString("JWT_SECRET"))
	authService := services.New(jwtAuth)
	middlewares.Init(authService)

	logger.Init()
	routing.Init()
	routing.RegisterRoutes()

	routing.Run()
}
