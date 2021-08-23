package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   string
}

type HealthCheck struct {
	Status string
	Time   time.Time
}

var books []Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	pgUrl, err := pq.ParseURL(os.Getenv("DB_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	// router.HandleFunc("/books", addBook).Methods("POST")
	// router.HandleFunc("/books", updateBook).Methods("PUT")
	// router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")
	router.HandleFunc("/health", healthCheckHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var book Book
	books = []Book{}

	rows, err := db.Query("select * from book")
	logFatal(err)

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	params := mux.Vars(r)

	rows := db.QueryRow(`SELECT * FROM book WHERE "ID" = $1`, params["id"])

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
