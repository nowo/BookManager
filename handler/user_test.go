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

func TestInitialGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/getUsers", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `No users found`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateUser(t *testing.T) {
	user := model.User{
		ID:       0,
		Name:     "",
		Email:    "",
		Password: "",
		Books:    nil,
	}
	body, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, "/createUser", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `User created successfully`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	expectedCount := 1
	if len(CacheDatabase.Users) != expectedCount {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUploadBookToUser(t *testing.T) {

	testCases := map[string]struct {
		params     map[string]string
		statusCode int
	}{
		"good params": {
			map[string]string{
				"userID": "1", "bookID": "1",
			},
			http.StatusOK,
		},
		"without params": {
			map[string]string{},
			http.StatusBadRequest,
		},
	}
	CacheDatabase.Users = []model.User{{
		ID:       1,
		Name:     "",
		Email:    "",
		Password: "",
		Books:    nil,
	}}
	CacheDatabase.Books = []model.Book{{
		ID:               1,
		Name:             "",
		Author:           "",
		Pages:            0,
		PercentageOfRead: 0,
	}}
	for tc, tp := range testCases {
		req, _ := http.NewRequest("GET", "/uploadBookToUser", nil)
		q := req.URL.Query()
		for k, v := range tp.params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
		rec := httptest.NewRecorder()
		handler := http.HandlerFunc(UploadBookToUser)
		handler.ServeHTTP(rec, req)
		res := rec.Result()
		if res.StatusCode != tp.statusCode {
			t.Errorf("`%v` failed, got %v, expected %v", tc, res.StatusCode, tp.statusCode)
		}
	}
}
