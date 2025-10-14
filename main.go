package main

import (
	"ticket/internal/utils"

	"github.com/rs/zerolog/log"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("config load failed")
	}

	if config.Debug == "true" {

	}

}
