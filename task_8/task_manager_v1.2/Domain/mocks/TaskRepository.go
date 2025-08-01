package mocks

import (
	"context"
	domain "t8/taskmanager/Domain"

	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Add(ctx context.Context, task *domain.Task) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}

func (m *MockTaskRepository) Remove(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockTaskRepository) Update(ctx context.Context, id string, task *domain.Task) (domain.Task, error) {
	args := m.Called(ctx, id, task)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetAll(ctx context.Context) ([]domain.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetOne(ctx context.Context, id string) (domain.Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Task), args.Error(1)
}
