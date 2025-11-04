package main

import (
	"ticket/cmd/api"
	"ticket/internal/auth"
	"ticket/internal/store"
	"ticket/pkg/config"
	"ticket/pkg/database"
	"ticket/pkg/logger"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	config.Set()
	database.Connect()
	logger.Init()

	storage := store.NewStorage(database.DB)
	jwt := viper.GetString("JwtSecret")
	authenticator := auth.NewJwtAuthenticator(jwt)

	app := &api.Application{
		Store:         storage,
		Authenticator: authenticator,
	}

	mux := app.Start()
	if err := app.Run(mux); err != nil {
		log.Fatal().Err(err).Msg("server failed to start")
	}
}
