package bootstrap

import (
	"ticket/internal/modules/auth/services"
	"ticket/pkg/auth"
	"ticket/pkg/config"
	"ticket/pkg/database"
	"ticket/pkg/helpers"
	"ticket/pkg/logger"
	"ticket/pkg/middlewares"
	"ticket/pkg/routing"
)

func Serve() {

	config.Set()
	// Initialize RabbitMQ with error handling
	/*if err := rabbitmq.Init(); err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}*/
	database.Connect()
	//authentication
	jwtAuth := auth.NewJwtAuthenticator(helpers.GenerateRandomKey())
	authService := services.New(jwtAuth)
	middlewares.Init(authService)

	logger.Init()
	routing.Init()
	routing.RegisterRoutes()

	routing.Run()
}
