package router

import (
	"bookservice/src/domain/core/model"
	"bookservice/src/domain/core/ports/outgoing"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func NewRouter(bookRepoAdapter outgoing.BookRepository) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", createBook(bookRepoAdapter)).Methods("POST")
	router.HandleFunc("/books", getBooks(bookRepoAdapter)).Methods("GET")
	router.HandleFunc("/books/{id}", getBook(bookRepoAdapter)).Methods("GET")
	router.HandleFunc("/books/{id}", updateBook(bookRepoAdapter)).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook(bookRepoAdapter)).Methods("DELETE")

	return router
}

func createBook(bookRepo outgoing.BookRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book model.Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		createdBook, err := bookRepo.CreateBook(book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdBook)
	}
}

func getBooks(bookRepo outgoing.BookRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := bookRepo.GetAllBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(books)
	}
}

func getBook(bookRepo outgoing.BookRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		idStr := params["id"]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
		}

		book, err := bookRepo.GetBookById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(book)
	}
}

func updateBook(bookRepo outgoing.BookRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		idStr := params["id"]

		var book model.Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
		}

		updatedBook, err := bookRepo.UpdateBook(id, book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedBook)
	}
}

func deleteBook(bookRepo outgoing.BookRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		idStr := params["id"]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
		}

		err = bookRepo.DeleteBook(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Book deleted")
	}
}
