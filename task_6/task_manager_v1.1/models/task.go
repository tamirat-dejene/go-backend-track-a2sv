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
	ID          string     `json:"id" binding:"uuid" bson:"_id"`
	Title       string     `json:"title" binding:"required" bson:"title"`
	Description string     `json:"description" binding:"required" bson:"description"`
	DueDate     time.Time  `json:"due_date" binding:"required" bson:"due_date"`
	Status      TaskStatus `json:"status" binding:"required" bson:"status"`
}
