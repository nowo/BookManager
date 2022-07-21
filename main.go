package main

import (
	"BookManagementApp/handler"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server started at port: 8080")
	http.HandleFunc("/getUsers", handler.GetUsers)
	http.HandleFunc("/createUser", handler.CreateUser)
	http.HandleFunc("/uploadBookToUser", handler.UploadBookToUser)
	http.HandleFunc("/createBook", handler.CreateBook)
	http.ListenAndServe(":8080", nil)
}
