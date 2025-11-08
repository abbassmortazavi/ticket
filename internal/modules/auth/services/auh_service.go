package services

import (
	"ticket/pkg/auth"
)

type AuthService struct {
	jwtAuth *auth.JWT
}

func New(jwtAuth *auth.JWT) *AuthService {
	return &AuthService{
		jwtAuth: jwtAuth,
	}
}

func (a *AuthService) GenerateToken(userID int, username string) (string, error) {
	return a.jwtAuth.GenerateToken(userID, username)
}

func (a *AuthService) ValidateToken(token string) (*auth.Claims, error) {
	return a.jwtAuth.ValidateToken(token)
}
