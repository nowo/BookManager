package main

import (
	"BookManagementApp/handler"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}
	fmt.Println("Server running on port:", port)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/getUsers", handler.GetUsers)
	http.HandleFunc("/createUser", handler.CreateUser)
	http.HandleFunc("/uploadBookToUser", handler.UploadBookToUser)
	http.HandleFunc("/createBook", handler.CreateBook)
	http.ListenAndServe(":"+port, nil)
}
