package db

import (
	"database/sql"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(dsn string) (*sql.DB, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("connect db failed")
	}
	//log.Fatal().Err(err).Msg("db open failed")

}
