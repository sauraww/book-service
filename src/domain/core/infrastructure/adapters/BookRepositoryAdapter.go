package adapters

import (
	"bookservice/src/domain/core/infrastructure"
	"bookservice/src/domain/core/model"
	"bookservice/src/domain/core/ports/outgoing"
)

type BookRepositoryAdapter struct {
	bookRepo *infrastructure.BookRepositoryImpl // Update this line
}

func NewBookRepositoryAdapter(bookRepo *infrastructure.BookRepositoryImpl) outgoing.BookRepository { // Update this line
	return &BookRepositoryAdapter{bookRepo: bookRepo}
}

func (a *BookRepositoryAdapter) CreateBook(book model.Book) (model.Book, error) { // Update this line
	return a.bookRepo.Create(book)
}

func (a *BookRepositoryAdapter) GetAllBooks() ([]model.Book, error) { // Update this line
	return a.bookRepo.GetAll()
}

func (a *BookRepositoryAdapter) GetBookById(id int) (model.Book, error) { // Update this line
	return a.bookRepo.GetByID(id)
}

func (a *BookRepositoryAdapter) UpdateBook(id int, book model.Book) (model.Book, error) { // Update this line
	return a.bookRepo.Update(id, book)
}

func (a *BookRepositoryAdapter) DeleteBook(id int) error { // Update this line
	return a.bookRepo.Delete(id)
}
