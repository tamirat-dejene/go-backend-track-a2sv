package usecases

import (
	"context"
	domain "t7/taskmanager/Domain"
	"time"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

// Login implements domain.UserUsecase.
func (u *userUsecase) Login(ctx context.Context, user *domain.User) (*domain.User, error) {
	context, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepository.Login(context, user)
}

// DeleteUser implements domain.UserUsecase.
func (u *userUsecase) Delete(ctx context.Context, user_name string) error {
	context, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepository.Delete(context, user_name)
}

// FindUser implements domain.UserUsecase.
func (u *userUsecase) FindOne(ctx context.Context, user_name string) (*domain.User, error) {
	context, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepository.FindOne(context, user_name)
}

// GetAll implements domain.UserUsecase.
func (u *userUsecase) GetAll(ctx context.Context) ([]domain.User, error) {
	context, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepository.GetAll(context)
}

// Register implements domain.UserUsecase.
func (u *userUsecase) Register(ctx context.Context, user *domain.User) (string, error) {
	context, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepository.Register(context, user)
}

func NewUserUsecase(userRepo domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepo,
		contextTimeout: timeout,
	}
}
