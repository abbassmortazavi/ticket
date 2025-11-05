package repositories

import (
	"context"
	"database/sql"
	models2 "ticket/internal/modules/user/models"
	"ticket/pkg/database"
)

type UserRepository struct {
	DB *sql.DB
}

func New() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}
func (u *UserRepository) Create(ctx context.Context, user models2.User) (int, error) {
	query := `insert into users (username, email, password, full_name, mobile) values ($1,$2, $3, $4,$5) returning id`
	var id int
	err := u.DB.QueryRowContext(ctx, query, user.Username, user.Email, user.Password, user.FullName, user.Mobile).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
