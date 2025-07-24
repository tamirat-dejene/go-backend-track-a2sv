package domain

import (
	"context"
	"time"
)

const (
	TaskCollection = "tasks"
)

type TaskStatus string

const (
	ONGOING TaskStatus = "ongoing"
	DONE    TaskStatus = "done"
)

type TaskIDUri struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type Task struct {
	ID          string     `json:"id" binding:"uuid" bson:"_id"`
	Title       string     `json:"title" binding:"required" bson:"title"`
	Description string     `json:"description" binding:"required" bson:"description"`
	DueDate     time.Time  `json:"due_date" binding:"required" bson:"due_date"`
	Status      TaskStatus `json:"status" binding:"required" bson:"status"`
}

type CreateTask struct {
	Title       string     `json:"title" binding:"required" bson:"title"`
	Description string     `json:"description" binding:"required" bson:"description"`
	DueDate     time.Time  `json:"due_date" binding:"required" bson:"due_date"`
	Status      TaskStatus `json:"status" binding:"required" bson:"status"`
}

type TaskRepository interface {
	Add(ctx context.Context, task *Task) error
	Remove(ctx context.Context, id string) error
	Update(ctx context.Context, id string, task *Task) (Task, error)
	GetAll(ctx context.Context) ([]Task, error)
	GetOne(ctx context.Context, id string) (Task, error)
}

type TaskUsecase interface {
	Add(ctx context.Context, task *Task) error
	Remove(ctx context.Context, id string) error
	Update(ctx context.Context, id string, task *Task) (Task, error)
	GetAll(ctx context.Context) ([]Task, error)
	GetOne(ctx context.Context, id string) (Task, error)
}
