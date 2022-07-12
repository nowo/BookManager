package main

import (
	"BookManagementApp/handler"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server  started at port 8080")
	http.HandleFunc("/get", handler.GetUser)
	http.HandleFunc("/create", handler.CreateUser)
	http.ListenAndServe(":8080", nil)

}
