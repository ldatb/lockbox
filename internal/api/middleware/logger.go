package middleware

import (
	"fmt"
	"net/http"

	"gitlab.com/xrs-cloud/lockbox/core/internal/global"
)

// LoggingMiddleware logs incoming HTTP requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log_msg := fmt.Sprintf("Request: %s %s", r.Method, r.RequestURI)
		global.Logger.Debug(log_msg)
		next.ServeHTTP(w, r)
	})
}
