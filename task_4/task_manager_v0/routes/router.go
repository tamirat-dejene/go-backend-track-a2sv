package routes

import (
	"net/http"
	"t4/taskmanager/controllers"

	"github.com/gin-gonic/gin"
)

/*
Package routes provides routing setup for the task manager application.
It defines API endpoints for task operations such as listing, retrieving,
creating, updating, and deleting tasks using the Gin web framework.
*/
func SetUpRouter(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) { ctx.Redirect(http.StatusPermanentRedirect, "/api") })

	var api = router.Group("/api")

	api.GET("/", controllers.Home)
	api.GET("/tasks", controllers.GetTasks)
	api.GET("/tasks/:id", controllers.GetTask)
	api.POST("/tasks/new", controllers.CreateTask)
	api.PUT("/tasks/edit/:id", controllers.UpdateTask)
	api.DELETE("/tasks/delete/:id", controllers.RemoveTask)
}
