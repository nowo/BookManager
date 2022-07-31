package handler

import (
	"BookManagementApp/CacheDatabase"
	"BookManagementApp/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if CacheDatabase.Users == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No users found"))
		return
	}
	response, err := json.Marshal(CacheDatabase.Users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	if user.Name == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	_, err = json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	CacheDatabase.Users = append(CacheDatabase.Users, user)
	w.Write([]byte("User created successfully"))
}

func AddBook(userID int, book model.Book) {
	CacheDatabase.Users[userID].Books = append(CacheDatabase.Users[userID].Books, book)
}

func UploadBookToUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	bookID := r.URL.Query().Get("bookID")

	if userID == "" || bookID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request"))

		return
	}

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	user, err := CacheDatabase.GetUsersIndex(userIDInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Println(err.Error())
		return
	}

	book, err := CacheDatabase.GetBook(bookIDInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Println(err.Error())

		return
	}
	AddBook(user, book)
	fmt.Println(CacheDatabase.Users)
	fmt.Println(CacheDatabase.Books)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book added successfully to user"))
}
