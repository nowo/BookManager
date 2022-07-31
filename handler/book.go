package handler

import (
	"BookManagementApp/CacheDatabase"
	"BookManagementApp/model"
	"encoding/json"
	"log"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := model.Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	if book.Name == "" || book.Author == "" || book.Pages == 0 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	_, err = json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	CacheDatabase.Books = append(CacheDatabase.Books, book)
	w.Write([]byte("Book created successfully"))
}
