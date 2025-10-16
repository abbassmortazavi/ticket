package main

import (
	"fmt"
	"ticket/internal/db"
	"ticket/internal/utils"

	"github.com/rs/zerolog/log"
)

func main() {
	utils.InitLogger()
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("config load failed")
	}

	if config.Debug == "true" {
		log.Info().Msg("debug mode enabled")
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name)

	_, err = db.ConnectDB(dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("connect db failed")
	}

}
