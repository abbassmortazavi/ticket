package store

import (
	"context"
	"database/sql"
	"ticket/internal/models"
)

type BusStore struct {
	db *sql.DB
}

func (b *BusStore) Create(ctx context.Context, bus models.Bus) (models.Bus, error) {
	//q := `insert into bus (bus_code) values ($1) returning bus_code;`
	return models.Bus{}, nil
}
