package api

import (
	"errors"
	"net/http"
	"ticket/internal/store"
	"ticket/internal/utils"
	"time"

	"github.com/go-chi/chi/v5"
	log2 "github.com/rs/zerolog/log"
)

type Application struct {
	Config utils.Config
	Store  store.Storage
}

func (app *Application) Start() http.Handler {
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/", app.Create)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", app.GetUser)
				r.Delete("/", app.Delete)
				r.Patch("/", app.Update)
			})
		})
	})
	return r
}
func (app *Application) Run(mux http.Handler) error {
	//adr := app.Config.AppPort
	adr := ":8081"
	srv := &http.Server{
		Addr:         adr,
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}
	log2.Info().Msg("starting server")
	// This will block until the server stops
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log2.Error().Err(err).Msg("server failed to start")
		return err
	}
	return nil
}
