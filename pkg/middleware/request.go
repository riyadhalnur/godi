package middleware

import (
	"context"
	"net/http"

	"github.com/gofrs/uuid"
)

// RequestID adds a uuid to all incoming requests
// and attaches it to the context
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID, _ := uuid.NewV4()

		ctx := r.Context()
		ctx = context.WithValue(r.Context(), "RequestID", requestID.String())

		w.Header().Set("X-Request-ID", requestID.String())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
