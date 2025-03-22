package handlers

import (
	"encoding/json"

	"my-book-api/models"
	"my-book-api/utils"
	"net/http"
	

	"github.com/google/uuid"
	
)

func GetBooks (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type","application/json")
	json.NewEncoder(w).Encode(utils.Books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.BookID = uuid.New().String()
	utils.Books = append(utils.Books, book)
	_ = utils.SaveBooks()
	json.NewEncoder(w).Encode(book)
}