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

func (j *JWT) GenerateToken(userID int, username string) (string, error) {

	//expirationTime := time.Now().Add(24 * time.Hour)
	expirationTime := time.Now().Add(20 * time.Minute) // Shorter lifetime

	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
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
