package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"api-rate-limiter/config"
	"api-rate-limiter/rate-limiter"
)

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get client identifier
		clientID := ratelimiter.GetClientID(r.RemoteAddr)
		fmt.Println("Incoming request from:", clientID)

		// Rate limit check
		allowed, remaining, resetTime := ratelimiter.AllowRequest(
			clientID,
			config.RequestLimit,
			config.TimeWindow,
		)

		// Blocked request
		if !allowed {
			fmt.Println("Request blocked:", clientID)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)

			// Anonymous struct (Unit 2 concept)
			response := struct {
				Status    string `json:"status"`
				Message   string `json:"message"`
				RetryAfter string `json:"retry_after"`
			}{
				Status:    "blocked",
				Message:   "Rate limit exceeded",
				RetryAfter: time.Unix(int64(resetTime), 0).Format("15:04:05"),
			}

			json.NewEncoder(w).Encode(response)
			return
		}

		// Allowed request
		fmt.Println("Request allowed:", clientID, "| Remaining:", remaining)
		next.ServeHTTP(w, r)
	})
}
