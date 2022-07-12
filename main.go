package main

import (
	"BookManagementApp/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/get", handler.GetUser)
	http.HandleFunc("/create", handler.CreateUser)
	http.ListenAndServe(":8080", nil)

}
