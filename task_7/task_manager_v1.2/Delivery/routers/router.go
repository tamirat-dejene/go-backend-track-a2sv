package routers

import (
	"net/http"
	"t7/taskmanager/Delivery/bootstrap"
	"t7/taskmanager/Delivery/controllers"
	domain "t7/taskmanager/Domain"
	"t7/taskmanager/Infrustructure/middlewares"
	repositories "t7/taskmanager/Repositories"
	usecases "t7/taskmanager/Usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
Package routes provides routing setup for the task manager application.
It defines API endpoints for task operations such as listing, retrieving,
creating, updating, and deleting tasks using the Gin web framework.
It also organizes authentication routes under "/api/auth" and task-related routes under "/api".
*/

func Setup(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) { ctx.Redirect(http.StatusPermanentRedirect, "/api") })
	// router.GET("/api")

	var auth = router.Group("/auth")
	ur := repositories.NewUserRepository(db, domain.UserCollection)
	uc := &controllers.UserController{
		UserUsecase: usecases.NewUserUsecase(ur, timeout),
		Env:         env,
	}

	auth.POST("/login", uc.Login)
	auth.GET("/logout", uc.Logout)
	auth.GET("/refresh", uc.Refresh)
	auth.POST("/register", uc.Register)
	auth.DELETE("/delete/:id", uc.Delete)
	auth.GET("/users", uc.GetAll)
	var api = router.Group("/api")
	tr := repositories.NewTaskRepository(db, domain.TaskCollection)
	tc := &controllers.TaskController{
		TaskUsecase: usecases.NewTaskUsecase(tr, timeout),
	}
	api.Use(middlewares.AuthMiddleWare(env))
	api.GET("/tasks", tc.GetAll)
	api.GET("/tasks/:id", tc.GetOne)
	api.POST("/tasks/new", tc.Create)
	api.PUT("/tasks/edit/:id", tc.Update)
	api.DELETE("/tasks/delete/:id", tc.Remove)
}
