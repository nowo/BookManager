package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBook(t *testing.T) {
	testCases := map[string]struct {
		statusCode int
		body       string
	}{
		"Proper Body": {
			http.StatusOK,
			`{"id":1,"name":"","author":"","pages":0,"percentageOfRead":0}`,
		},
		"Empty Body": {
			http.StatusBadRequest,
			"",
		},
	}
	for tc, tp := range testCases {
		req, _ := http.NewRequest(http.MethodPost, "/createBook", bytes.NewBufferString(tp.body))
		rec := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateUser)
		handler.ServeHTTP(rec, req)
		res := rec.Result()
		if res.StatusCode != tp.statusCode {
			t.Errorf("`%v` failed, got %v, expected %v", tc, res.StatusCode, tp.statusCode)
		}
	}
}
