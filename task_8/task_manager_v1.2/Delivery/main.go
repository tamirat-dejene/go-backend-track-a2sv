package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"t7/taskmanager/Delivery/bootstrap"
	"t7/taskmanager/Delivery/routers"
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
	app := bootstrap.App()
	env := app.Env
	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	router := gin.Default()
	routers.Setup(env, timeout, db, router)

	srv := &http.Server{
		Addr:    env.ServerAddress,
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


// package main

// import (
// 	"fmt"
// )

// func Calculate(x int) (result int) {
// 	result = x + 2
// 	return result
// }

// type MessageService interface {
// 	SendChargeNotification(int) error
// }
// type SMSService struct{}
// func (ss *SMSService) SendChargeNotification(int) error {
// 	fmt.Println("Sending Production Charge Notification")
// 	return nil
// }

// type MyService struct {
// 	messageService MessageService
// }
// func (ms *MyService) ChargeCustomer(value int) error {
// 	ms.messageService.SendChargeNotification(value)
// 	fmt.Printf("Charging customer for the value of %d\n", value)
// 	return  nil
// }

// func main() {
// 	// fmt.Println("Hello World!")

// 	smsService := SMSService{}
// 	myService := MyService{messageService: &smsService}
// 	myService.ChargeCustomer(100)
// }