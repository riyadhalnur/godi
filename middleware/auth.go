package middleware

import (
	"net/http"
)

// Authenticate middleware is used to authenticate protected
// api routes. Mounted to authenticate any route under /api
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
