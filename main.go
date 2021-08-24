package main

import (
	"book-list/controllers"
	"book-list/driver"
	"book-list/models"
	bookRepository "book-list/repository/book"
	"book-list/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var books []models.Book

var db *sql.DB

func init() {
	err := gotenv.Load()
	utils.LogFatal(err)
}

func main() {

	db = driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()
	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")
	router.HandleFunc("/health", healthCheckHandler).Methods("GET")

	fmt.Println("Server is running on port http://localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
}


func removeBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Exec(`DELETE FROM book WHERE "ID" = $1`, params["id"])
	utils.LogFatal(err)

	rowsDeleted, err := result.RowsAffected()

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(rowsDeleted)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var cError models.Error

	err := json.NewDecoder(r.Body).Decode(&book)
	utils.LogFatal(err)

	err = bookRepository.BookRepository{}.AddBook(db, book)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, cError)
	}
	utils.SendSuccess(w, nil, http.StatusCreated)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	utils.LogFatal(err)

	updateSql := `
		UPDATE book SET "Title"=$1, "Author"=$2, "Year"=$3 
		WHERE "ID"=$4 
		RETURNING "ID";
	`
	result, err := db.Exec(updateSql, &book.Title, &book.Author, &book.Year, &book.ID)
	rowsUpdated, err := result.RowsAffected()
	utils.LogFatal(err)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rowsUpdated)
}
