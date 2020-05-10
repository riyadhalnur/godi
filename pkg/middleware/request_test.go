package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/godi/pkg/server/util"

	"github.com/stretchr/testify/assert"
)

func TestRequestIDMiddleware(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		assert.Nil(t, err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Context().Value(util.RequestIDKey).(string)
		assert.Greater(t, len(requestID), 1)
	})

	rr := httptest.NewRecorder()

	handler := RequestID(testHandler)
	handler.ServeHTTP(rr, req)

	assert.NotEmpty(t, rr.Header().Get("X-Request-ID"))
}
