package core

import (
	"bookservice/src/domain/core/model"
	"bookservice/src/domain/core/ports/outgoing"
)

type BookFacade struct {
	Repository outgoing.BookRepository
}

func (f *BookFacade) Create(book model.Book) (model.Book, error) {
	return f.Repository.CreateBook(book)
}

func (f *BookFacade) GetAll() ([]model.Book, error) {
	return f.Repository.GetAllBooks()
}

func (f *BookFacade) GetByID(id int) (model.Book, error) {
	return f.Repository.GetBookById(id)
}

func (f *BookFacade) Update(id int, book model.Book) (model.Book, error) {
	return f.Repository.UpdateBook(id, book)
}

func (f *BookFacade) Delete(id int) error {
	return f.Repository.DeleteBook(id)
}
