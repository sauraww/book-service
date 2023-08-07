package sqlstore

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SQLClient interface {
	InsertBook(book Book) (uint, error)
	UpsertBook(book Book) (uint, error)
	UpdateBook(book Book) error
}

type Book struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"not null"`
	Author string `gorm:"not null"`
	ISBN   string `gorm:"unique;not null"`
}

type GORMClient struct {
	db *gorm.DB
}

func NewGORMClient(dsn string) (*GORMClient, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &GORMClient{db: db}, nil
}

func (g *GORMClient) InsertBook(book Book) (uint, error) {
	result := g.db.Create(&book)
	return book.ID, result.Error
}

func (g *GORMClient) UpsertBook(book Book) (uint, error) {
	// Check if the book with the same ISBN exists
	var existingBook Book
	if err := g.db.Where("isbn = ?", book.ISBN).First(&existingBook).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Book does not exist, so create a new record
			return g.InsertBook(book)
		}
		return 0, err
	}

	// Book exists, update the record
	book.ID = existingBook.ID
	return book.ID, g.UpdateBook(book)
}

func (g *GORMClient) UpdateBook(book Book) error {
	result := g.db.Save(&book)
	return result.Error
}

func main() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	client, err := NewGORMClient(dsn)
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	client.db.AutoMigrate(&Book{})

	// Example usage
	book := Book{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", ISBN: "1234567890"}
	id, err := client.InsertBook(book)
	if err != nil {
		panic(err)
	}
	book.ID = id
	book.Author = "Fitzgerald"
	err = client.UpdateBook(book)
	if err != nil {
		panic(err)
	}
}
