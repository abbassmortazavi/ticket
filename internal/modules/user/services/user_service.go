package services

import (
	"context"
	"ticket/internal/modules/user/models"
	userRepository "ticket/internal/modules/user/repositories"
)

type UserService struct {
	userRepository userRepository.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: userRepository.New(),
	}
}
func (u *UserService) CreateUser(ctx context.Context, user models.User) (int, error) {
	return u.userRepository.Create(ctx, user)
}
func (u *UserService) GetById(ctx context.Context, id int) (models.User, error) {
	return u.userRepository.GetById(ctx, id)
}
func (u *UserService) GetByUsername(ctx context.Context, username string) (models.User, error) {
	return u.userRepository.GetByUsername(ctx, username)
}
