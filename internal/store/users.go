package store

import (
	"context"
	"database/sql"
	"ticket/internal/models"
)

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user models.User) (int, error) {
	query := `insert into users (username, email, password, full_name, mobile) values ($1,$2, $3, $4,$5) returning id`
	var id int
	err := s.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password, user.FullName, user.Mobile).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
func (s *UserStore) GetUser(ctx context.Context, id int) (models.User, error) {
	var user models.User
	query := `select * from users where id = $1`
	err := s.db.QueryRowContext(ctx, query, id).Scan(
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
func (s *UserStore) Delete(ctx context.Context, id int) error {
	query := `delete from users where id = $1`
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) Update(ctx context.Context, user models.User) (int, error) {
	query := `update users set username= $1, email=$2, password=$3, full_name=$4, mobile=$5 where id=$6 returning id, username`
	err := s.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password, user.FullName, user.Mobile, user.ID).Scan(&user.ID)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (s *UserStore) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var user models.User
	query := `select * from users where username = $1`
	err := s.db.QueryRowContext(ctx, query, username).Scan(
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
