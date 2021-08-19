package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

type HealthCheck struct {
	Status string    `json:status`
	Time   time.Time `json:time`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books,
		Book{ID: 1, Title: "Go Lang Pointers", Author: "MR. Go Lang", Year: "2010"},
		Book{ID: 2, Title: "Go Lang History", Author: "MR. Go Lang", Year: "2009"},
	)

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")
	router.HandleFunc("/health", healthCheckHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	var resp = HealthCheck{
		Status: "Up",
		Time:   time.Now(),
	}
	json.NewEncoder(w).Encode(resp)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("You got books")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])
	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("You updated a book")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("You got books")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("You got books")
}
