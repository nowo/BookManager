package CacheDatabase

import (
	"BookManagementApp/model"
	"errors"
)

var Users []model.User
var Books []model.Book

func GetUsersIndex(id int) (int, error) {
	for index, user := range Users {
		if user.ID == id {
			return index, nil
		}
	}
	return 0, errors.New("User not found")
}

func GetBook(id int) (model.Book, error) {
	for _, book := range Books {
		if book.ID == id {
			return book, nil
		}
	}
	return model.Book{}, errors.New("Book not found")
}
