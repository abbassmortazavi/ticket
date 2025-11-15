package bootstrap

import (
	"ticket/internal/modules/auth/repositories"
	auth2 "ticket/internal/modules/auth/services"
	"ticket/pkg/auth"
	"ticket/pkg/config"
	"ticket/pkg/database"
	"ticket/pkg/helpers"
	"ticket/pkg/html"
	"ticket/pkg/logger"
	"ticket/pkg/middlewares"
	"ticket/pkg/routing"
	"ticket/pkg/static"
)

func Serve() {

	config.Set()
	// Initialize RabbitMQ with error handling
	/*if err := rabbitmq.Init(); err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}*/
	database.Connect()
	//authentication
	jwtAuth := auth.NewJwtAuthenticator(helpers.GenerateRandomKey(), repositories.UserTokenRepository{
		DB: database.DB,
	})
	authService := auth2.New(jwtAuth)
	middlewares.Init(authService)

	logger.Init()
	routing.Init()
	html.LoadHtml(routing.GetRouter())
	routing.RegisterRoutes()
	static.LoadAsset(routing.GetRouter())
	routing.Run()
}
