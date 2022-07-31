package handler

import (
	"BookManagementApp/CacheDatabase"
	"BookManagementApp/model"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInitialGetUser(t *testing.T) {
	testCases := map[string]struct {
		statusCode int
		function   func()
	}{
		"User Found": {
			http.StatusOK,
			func() {
				CacheDatabase.Users = []model.User{{
					ID:       1,
					Name:     "",
					Email:    "",
					Password: "",
					Books:    nil,
				}}
			},
		},
		"There is No User": {
			http.StatusNotFound,
			nil,
		},
	}
	for tc, tp := range testCases {
		if tp.function != nil {
			tp.function()
		}
		req, _ := http.NewRequest(http.MethodGet, "/getUsers", nil)
		q := req.URL.Query()
		req.URL.RawQuery = q.Encode()
		rec := httptest.NewRecorder()
		handler := http.HandlerFunc(GetUsers)
		handler.ServeHTTP(rec, req)
		res := rec.Result()
		if res.StatusCode != tp.statusCode {
			t.Errorf("`%v` failed, got %v, expected %v", tc, res.StatusCode, tp.statusCode)
		}
	}
}

func TestCreateUser(t *testing.T) {
	testCases := map[string]struct {
		statusCode int
		body       string
	}{
		"Proper Body": {
			http.StatusOK,
			`{"id":1,"name":"test","email":"test","password":"test","books":[]}`,
		},
		"Empty Body": {
			http.StatusBadRequest,
			"",
		},
	}
	for tc, tp := range testCases {
		req, _ := http.NewRequest(http.MethodPost, "/createUser", bytes.NewBufferString(tp.body))
		rec := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateUser)
		handler.ServeHTTP(rec, req)
		res := rec.Result()
		if res.StatusCode != tp.statusCode {
			t.Errorf("`%v` failed, got %v, expected %v", tc, res.StatusCode, tp.statusCode)
		}
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
		Name:     "test",
		Email:    "test",
		Password: "test",
		Books:    nil,
	}}
	CacheDatabase.Books = []model.Book{{
		ID:               1,
		Name:             "test",
		Author:           "test",
		Pages:            123,
		PercentageOfRead: 0,
	}}
	for tc, tp := range testCases {
		req, _ := http.NewRequest(http.MethodGet, "/uploadBookToUser", nil)
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
