package bookRepository

import (
	"book-list/models"
	"book-list/utils"
	"database/sql"
)

type BookRepository struct {}

func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) ([]models.Book, error)  {

	rows, err := db.Query("SELECT * FROM book")

	if err != nil {
		return []models.Book{}, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		utils.LogFatal(err)

		books = append(books, book)
	}

	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}

func (b BookRepository) GetBook(db *sql.DB, book models.Book, bookID string) (models.Book, error)  {

	rows := db.QueryRow(`SELECT * FROM book WHERE "ID" = $1`, bookID)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (b BookRepository) AddBook(db *sql.DB, book models.Book) (error)  {
	var bookID int

	addBookSQL := `
		INSERT INTO book ("Title", "Author", "Year")
		VALUES ($1, $2, $3)
		RETURNING "ID";
	`
	err := db.QueryRow(addBookSQL, book.Title, book.Author, book.Year).Scan(&bookID)

	if err != nil {
		return err
	}

	return nil
}



