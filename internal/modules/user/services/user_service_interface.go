package services

import (
	"context"
	"ticket/internal/modules/user/models"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GetById(ctx context.Context, id int) (models.User, error)
	GetByUsername(ctx context.Context, username string) (models.User, error)
}
