package repositories

import (
	"context"
	models2 "ticket/internal/modules/user/models"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user models2.User) (int, error)
}
