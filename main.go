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
	//utils.InitLogger()
	logger := zerolog.New(nil)
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("config load failed")
	}

	if config.Debug == "true" {
		log.Info().Msg("debug mode enabled")
	}
	/*dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	config.Username,
	config.Password,
	config.Host,
	config.Port,
	config.Name)*/
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

	log.Debug().Msg("dsn: " + dsn)

	database, err := db.New(dsn, config.MaxIdleTimeout, config.MaxConn, config.MaxIdle)
	if err != nil {
		log.Fatal().Err(err).Msg("connect db failed")
	}
	log.Info().Msg("connect db success")
	storage := store.NewStorage(database)
	app := &api.Application{
		Config: config,
		Store:  storage,
		Logger: logger,
	}
	mux := app.Start()
	if err := app.Run(mux); err != nil {
		logger.Fatal().Err(err).Msg("server failed to start")
	}
}
