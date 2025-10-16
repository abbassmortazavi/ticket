package store

import "database/sql"

type Storage struct {
}

func NewStorage(db *sql.DB) Storage {
	return Storage{}
}
