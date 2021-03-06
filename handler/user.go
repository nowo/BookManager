package handler

import (
	"BookManagementApp/CacheDatabase"
	"BookManagementApp/helper"
	"BookManagementApp/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if CacheDatabase.Users == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No users found"))
		helper.WriteToLogFile(nil, 41233123, "GetUsers request called when there is no user")
		return
	}
	response, err := json.Marshal(CacheDatabase.Users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		helper.WriteToLogFile(nil, 34392714, "Error in marshaling the CacheDatabase User slice") // I made random number to make it unique
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
		helper.WriteToLogFile(err, 80219016, "Error in decoding the request body")
		return
	}
	if user.Name == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		helper.WriteToLogFile(err, 80219016, "Error in decoding the request body")
		return
	}
	_, err = json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		helper.WriteToLogFile(err, 26193790, "Error in marshaling the user model")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	CacheDatabase.Users = append(CacheDatabase.Users, user)
	w.Write([]byte("User created successfully"))
}

func UploadBookToUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	bookID := r.URL.Query().Get("bookID")

	if userID == "" || bookID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request"))

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		helper.WriteToLogFile(err, 87358402, "Error in converting userID to int")
		return
	}
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		helper.WriteToLogFile(err, 89594183, "Error in converting bookID to int")
		return
	}
	for _, user := range CacheDatabase.Users {
		if userIDInt == user.ID {
			for _, book := range CacheDatabase.Books {
				if bookIDInt == book.ID {
					user.Books = append(user.Books, book)
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("Book added successfully to user"))
					return
				}
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("User couldnt added successfully to user"))
}
