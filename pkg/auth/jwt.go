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
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JWT struct {
	SigningKey []byte
}

func (j *JWT) GenerateToken(userID int, username string) (string, string) {
	accessExpiry := time.Now().Add(time.Minute * 2)
	refreshExpiry := time.Now().Add(time.Minute * 5)

	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpiry),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "auth",
		},
	}

	refreshExpiryClaims := &Claims{
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

	return signedAccessToken, signedRefreshToken
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
