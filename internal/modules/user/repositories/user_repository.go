package repositories

import (
	"context"
	"database/sql"
	models "ticket/internal/modules/user/models"
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
func (u *UserRepository) Create(ctx context.Context, user models.User) (int, error) {
	query := `insert into users (username, email, password, full_name, mobile) values ($1,$2, $3, $4,$5) returning id`
	var id int
	err := u.DB.QueryRowContext(ctx, query, user.Username, user.Email, user.Password, user.FullName, user.Mobile).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
func (u *UserRepository) GetById(ctx context.Context, id int) (models.User, error) {
	var user models.User
	query := `select * from users where id = $1`
	err := u.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.Mobile,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (u *UserRepository) GetByUsername(ctx context.Context, username string) (models.User, error) {
	var user models.User
	query := `select * from users where username = $1`
	err := u.DB.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.Mobile,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (u *UserRepository) Delete(ctx context.Context, id int) error {
	query := `delete from users where id = $1`
	_, err := u.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserRepository) Update(ctx context.Context, user models.User) (int, error) {
	query := `update users set username=$1,email=$2, password=$3, full_name=$4, mobile=$5 where id=$6 returning id, username`
	err := u.DB.QueryRowContext(ctx, query, user.ID, user.Username, user.Email, user.Password, user.FullName, user.Mobile, user.ID).Scan(&user.ID)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
