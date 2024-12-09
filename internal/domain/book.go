package domain

import (
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/types"
)

// database interface
type IBookRepo interface {
	GetBooks(bookId uint) []models.Book
	CreateBook(book models.Book) error
	// UpdateBook(book *models.Book) error
	// DeleteBook(bookId uint) error
}

// service interface
type IBookService interface {
	GetBooks(bookId uint) ([]types.BookRequest, error)
	CreateBook(book models.Book) error
	// UpdateBook(book *models.Book) error
	// DeleteBook(bookId uint) error
}
