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

func (a *AuthService) GenerateToken(userID int, username string) *auth.TokenResponse {
	res := a.jwtAuth.GenerateToken(userID, username)
	return &auth.TokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		ExpiresAt:    res.ExpiresAt,
	}
}

func (a *AuthService) ValidateToken(token string) (*auth.Claims, error) {
	return a.jwtAuth.ValidateToken(token)
}
func (a *AuthService) RefreshToken(refreshToken string) (*auth.TokenResponse, error) {
	tokens, err := a.jwtAuth.RefreshAccessToken(refreshToken)
	if err != nil {
		return nil, err
	}
	return &auth.TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		ExpiresAt:    tokens.ExpiresAt,
	}, nil
}
