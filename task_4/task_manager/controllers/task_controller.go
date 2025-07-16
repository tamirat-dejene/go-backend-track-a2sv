package controllers

/*
Package controllers provides HTTP handler functions for managing tasks.
It includes endpoints for listing, retrieving, creating, updating, and deleting tasks.
All handlers use Gin for routing and JSON responses.
*/

import (
	"fmt"
	"net/http"
	"t4/taskmanager/data"
	"t4/taskmanager/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IDUri struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func Home(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message":     "🎉 Welcome to the Task Manager API!",
		"description": "Effortlessly manage your tasks with our simple RESTful interface.",
		"endpoints": []string{
			"GET    /tasks         - List all tasks",
			"GET    /tasks/:id     - Get a specific task",
			"POST   /tasks         - Create a new task",
			"PUT    /tasks/:id     - Update an existing task",
			"DELETE /tasks/:id     - Remove a task",
		},
		"docs":      "Visit /docs for API documentation.",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

func GetTasks(ctx *gin.Context) {
	tasks, err := data.GetTasks()

	if err != nil {
		ctx.IndentedJSON(500, gin.H{"message": "Failed to retrieve tasks", "error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, tasks)
}

func GetTask(ctx *gin.Context) {
	var id IDUri
	if err := ctx.ShouldBindUri(&id); err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid UUID format for task ID", "error": err.Error()})
		return
	}

	task, err := data.GetTask(id.ID)

	if err != nil {
		ctx.JSON(404, gin.H{"message": "Task not found", "error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, task)
}

func RemoveTask(ctx *gin.Context) {
	var id IDUri
	if err := ctx.ShouldBindUri(&id); err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid UUID format for task ID", "error": err.Error()})
		return
	}

	err := data.RemoveTask(id.ID)
	if err != nil {
		ctx.JSON(404, gin.H{"message": "Task not found", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("task with id %s removed", id.ID)})
}

type CreateTaskRequest struct {
	Title       string            `json:"title" binding:"required"`
	Description string            `json:"description" binding:"required"`
	DueDate     time.Time         `json:"due_date" binding:"required"`
	Status      models.TaskStatus `json:"status" binding:"required"`
}

func CreateTask(ctx *gin.Context) {
	var ctr CreateTaskRequest

	if err := ctx.ShouldBind(&ctr); err != nil {
		ctx.IndentedJSON(400, gin.H{"message": "Validation failed: please check your input fields", "error": err.Error()})
		return
	}

	var task = models.Task{
		ID:          uuid.New().String(),
		Title:       ctr.Title,
		Description: ctr.Description,
		DueDate:     ctr.DueDate,
		Status:      ctr.Status,
	}

	if err := data.AddTask(&task); err != nil {
		ctx.JSON(500, gin.H{"message": "Failed to create task", "error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, task)
}

func UpdateTask(ctx *gin.Context) {
	var utr models.Task
	var id IDUri

	if err := ctx.ShouldBindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid UUID format for task ID", "error": err.Error()})
		return
	}

	if err := ctx.ShouldBind(&utr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed: please check your input fields", "error": err.Error()})
		return
	}

	_, err := data.UpdateTask(id.ID, utr)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Failed to update the task", "error": err.Error()})
		return
	}

	// ctx.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
	ctx.Redirect(http.StatusFound, "/api/tasks")
}
