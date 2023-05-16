package main

import (
	"bookservice/src/container"
	"bookservice/src/mlog" // Add this import
	"bookservice/src/router"
	"log"
	"net/http"

	"fmt"
	"os"

	// "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	mlog.InitLogger() // Initialize the logger

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSslmode := os.Getenv("DB_SSLMODE")
	dbHost := os.Getenv("DB_HOST")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s", dbUser, dbPassword, dbName, dbSslmode, dbHost)

	bookRepoAdapter, err := container.InitBookRepository(connectionString)

	if err != nil {
		mlog.Error.Fatal(err) // Use the mlog.Error logger
	}

	router := router.NewRouter(bookRepoAdapter)

	mlog.Info.Println("Starting server at :8000") // Use the mlog.Info logger
	log.Fatal(http.ListenAndServe(":8000", router))
}
