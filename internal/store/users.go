package store

import (
	"context"
	"database/sql"
)

type UserStore struct {
	db *sql.DB
}

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	Mobile    string `json:"mobile,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"Updated_at"`
}

func (s *UserStore) Create(ctx context.Context, user User) (int, error) {
	query := `insert into users (username, email, password, full_name, mobile) values ($1,$2, $3, $4,$5) returning id`
	var id int
	err := s.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password, user.FullName, user.Mobile).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
