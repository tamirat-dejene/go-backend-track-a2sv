package usecases

import (
	"context"
	"errors"
	domain "t8/taskmanager/Domain"
	"t8/taskmanager/Domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestAddSuccess should add a task successfully
func TestAddSuccess(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	usecase := &taskUsecase{
		taskRepository: mockTaskRepo,
		contextTimeout: 5 * time.Second,
	}

	task := &domain.Task{
		ID:          "1",
		Title:       "Test task",
		Description: "desc",
		DueDate:     time.Now(),
		Status:      "pending",
	}

	mockTaskRepo.On("Add", mock.Anything, task).Return(nil).Once()

	err := usecase.Add(context.Background(), task)

	assert.NoError(t, err)
	mockTaskRepo.AssertExpectations(t)
}

// TestAddFails should return an error when adding a task fails
func TestAddFails(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	usecase := &taskUsecase{
		taskRepository: mockTaskRepo,
		contextTimeout: 5 * time.Second,
	}

	task := &domain.Task{
		ID:          "1",
		Title:       "Test task",
		Description: "desc",
		DueDate:     time.Now(),
		Status:      "pending",
	}

	mockTaskRepo.On("Add", mock.Anything, task).Return(errors.New("insert error")).Once()

	err := usecase.Add(context.Background(), task)

	assert.Error(t, err)
	mockTaskRepo.AssertExpectations(t)
}

// TestGetAllSuccess should return all tasks successfully
func TestGetAllSuccess(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	usecase := &taskUsecase{
		taskRepository: mockTaskRepo,
		contextTimeout: 5 * time.Second,
	}

	tasks := []domain.Task{
		{ID: "1", Title: "task1"},
		{ID: "2", Title: "task2"},
	}

	mockTaskRepo.On("GetAll", mock.Anything).Return(tasks, nil).Once()

	result, err := usecase.GetAll(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, tasks, result)
	mockTaskRepo.AssertExpectations(t)
}

// TestGetAllFails should return an error when fetching tasks fails
func TestGetAllFails(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	usecase := &taskUsecase{
		taskRepository: mockTaskRepo,
		contextTimeout: 5 * time.Second,
	}

	mockTaskRepo.On("GetAll", mock.Anything).Return([]domain.Task{}, errors.New("fetch error")).Once()

	result, err := usecase.GetAll(context.Background())

	assert.Error(t, err)
	assert.Empty(t, result)
	mockTaskRepo.AssertExpectations(t)
}

// TestGetOneSuccess should return a task by ID successfully
func TestGetOneSuccess(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	usecase := &taskUsecase{
		taskRepository: mockTaskRepo,
		contextTimeout: 5 * time.Second,
	}

	task := domain.Task{ID: "1", Title: "task1"}

	mockTaskRepo.On("GetOne", mock.Anything, "1").Return(task, nil).Once()

	result, err := usecase.GetOne(context.Background(), "1")

	assert.NoError(t, err)
	assert.Equal(t, task, result)
	mockTaskRepo.AssertExpectations(t)
}

// TestGetOneFails should return an error when fetching a task by ID fails
func TestGetOneFails(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	usecase := &taskUsecase{
		taskRepository: mockTaskRepo,
		contextTimeout: 5 * time.Second,
	}

	mockTaskRepo.On("GetOne", mock.Anything, "1").Return(domain.Task{}, errors.New("not found")).Once()

	result, err := usecase.GetOne(context.Background(), "1")

	assert.Error(t, err)
	assert.Equal(t, domain.Task{}, result)
	mockTaskRepo.AssertExpectations(t)
}

// TestRemoveSuccess should remove a task by ID successfully
func TestRemoveSuccess(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	usecase := &taskUsecase{
		taskRepository: mockTaskRepo,
		contextTimeout: 5 * time.Second,
	}

	mockTaskRepo.On("Remove", mock.Anything, "1").Return(nil).Once()

	err := usecase.Remove(context.Background(), "1")

	assert.NoError(t, err)
	mockTaskRepo.AssertExpectations(t)
}

// TestRemoveFails should return an error when removing a task by ID fails
func TestRemoveFails(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	usecase := &taskUsecase{
		taskRepository: mockTaskRepo,
		contextTimeout: 5 * time.Second,
	}

	mockTaskRepo.On("Remove", mock.Anything, "1").Return(errors.New("delete error")).Once()

	err := usecase.Remove(context.Background(), "1")

	assert.Error(t, err)
	mockTaskRepo.AssertExpectations(t)
}

// TestUpdateSuccess should update a task successfully
func TestUpdateSuccess(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	usecase := &taskUsecase{
		taskRepository: mockTaskRepo,
		contextTimeout: 5 * time.Second,
	}

	task := &domain.Task{
		ID:          "1",
		Title:       "Updated task",
		Description: "Updated desc",
		DueDate:     time.Now(),
		Status:      "completed",
	}

	mockTaskRepo.On("Update", mock.Anything, "1", task).Return(*task, nil).Once()

	result, err := usecase.Update(context.Background(), "1", task)

	assert.NoError(t, err)
	assert.Equal(t, *task, result)
	mockTaskRepo.AssertExpectations(t)
}

// TestUpdateFails should return an error when updating a task fails
func TestUpdateFails(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	usecase := &taskUsecase{
		taskRepository: mockTaskRepo,
		contextTimeout: 5 * time.Second,
	}

	task := &domain.Task{
		ID:          "1",
		Title:       "Updated task",
		Description: "Updated desc",
		DueDate:     time.Now(),
		Status:      "completed",
	}

	mockTaskRepo.On("Update", mock.Anything, "1", task).Return(domain.Task{}, errors.New("update error")).Once()

	result, err := usecase.Update(context.Background(), "1", task)

	assert.Error(t, err)
	assert.Equal(t, domain.Task{}, result)
	mockTaskRepo.AssertExpectations(t)
}
