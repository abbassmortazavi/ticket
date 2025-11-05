package routing

import (
	"net/http"
	"ticket/pkg/config"

	"github.com/rs/zerolog/log"
)

func Run() {
	router := GetRouter()

	configs := config.Get()

	server := &http.Server{

		Addr:    configs.Host + ":" + configs.AppPort,
		Handler: router,
	}
	log.Info().Msgf("listening on %s", configs.Host+":"+configs.AppPort)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
		panic(err)
		return
	}
}
