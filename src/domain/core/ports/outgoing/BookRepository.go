package outgoing

import (
	"bookservice/src/domain/core/model"
)

type BookRepository interface {
	CreateBook(book model.Book) (model.Book, error)
	GetAllBooks() ([]model.Book, error)
	GetBookById(id int) (model.Book, error)
	UpdateBook(id int, book model.Book) (model.Book, error)
	DeleteBook(id int) error
}
