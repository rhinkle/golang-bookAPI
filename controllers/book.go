package controllers

import (
	"book-list/models"
	bookRepository "book-list/repository/book"
	"book-list/utils"
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

type Controller struct {}

var books []models.Book

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var cError models.Error

		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}

		books, err := bookRepo.GetBooks(db, book, books)
		if err != nil {
			cError.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, cError)
		}

		utils.SendSuccess(w, books, 200)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var cError models.Error

		params := mux.Vars(r)
		bookRepo := bookRepository.BookRepository{}

		book, err := bookRepo.GetBook(db, book, params["id"])

		if err != nil {
			cError.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, cError)
		}

		utils.SendSuccess(w, book, 200)
	}
}