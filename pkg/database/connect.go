package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq" // This is the PostgreSQL driver import

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Connect() {
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

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal().Err(err).Msg("connect db failed")
	}

	db.SetMaxIdleConns(viper.GetInt("DB_MAX_IDLE"))
	db.SetMaxOpenConns(viper.GetInt("DB_MAX_CONN"))
	duration, err := time.ParseDuration(viper.GetString("DB_MAX_IDLE_TIMEOUT"))
	if err != nil {
		log.Fatal().Err(err).Msg("parse duration failed")
	}
	db.SetConnMaxIdleTime(duration)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal().Err(err).Msg("ping db failed")
	}
	DB = db
}
