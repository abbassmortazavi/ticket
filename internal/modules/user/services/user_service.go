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
