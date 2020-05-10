package util

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRespondJSON(t *testing.T) {
	w := httptest.NewRecorder()

	RespondJSON(w, &Response{
		StatusCode: http.StatusOK,
		Body:       `{"test": "hello"}`,
	})

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.JSONEq(t, `{"test": "hello"}`, string(body))
}

func TestErrorJSON(t *testing.T) {
	err := errors.New("some random error")
	w := httptest.NewRecorder()

	ErrorJSON(w, &ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	})

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.JSONEq(t, `{"code":500,"message":"some random error"}`, string(body))
}
