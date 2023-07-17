package main

import (
	"bookservice/src/cmd/proto"
	"bookservice/src/container"
	"bookservice/src/domain/core/model"
	"bookservice/src/domain/core/ports/outgoing"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedBookServiceServer
	bookRepository outgoing.BookRepository
}

func (s *server) CreateBook(ctx context.Context, in *proto.Book) (*proto.Book, error) {
	// Convert the proto.Book to domain.core.model.Book
	bookToCreate := model.Book{
		Title:  in.GetTitle(),
		Author: in.GetAuthor(),
		ISBN:   in.GetIsbn(),
	}

	createdBook, err := s.bookRepository.CreateBook(bookToCreate)
	if err != nil {
		return nil, err
	}

	// Convert createdBook (type model.Book) back to proto.Book before returning
	return &proto.Book{
		Id:     int32(createdBook.ID),
		Title:  createdBook.Title,
		Author: createdBook.Author,
		Isbn:   createdBook.ISBN,
	}, nil
}

func (s *server) GetBookById(ctx context.Context, id *proto.BookId) (*proto.Book, error) {
	book, err := s.bookRepository.GetBookById(int(id.GetId()))
	if err != nil {
		return nil, err
	}

	return &proto.Book{
		Id:     int32(book.ID),
		Title:  book.Title,
		Author: book.Author,
		Isbn:   book.ISBN,
	}, nil
}

func (s *server) UpdateBook(ctx context.Context, book *proto.Book) (*proto.Book, error) {
	bookToUpdate := model.Book{
		ID:     int(book.GetId()),
		Title:  book.GetTitle(),
		Author: book.GetAuthor(),
		ISBN:   book.GetIsbn(),
	}

	updatedBook, err := s.bookRepository.UpdateBook(bookToUpdate.ID, bookToUpdate)
	if err != nil {
		return nil, err
	}

	return &proto.Book{
		Id:     int32(updatedBook.ID),
		Title:  updatedBook.Title,
		Author: updatedBook.Author,
		Isbn:   updatedBook.ISBN,
	}, nil
}

func (s *server) GetAllBooks(ctx context.Context, empty *proto.Empty) (*proto.BookList, error) {

	books, err := s.bookRepository.GetAllBooks()
	if err != nil {
		return nil, err
	}

	bookList := &proto.BookList{}

	for _, book := range books {
		bookList.Books = append(bookList.Books, &proto.Book{
			Id:     int32(book.ID),
			Title:  book.Title,
			Author: book.Author,
			Isbn:   book.ISBN,
		})
	}

	return bookList, nil

}
func (s *server) DeleteBook(ctx context.Context, id *proto.BookId) (*proto.Book, error) {
	err := s.bookRepository.DeleteBook(int(id.GetId()))
	if err != nil {
		return nil, err
	}

	// We can return a default Book object here, or potentially fetch the Book that was deleted
	// prior to deletion and return that.
	return &proto.Book{}, nil
}

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSslmode := os.Getenv("DB_SSLMODE")
	dbHost := os.Getenv("DB_HOST")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s", dbUser, dbPassword, dbName, dbSslmode, dbHost)

	bookRepo, err := container.InitBookRepository(connectionString)
	if err != nil {
		log.Fatalf("Failed to initialize Book Repository: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterBookServiceServer(s, &server{bookRepository: bookRepo})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
