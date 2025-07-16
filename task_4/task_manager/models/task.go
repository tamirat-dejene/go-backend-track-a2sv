package models

import (
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

const (
	ONGOING TaskStatus = "ongoing"
	DONE    TaskStatus = "done"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description" binding:"required"`
	DueDate     time.Time  `json:"due_date" binding:"required"`
	Status      TaskStatus `json:"status" binding:"required"`
}
