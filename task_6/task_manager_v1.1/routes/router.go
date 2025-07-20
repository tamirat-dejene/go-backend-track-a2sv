package routes

import (
	"net/http"
	"t4/taskmanager/controllers"
	"t4/taskmanager/middlewares"

	"github.com/gin-gonic/gin"
)

/*
Package routes provides routing setup for the task manager application.
It defines API endpoints for task operations such as listing, retrieving,
creating, updating, and deleting tasks using the Gin web framework.
It also organizes authentication routes under "/api/auth" and task-related routes under "/api".
*/

func SetUpRouter(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) { ctx.Redirect(http.StatusPermanentRedirect, "/api") })

	var api = router.Group("/api")
	var auth = router.Group("/auth")

	auth.POST("/login", controllers.LoginUser)
	auth.POST("/register", controllers.RegisterUser)
	auth.POST("/refresh", controllers.RefreshUserToken)
	auth.POST("/logout", controllers.LogoutUser)
	auth.DELETE("/delete/:id", controllers.DeleteUser)
	auth.GET("/users", middlewares.AuthMiddleWare(), controllers.GetUsers)

	api.Use(middlewares.AuthMiddleWare())

	api.GET("/", controllers.Home)
	api.GET("/tasks", controllers.GetTasks)
	api.GET("/tasks/:id", controllers.GetTask)
	api.POST("/tasks/new", controllers.CreateTask)
	api.PUT("/tasks/edit/:id", controllers.UpdateTask)
	api.DELETE("/tasks/delete/:id", controllers.RemoveTask)
}
