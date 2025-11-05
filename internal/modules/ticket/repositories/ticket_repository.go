package repositories

import (
	"database/sql"
	"ticket/pkg/database"
)

type TicketRepository struct {
	DB *sql.DB
}

func New() *TicketRepository {
	return &TicketRepository{
		DB: database.Connection(),
	}

}
