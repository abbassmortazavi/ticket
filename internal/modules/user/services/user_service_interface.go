package services

import (
	"context"
	"ticket/internal/modules/user/models"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
}
