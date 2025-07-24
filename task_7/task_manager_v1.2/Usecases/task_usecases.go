package usecases

import (
	"context"
	domain "t7/taskmanager/Domain"
	"time"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

// Add implements domain.TaskUsecase.
func (t *taskUsecase) Add(ctx context.Context, task *domain.Task) error {
	c, cancel := context.WithTimeout(ctx, t.contextTimeout)
	defer cancel()

	return t.taskRepository.Add(c, task)
}

// GetAll implements domain.TaskUsecase.
func (t *taskUsecase) GetAll(ctx context.Context) ([]domain.Task, error) {
	c, cancel := context.WithTimeout(ctx, t.contextTimeout)
	defer cancel()

	return t.taskRepository.GetAll(c)
}

// GetOne implements domain.TaskUsecase.
func (t *taskUsecase) GetOne(ctx context.Context, id string) (domain.Task, error) {
	c, cancel := context.WithTimeout(ctx, t.contextTimeout)
	defer cancel()

	return t.taskRepository.GetOne(c, id)
}

// Remove implements domain.TaskUsecase.
func (t *taskUsecase) Remove(ctx context.Context, id string) error {
	c, cancel := context.WithTimeout(ctx, t.contextTimeout)
	defer cancel()

	return t.taskRepository.Remove(c, id)
}

// Update implements domain.TaskUsecase.
func (t *taskUsecase) Update(ctx context.Context, id string, task *domain.Task) (domain.Task, error) {
	c, cancel := context.WithTimeout(ctx, t.contextTimeout)
	defer cancel()

	return t.taskRepository.Update(c, id, task)
}

func NewTaskUsecase(taskRepo domain.TaskRepository, timeout time.Duration) domain.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepo,
		contextTimeout: timeout,
	}
}
