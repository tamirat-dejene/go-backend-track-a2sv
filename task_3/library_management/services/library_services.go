package services

import (
	"fmt"
	"task3/lm/models"
)

// LibraryManager defines the interface for library management operations.
type LibraryManager interface {
	AddBook(book models.Book)
	AddMember(member models.Member)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

// Library implements the LibraryManager interface.
type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

// AddBook adds a new book to the library.
// If a book with the same ID already exists, it prints an error message.
func (lib *Library) AddBook(book models.Book) {
	_, ok := lib.Books[book.ID]

	if ok {
		fmt.Printf("Error: Book with id %d already found.\n", book.ID)
		return
	}

	lib.Books[book.ID] = book
	fmt.Printf("Book with id %d added successfully.\n", book.ID)
}

// AddMember adds a new member to the library.
// If a member with the same ID already exists, it prints an error message.
func (lib *Library) AddMember(member models.Member) {
	_, ok := lib.Members[member.ID]

	if ok {
		fmt.Printf("Error: Member with id %d already found.\n", member.ID)
		return
	}

	lib.Members[member.ID] = member
	fmt.Printf("Member with id %d added successfully.\n", member.ID)
}

// RemoveBook removes a book from the library by its ID.
// If the book does not exist, it prints an error message.
// It does not return any value.
// It is assumed that the book ID is unique.
func (lib *Library) RemoveBook(bookId int) {
	_, ok := lib.Books[bookId]

	if !ok {
		println("Error: Book not found")
		return
	}
	delete(lib.Books, bookId)
	fmt.Printf("Book with id %d removed successfully.\n", bookId)
}

// BorrowBook allows a member to borrow a book.
// It checks if the book and member exist, and if the book is available.
// If successful, it updates the book's status and adds it to the member's borrowed books.
func (lib *Library) BorrowBook(bookId, memberId int) error {
	book, ok := lib.Books[bookId]
	if !ok {
		return fmt.Errorf("book with id %d not found", bookId)
	}

	member, ok := lib.Members[memberId]
	if !ok {
		return fmt.Errorf("member with id %d not found", memberId)
	}

	if book.Status == models.Borrowed {
		return fmt.Errorf("book with id %d is already borrowed", bookId)
	}

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	lib.Members[memberId] = member

	book.Status = models.Borrowed
	lib.Books[bookId] = book

	fmt.Printf("Book with id %d borrowed by member %d successfully.\n", bookId, memberId)
	return nil
}

// ReturnBook allows a member to return a borrowed book.
// It checks if the book and member exist, and if the book is currently borrowed.
// If successful, it updates the book's status and removes it from the member's borrowed books.
// It returns an error if the book is not found or if it is already available.
func (lib* Library) ReturnBook(bookId, memberId int) error {
	book, ok := lib.Books[bookId]
	if !ok {
		return fmt.Errorf("book with id %d not found", bookId)
	}

	member, ok := lib.Members[memberId]
	if !ok {
		return fmt.Errorf("member with id %d not found", memberId)
	}

	if book.Status == models.Available {
		return fmt.Errorf("book with id %d is already available", bookId)
	}

	for i, b := range member.BorrowedBooks {
		if b.ID == bookId {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}

	book.Status = models.Available
	lib.Books[bookId] = book
	lib.Members[memberId] = member
	fmt.Printf("Book with id %d returned by member %d successfully.\n", bookId, memberId)

	return nil
}

// ListAvailableBooks returns a slice of all available books in the library.
// It iterates through the library's books and collects those with the status "Available".
// It returns a slice of models.Book.
func (lib *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range lib.Books {
		if book.Status == models.Available {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

// ListBorrowedBooks returns a slice of books borrowed by a specific member.
// It checks if the member exists and returns their borrowed books.
// If the member does not exist, it prints an error message and returns nil.
// It returns a slice of models.Book.
func (lib *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, ok := lib.Members[memberID]
	if !ok {
		fmt.Printf("Error: Member with id %d not found.\n", memberID)
		return nil
	}
	return member.BorrowedBooks
}