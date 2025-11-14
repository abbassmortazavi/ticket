package services

import "ticket/pkg/auth"

type AuthServiceInterface interface {
	GenerateToken(userID int, username string) *auth.TokenResponse
	ValidateToken(token string) (*auth.Claims, error)
	RefreshToken(refreshToken string) (*auth.TokenResponse, error)
}
