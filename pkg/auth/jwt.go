package auth

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func NewJwtAuthenticator(j string) *JWT {
	if j == "" {
		panic("JWT secret cannot be empty")
	}
	JwtAuthenticator = &JWT{
		SigningKey: []byte(j),
	}
	log.Println("JWT Authenticator initialized")
	return JwtAuthenticator
}

type Claims struct {
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

type JWT struct {
	SigningKey []byte
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

func (j *JWT) GenerateToken(userID int, username string) *TokenResponse {
	accessExpiry := time.Now().Add(time.Minute * 2)
	refreshExpiry := time.Now().Add(time.Minute * 5)

	claims := Claims{
		UserID:    userID,
		Username:  username,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpiry),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "auth",
		},
	}

	refreshExpiryClaims := &Claims{
		UserID:    userID,
		Username:  username,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpiry),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedAccessToken, err := accessToken.SignedString(j.SigningKey)
	if err != nil {
		panic(err)
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshExpiryClaims)
	signedRefreshToken, err := refreshToken.SignedString(j.SigningKey)
	if err != nil {
		panic(err)
	}

	return &TokenResponse{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
		ExpiresAt:    accessExpiry.Unix(),
	}
}
func (j *JWT) RefreshAccessToken(refreshToken string) (*TokenResponse, error) {
	claims, err := j.ValidateToken(refreshToken)
	if err != nil {
		return nil, err
	}
	if claims.TokenType != "refresh" {
		return nil, errors.New("invalid token type")
	}
	return j.GenerateToken(claims.UserID, claims.Username), nil
}

func (j *JWT) ValidateToken(token string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
