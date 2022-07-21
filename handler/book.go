package handler

import (
	"BookManagementApp/CacheDatabase"
	"BookManagementApp/helper"
	"BookManagementApp/model"
	"encoding/json"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := model.Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		helper.WriteToLogFile(err, 91024760, "Bad Body")

		return
	}
	if book.Name == "" || book.Author == "" || book.Pages == 0 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		helper.WriteToLogFile(err, 91024760, "Bad Body")
		return
	}
	_, err = json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		helper.WriteToLogFile(err, 98094760, "Error in marshaling the book model")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	CacheDatabase.Books = append(CacheDatabase.Books, book)
	w.Write([]byte("Book created successfully"))
}
