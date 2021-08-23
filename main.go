package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
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
	log.Println(pgUrl)

	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")
	router.HandleFunc("/health", healthCheckHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

}

func removeBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

}

func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

}
