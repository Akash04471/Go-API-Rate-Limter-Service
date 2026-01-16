package middleware

import (
	"fmt"
	"net/http"
	"time"

	"api-rate-limiter/config"
	"api-rate-limiter/ratelimiter"
)

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		clientID := ratelimiter.ClientID(r.RemoteAddr)

		fmt.Println("Incoming request from:", clientID)

		allowed := ratelimiter.AllowRequest(
			clientID,
			config.RequestLimit,
			time.Duration(config.TimeWindowSeconds)*time.Second,
		)

		// if-else control flow
		if !allowed {
			fmt.Println("Request blocked:", clientID)
			w.WriteHeader(http.StatusTooManyRequests)

			// Anonymous struct for JSON-style response
			response := struct {
				Status  string
				Message string
			}{
				Status:  "blocked",
				Message: "Rate limit exceeded",
			}

			fmt.Fprintf(w, "%+v\n", response)
			return
		}

		fmt.Println("Request allowed:", clientID)
		next.ServeHTTP(w, r)
	})
}
