package container

import (
	"bookservice/src/domain/core/infrastructure"
	"bookservice/src/domain/core/infrastructure/adapters"
	"bookservice/src/domain/core/ports/outgoing"
)

func InitBookRepository(dbConfig string) (outgoing.BookRepository, error) {
	db, err := infrastructure.NewDBConnection(dbConfig)
	if err != nil {
		return nil, err
	}

	bookRepoImpl := infrastructure.NewBookRepositoryImpl(db)
	bookRepoAdapter := adapters.NewBookRepositoryAdapter(bookRepoImpl)

	return bookRepoAdapter, nil
}
