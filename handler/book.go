package handler

import (
	"BookManagementApp/CacheDatabase"
	"BookManagementApp/model"
	"encoding/json"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := model.Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	CacheDatabase.Books = append(CacheDatabase.Books, book)
	w.Write([]byte("Book created successfully"))
}