package main

import (
	"fmt"
	"os"
	"ticket/cmd/api"
	"ticket/internal/auth"
	"ticket/internal/db"
	"ticket/internal/store"
	"ticket/pkg/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	config.Set()
	// Initialize global logger
	initGlobalLogger(viper.GetString("APP_DEBUG") == "true")

	if viper.GetString("APP_DEBUG") == "true" {
		log.Info().Msg("debug mode enabled")
	}

	var myHost string
	if host := os.Getenv("DB_HOST"); host != "" {
		myHost = host
	} else {
		myHost = viper.GetString("DB_HOST")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		myHost,
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"))

	log.Info().Msgf("dsn: %s", dsn)

	database, err := db.New(dsn, viper.GetString("DB_MAX_IDLE_TIMEOUT"), viper.GetInt("DB_MAX_CONN"), viper.GetInt("DB_MAX_IDLE"))
	if err != nil {
		log.Fatal().Err(err).Msg("connect db failed")
	}

	storage := store.NewStorage(database)
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

func initGlobalLogger(debug bool) {
	// Console output for development
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
	}

	log.Logger = zerolog.New(output).With().Timestamp().Logger()

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
