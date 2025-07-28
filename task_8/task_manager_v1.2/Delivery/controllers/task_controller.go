package controllers

import (
	"fmt"
	"net/http"
	domain "t8/taskmanager/Domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

func (tc *TaskController) GetAll(ctx *gin.Context) {
	tasks, err := tc.TaskUsecase.GetAll(ctx)

	if err != nil {
		ctx.IndentedJSON(500, domain.ErrorResponse{Message: "Failed to retrieve tasks", Error: err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetOne(ctx *gin.Context) {
	var id domain.TaskIDUri
	if err := ctx.ShouldBindUri(&id); err != nil {
		ctx.JSON(400, domain.ErrorResponse{Message: "Invalid UUID format for task ID", Error: err.Error()})
		return
	}

	task, err := tc.TaskUsecase.GetOne(ctx, id.ID)

	if err != nil {
		ctx.JSON(404, domain.ErrorResponse{Message: "Task not found", Error: err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, task)
}

func (tc *TaskController) Remove(ctx *gin.Context) {
	var id domain.TaskIDUri
	if err := ctx.ShouldBindUri(&id); err != nil {
		ctx.JSON(400, domain.ErrorResponse{Message: "Invalid UUID format for task ID", Error: err.Error()})
		return
	}

	err := tc.TaskUsecase.Remove(ctx, id.ID)
	if err != nil {
		ctx.JSON(404, domain.ErrorResponse{Message: "Task not found", Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, domain.SuccessResponse{Message: fmt.Sprintf("task with id %s removed", id.ID)})
}

func (tc *TaskController) Create(ctx *gin.Context) {
	var ctr domain.CreateTask

	if err := ctx.ShouldBind(&ctr); err != nil {
		ctx.IndentedJSON(400, domain.ErrorResponse{Message: "Validation failed: please check your input fields", Error: err.Error()})
		return
	}

	var task = domain.Task{
		ID:          uuid.New().String(),
		Title:       ctr.Title,
		Description: ctr.Description,
		DueDate:     ctr.DueDate,
		Status:      ctr.Status,
	}

	if err := tc.TaskUsecase.Add(ctx, &task); err != nil {
		ctx.JSON(500, domain.ErrorResponse{Message: "Failed to create task", Error: err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, task)
}

func (tc *TaskController) Update(ctx *gin.Context) {
	var utr domain.Task
	var id domain.TaskIDUri

	if err := ctx.ShouldBindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid UUID format for task ID", Error: err.Error()})
		return
	}

	if err := ctx.ShouldBind(&utr); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Validation failed: please check your input fields", Error: err.Error()})
		return
	}

	_, err := tc.TaskUsecase.Update(ctx, id.ID, &utr)

	if err != nil {
		ctx.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Failed to update the task", Error: err.Error()})
		return
	}

	// ctx.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
	ctx.Redirect(http.StatusFound, "/api/tasks")
}
