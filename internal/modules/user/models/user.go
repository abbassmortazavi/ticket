package models

import (
	"time"

	log2 "github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Mobile    string    `json:"mobile,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log2.Fatal().Err(err).Msg("failed to hash password")
		return err
	}
	u.Password = string(bytes)
	return nil
}
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		log2.Fatal().Err(err).Msg("failed to compare password")
		return false
	}
	return true
}
