package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := NewRouter()
	router.ServeHTTP(rr, req)

	return rr
}

func TestGetUser(t *testing.T) {
	for _, tt := range []struct {
		statusCode   int
		expectedBody string
		authHeader   string
	}{
		{
			403,
			"",
			"",
		},
		{
			200,
			`{"id":"1","name":"Test User","roles":["admin"]}`,
			"randomToken",
		},
	} {
		req, _ := http.NewRequest("GET", "/api/user/1", nil)
		req.Header.Set("Authorization", tt.authHeader)

		response := executeRequest(req)

		if tt.statusCode != response.Code {
			t.Errorf("Expected response code %d. Got %d\n", tt.statusCode, response.Code)
		}

		if response.Body.String() != tt.expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v",
				response.Body.String(), tt.expectedBody)
		}
	}
}

func TestCreateUser(t *testing.T) {
	for _, tt := range []struct {
		statusCode   int
		expectedBody string
		authHeader   string
	}{
		{
			403,
			"",
			"",
		},
		{
			200,
			`{"id":"2","name":"Test","roles":["member"]}`,
			"randomToken",
		},
	} {
		reqBody, _ := json.Marshal(&User{
			Name:  "Test",
			Roles: []string{"member"},
		})

		req, _ := http.NewRequest("POST", "/api/user", bytes.NewBuffer(reqBody))
		req.Header.Set("Authorization", tt.authHeader)

		response := executeRequest(req)

		if tt.statusCode != response.Code {
			t.Errorf("Expected response code %d. Got %d\n", tt.statusCode, response.Code)
		}

		if response.Body.String() != tt.expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v",
				response.Body.String(), tt.expectedBody)
		}
	}
}
