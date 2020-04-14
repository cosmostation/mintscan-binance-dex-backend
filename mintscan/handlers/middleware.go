package handlers

import (
	"log"
	"net/http"

	"github.com/tomasen/realip"
)

// MiddlewareLogRequest logs incoming requests and calls next handler
func MiddlewareLogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		clientIP := realip.FromRequest(r)
		log.Printf("%s %s [%s]\n", r.Method, r.URL, clientIP)

		next.ServeHTTP(rw, r)
	})
}
