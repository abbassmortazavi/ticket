package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func New(dsn, maxIdleTime string, maxOpencons, maxIdlecons int) (*sql.DB, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("connect db failed")
	}
	db.SetMaxIdleConns(maxIdlecons)
	db.SetMaxOpenConns(maxOpencons)
	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		log.Fatal().Err(err).Msg("parse duration failed")
	}
	db.SetConnMaxIdleTime(duration)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal().Err(err).Msg("ping db failed")
	}
	return db, nil
}
