package controllers


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
	ID uuid.UUID `uri:"id" binding:"required,uuid"`
}

func Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to the Task Manager API!"})
}

func GetTasks(ctx *gin.Context) {
	tasks, err := data.GetTasks()

	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": "internal server error"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, tasks)
}

func GetTask(ctx *gin.Context) {
	var id IDUri
	if err := ctx.ShouldBindUri(&id); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	task, err := data.GetTask(id.ID)

	if err != nil {
		ctx.JSON(404, gin.H{"msg": err.Error()})
	}

	ctx.IndentedJSON(http.StatusOK, task)
}

func RemoveTask(ctx *gin.Context) {
	var id IDUri
	if err := ctx.ShouldBindUri(&id); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	err := data.RemoveTask(id.ID)
	if err != nil {
		ctx.JSON(404, gin.H{"msg": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("task with id %s removed\n", id.ID)})
}

type CreateTaskRequest struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	DueDate     time.Time         `json:"due_date"`
	Status      models.TaskStatus `json:"status"`
}

func CreateTask(ctx *gin.Context) {
	var ctr CreateTaskRequest

	if err := ctx.ShouldBind(&ctr); err != nil {
		ctx.JSON(400, gin.H{"msg": ""})
		return
	}

	var task = models.Task{
		ID: uuid.New(),
		Title: ctr.Title,
		Description: ctr.Description,
		DueDate: ctr.DueDate,
		Status: ctr.Status,
	}
	
	if err := data.AddTask(&task); err != nil {
		ctx.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, task)
}

func UpdateTask(ctx *gin.Context) {
	var utr models.Task

	if err := ctx.ShouldBind(&utr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	_, err := data.UpdateTask(utr.ID, utr)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	ctx.Redirect(http.StatusAccepted, "/tasks")
}