package middleware

import (
	"net/http"
)

// AuthenticationMiddleware validates the authentication token in the request
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement authentication logic here (e.g., check JWT token)
		next.ServeHTTP(w, r)
	})
}
