package infrastructure

import (
	"bookservice/src/domain/core/model"
	"database/sql"
	"fmt"
)

type BookRepositoryImpl struct {
	db *sql.DB
}

func NewBookRepositoryImpl(db *sql.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{db: db}
}

func NewDBConnection(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (r *BookRepositoryImpl) Create(book model.Book) (model.Book, error) {
	sqlStatement := `INSERT INTO books (title, author, isbn) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(sqlStatement, book.Title, book.Author, book.ISBN).Scan(&book.ID)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func (r *BookRepositoryImpl) GetAll() ([]model.Book, error) {
	rows, err := r.db.Query("SELECT id, title, author, isbn FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *BookRepositoryImpl) GetByID(id int) (model.Book, error) {
	var book model.Book
	sqlStatement := "SELECT id, title, author, isbn FROM books WHERE id=$1"
	row := r.db.QueryRow(sqlStatement, id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN)
	if err == sql.ErrNoRows {
		return model.Book{}, fmt.Errorf("Book not found")
	} else if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func (r *BookRepositoryImpl) Update(id int, book model.Book) (model.Book, error) {
	sqlStatement := `UPDATE books SET title=$1, author=$2, isbn=$3 WHERE id=$4`
	result, err := r.db.Exec(sqlStatement, book.Title, book.Author, book.ISBN, id)
	if err != nil {
		return model.Book{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.Book{}, err
	}

	if rowsAffected == 0 {
		return model.Book{}, fmt.Errorf("Book not found")
	}

	return book, nil
}

func (r *BookRepositoryImpl) Delete(id int) error {
	sqlStatement := "DELETE FROM books WHERE id=$1"
	result, err := r.db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Book not found")
	}

	return nil
}
