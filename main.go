package main

import (
	"fmt"
	"os"
	"ticket/cmd/api"
	"ticket/internal/db"
	"ticket/internal/store"
	"ticket/internal/utils"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("config load failed")
	}

	// Initialize global logger
	initGlobalLogger(config.Debug == "true")

	if config.Debug == "true" {
		log.Info().Msg("debug mode enabled")
	}

	if host := os.Getenv("DB_HOST"); host != "" {
		config.Host = host
		log.Info().Msg("db host set to: " + config.Host)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.Name)

	database, err := db.New(dsn, config.MaxIdleTimeout, config.MaxConn, config.MaxIdle)
	if err != nil {
		log.Fatal().Err(err).Msg("connect db failed")
	}
	log.Info().Msg("connect db success")

	storage := store.NewStorage(database)
	app := &api.Application{
		Config: config,
		Store:  storage,
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
