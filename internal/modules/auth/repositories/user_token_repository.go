package repositories

import (
	"context"
	"database/sql"
	"errors"
	"ticket/internal/modules/auth/models"
	"ticket/pkg/database"
)

type UserTokenRepository struct {
	DB *sql.DB
}

func New() *UserTokenRepository {
	return &UserTokenRepository{
		DB: database.Connection(),
	}
}
func (r *UserTokenRepository) FindByToken(ctx context.Context, token string) (*models.UserToken, error) {
	userToken := &models.UserToken{}
	query := `SELECT * FROM user_tokens WHERE hash_token = $1`
	row := r.DB.QueryRowContext(ctx, query, token)
	err := row.Scan(
		&userToken.ID,
		&userToken.UserID,
		&userToken.TokenType,
		&userToken.HashToken,
		&userToken.IsRevoked,
		&userToken.ExpiredAt,
		&userToken.CreatedAt,
		&userToken.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
	}
	return userToken, nil
}
func (r *UserTokenRepository) FindByUserID(ctx context.Context, id int) (*models.UserToken, error) {
	userToken := &models.UserToken{}
	query := `SELECT * FROM user_tokens WHERE user_id = $1`
	row := r.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&userToken.ID,
		&userToken.UserID,
		&userToken.TokenType,
		&userToken.HashToken,
		&userToken.ExpiredAt,
		&userToken.IsRevoked,
		&userToken.CreatedAt,
		&userToken.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
	}
	return userToken, nil
}
func (r *UserTokenRepository) Create(ctx context.Context, userToken *models.UserToken) error {
	query := `INSERT INTO user_tokens (user_id, token_type, hash_token, expired_at, is_revoked) 
			values ($1, $2, $3, $4, $5) RETURNING *`
	row := r.DB.QueryRowContext(ctx, query,
		userToken.UserID, userToken.TokenType,
		userToken.HashToken, userToken.ExpiredAt, userToken.IsRevoked)
	err := row.Scan(
		&userToken.ID,
		&userToken.UserID,
		&userToken.TokenType,
		&userToken.HashToken,
		&userToken.ExpiredAt,
		&userToken.IsRevoked,
		&userToken.CreatedAt,
		&userToken.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserTokenRepository) RevokeAllUserTokens(ctx context.Context, userID int) error {
	query := `DELETE FROM user_tokens WHERE user_id = $1`
	_, err := r.DB.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserTokenRepository) CleanedAllExpiredTokens(ctx context.Context) error {
	query := `DELETE FROM user_tokens WHERE expired_at > CURRENT_TIMESTAMP`
	_, err := r.DB.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
