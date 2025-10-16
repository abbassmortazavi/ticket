package api

import (
	"errors"
	"log"
	"net/http"
	"ticket/internal/store"
	"ticket/internal/utils"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type Application struct {
	Config utils.Config
	Store  store.Storage
	Logger zerolog.Logger
}

func (app *Application) Start() http.Handler {
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		r.Route("/posts", func(r chi.Router) {
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", app.GetUser)
			})
		})
	})
	return r
}
func (app *Application) Run(mux http.Handler) error {

	log.Println("starting server")

	adr := app.Config.AppPort
	srv := &http.Server{
		Addr:         adr,
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}
	app.Logger.Info().Msg("starting server")
	// This will block until the server stops
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		app.Logger.Error().Err(err).Msg("server failed to start")
		return err
	}
	return nil
}
