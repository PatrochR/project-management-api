package middleware

import (
	"log"
	"net/http"
)

// FIXME: make it better
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method)
		next.ServeHTTP(w, r)
	})
}
