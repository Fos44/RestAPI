package controllers

import(
	"books-list/models"
	"books-list/repository/book"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Controller struct{}

var books []models.Book

func logFatal(err error){
	if err != nil{
		log.Fatal(err)
	}
}

func (c Controller) GetBooks(db *sql.DB) http.HandleFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}

		bookRepo := bookrepository.Bookrepository{}
		books = bookRepo.GetBooks(db, book, books)

		json.NewEncoder(w).Encode(books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandleFunc{
 return func (w http.ResponseWriter, r *http.Request) {
	var book models.Book
	params := mux.Vars(r)

	books = []model.Book{}
	bookRepo := bookrepository.Bookrepository{}

	id, err := strconv.Atoi(params["id"])
	logFatal(err)

	book = bookRepo.GetBook(db, book, id)

	json.NewEncoder(w).Encode(book)
}
}

func (c Controller) AddBook(db *sql.DB) http.HandleFunc{
	return func (w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)

	bookRepo := bookrepository.Bookrepository{}
	bookID = bookRepo.AddBook(db, book)

	json.NewEncoder(w).Encode(bookID)

}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandleFunc{
	return func (w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	bookRepo := bookrepository.Bookrepository{}
	rowsUpdated := bookRepo.UpdateBook(db, book)

	json.NewEncoder(w).Encode(rowsUpdated)

}
}

func (c Controller) RemoveBook(*sql.DB) http.HandleFunc{
	return func (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookRepo := bookrepository.Bookrepository{}

	id, err := strconv.Atoi(params["id"])
	logFatal(err)

	rowsDeleted := bookRepo.RemoveBook(db, id)

	json.NewEncoder(w).Encode(rowsDeleted)
}
}