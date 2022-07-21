package handler

import (
	"BookManagementApp/CacheDatabase"
	"BookManagementApp/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBook(t *testing.T) {
	book := model.Book{
		ID:               0,
		Name:             "",
		Author:           "",
		Pages:            0,
		PercentageOfRead: 0,
	}
	body, err := json.Marshal(book)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/createBook", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBook)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `Book created successfully`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	expectedCount := 1
	if len(CacheDatabase.Books) != expectedCount {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
