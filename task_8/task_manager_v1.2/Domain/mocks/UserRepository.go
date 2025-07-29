package mocks

import (
	"context"
	domain "t8/taskmanager/Domain"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindOne(ctx context.Context, user_name string) (*domain.User, error) {
	args := m.Called(ctx, user_name)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) Register(ctx context.Context, user *domain.User) (string, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(string), args.Error(1)
}

func (m *MockUserRepository) Delete(ctx context.Context, user_name string) error {
	args := m.Called(ctx, user_name)
	return args.Error(0)
}

func (m *MockUserRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserRepository) Login(ctx context.Context, user *domain.User) (*domain.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*domain.User), args.Error(1)
}
