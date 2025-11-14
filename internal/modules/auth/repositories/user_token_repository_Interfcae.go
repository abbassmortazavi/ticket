package repositories

import (
	"context"
	"ticket/internal/modules/auth/models"
)

type UserTokenRepositoryInterface interface {
	Create(ctx context.Context, token *models.UserToken) error
	FindByUserID(ctx context.Context, id int) (*models.UserToken, error)
	FindByToken(ctx context.Context, token string) (*models.UserToken, error)
	Delete(ctx context.Context, id int) error
	RevokeAllUserTokens(ctx context.Context, id int) error
	CleanedAllExpiredTokens(ctx context.Context) error
}
