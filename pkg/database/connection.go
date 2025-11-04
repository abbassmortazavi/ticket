package database

import "database/sql"

func Connection() *sql.DB {
	return DB
}
