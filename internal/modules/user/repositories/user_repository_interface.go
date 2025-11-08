package repositories

import (
	"context"
	"ticket/internal/modules/user/models"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user models.User) (int, error)
	GetById(ctx context.Context, id int) (models.User, error)
	GetByUsername(ctx context.Context, username string) (models.User, error)
	Delete(ctx context.Context, id int) error
}
