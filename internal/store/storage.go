package store

import (
	"context"
	"database/sql"
	"ticket/internal/models"
)

type Storage struct {
	User interface {
		Create(ctx context.Context, user models.User) (int, error)
		GetUser(ctx context.Context, id int) (models.User, error)
		GetUserByUsername(ctx context.Context, username string) (models.User, error)
		Delete(ctx context.Context, id int) error
		Update(ctx context.Context, user models.User) (int, error)
	}
	Bus interface {
		Create(ctx context.Context, bus models.Bus) (models.Bus, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		User: &UserStore{db},
	}
}
