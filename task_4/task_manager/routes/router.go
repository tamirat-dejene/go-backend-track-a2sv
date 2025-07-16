package routes

import (
	"t4/taskmanager/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(router *gin.Engine) {
	router.GET("/", controllers.Home)
	var api = router.Group("/api")

	api.GET("/tasks", controllers.GetTasks)
	api.GET("/tasks/:id", controllers.GetTask)
	api.POST("/task", controllers.CreateTask)
	api.PUT("/task/:id", controllers.UpdateTask)
	api.DELETE("/task/:id", controllers.RemoveTask)
}
