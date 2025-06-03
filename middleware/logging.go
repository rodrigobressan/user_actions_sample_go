package middleware

import (
	"log"
	"net/http"
	"time"
)

// statusRecorder is a wrapper to capture HTTP response status codes
type statusRecorder struct {
	http.ResponseWriter
	status int
}

// WriteHeader captures the status code for logging
func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

// LoggingMiddleware logs incoming HTTP requests with method, path, status, and duration
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(rec, r)

		duration := time.Since(start)
		log.Printf("[%s] %s %d %s", r.Method, r.URL.Path, rec.status, duration)
	})
}
