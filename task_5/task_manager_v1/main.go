package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"t4/taskmanager/data"
	"t4/taskmanager/routes"
	"time"

	"github.com/gin-gonic/gin"
)

// close gracefully shuts down the HTTP server when an interrupt signal is received.
func close(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func main() {
	// Initialize MongoDB
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}
	if err := data.InitMongo(mongoURI); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	log.Println("MongoDB connected.:27017")

	// Set up Gin Router
	router := gin.Default()
	routes.SetUpRouter(router)

	srv := &http.Server{
		Addr:    ":1337",
		Handler: router.Handler(),
	}
	// Start HTTP Server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	close(srv)
}
