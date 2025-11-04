package main

import (
	"ticket/cmd"
	"ticket/cmd/api"
	"ticket/internal/auth"
	"ticket/internal/store"
	"ticket/pkg/database"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	cmd.Execute()

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
