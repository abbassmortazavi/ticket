package services

import "ticket/pkg/auth"

type AuthServiceInterface interface {
	GenerateToken(userID int, username string) (string, error)
	ValidateToken(token string) (*auth.Claims, error)
}
