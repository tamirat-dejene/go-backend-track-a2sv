package controllers

import (
	"fmt"
	"task3/lm/models"
	"task3/lm/services"
)

func CreateBook() models.Book {
	var book models.Book
	fmt.Println("Enter Book ID:")
	fmt.Scanln(&book.ID)
	fmt.Println("Enter Book Title:")
	fmt.Scanln(&book.Title)
	fmt.Println("Enter Book Author:")
	fmt.Scanln(&book.Author)
	book.Status = models.Available
	return book
}

func CreateMember() models.Member {
	var member models.Member
	fmt.Println("Enter Member ID:")
	fmt.Scanln(&member.ID)
	fmt.Println("Enter Member Name:")
	fmt.Scanln(&member.Name)
	member.BorrowedBooks = []models.Book{}
	return member
}

func Operations(lib *services.Library) {
	fmt.Println("\nConsole-Based Library Management System")
	fmt.Println("---------------------------------------")
	fmt.Println("Available Operations:")
	fmt.Println("1. Add Book")
	fmt.Println("2. Add Member")
	fmt.Println("3. Remove Book")
	fmt.Println("4. Borrow Book")
	fmt.Println("5. Return Book")
	fmt.Println("6. List Available Books")
	fmt.Println("7. List Borrowed Books")
	fmt.Println("8. Exit")
	fmt.Println("-------------------------")

	fmt.Println("Please enter a command number:")
	var command int
	fmt.Scanln(&command)

	switch command {
	case 1:
		newBook := CreateBook()
		lib.AddBook(newBook)
		fmt.Println("Book added successfully.")
	case 2:
		newMember := CreateMember()
		lib.AddMember(newMember)
		fmt.Println("Member added successfully.")
	case 3:
		var bookID int
		fmt.Println("Enter Book ID to remove:")
		fmt.Scanln(&bookID)
		lib.RemoveBook(bookID)
		fmt.Println("Book removed successfully.")
	case 4:
		var bookID, memberID int
		fmt.Println("Enter Book ID to borrow:")
		fmt.Scanln(&bookID)
		fmt.Println("Enter Member ID:")
		fmt.Scanln(&memberID)
		err := lib.BorrowBook(bookID, memberID)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Book borrowed successfully.")
		}
	case 5:
		var bookID, memberID int
		fmt.Println("Enter Book ID to return:")
		fmt.Scanln(&bookID)
		fmt.Println("Enter Member ID:")
		fmt.Scanln(&memberID)
		err := lib.ReturnBook(bookID, memberID)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Book returned successfully.")
		}
	case 6:
		availableBooks := lib.ListAvailableBooks()
		fmt.Println("Available Books:")
		for _, book := range availableBooks {
			fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	case 7:
		var memberID int
		fmt.Println("Enter Member ID to list borrowed books:")
		fmt.Scanln(&memberID)
		borrowedBooks := lib.ListBorrowedBooks(memberID)
		fmt.Println("Borrowed Books:")
		for _, book := range borrowedBooks {
			fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	case 8:
		fmt.Println("Exiting the system...")
		return
	default:
		fmt.Println("Invalid command. Please try again.")
	}
	fmt.Println("Returning to main menu...")
	Operations(lib)
}
