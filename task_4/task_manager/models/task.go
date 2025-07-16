package models

import (
	"time"
)

type TaskStatus string

const (
	ONGOING TaskStatus = "ongoing"
	DONE    TaskStatus = "done"
)

type Task struct {
	ID          string     `json:"id" binding:"uuid"`
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description" binding:"required"`
	DueDate     time.Time  `json:"due_date" binding:"required"`
	Status      TaskStatus `json:"status" binding:"required"`
}
