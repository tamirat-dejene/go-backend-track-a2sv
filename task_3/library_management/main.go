package main

import (
	"fmt"
	"os"
	"task3/lm/controllers"
	"task3/lm/models"
	"task3/lm/services"
)

func main() {
	// This is the entry point of the application.
	fmt.Println("Library Management System is running...")
	lib := &services.Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}

	args := os.Args[1:]
	fmt.Println("Librarians:", args)
	controllers.Operations(lib)
}
