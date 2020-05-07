package middleware

import (
	"log"
	"net/http"
	"os"
)

// Logging is a basic logging middleware that logs
// any incoming request to stdout
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.New(os.Stdout, "[INFO]: ", log.LstdFlags)
		logger.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		next.ServeHTTP(w, r)
	})
}
