package main

import (
	"fmt"
	"net/http"

	"api-rate-limiter/config"
	"api-rate-limiter/middleware"
)

func main() {

	fmt.Println("Starting API Rate Limiter Server...")

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Request successful")
	})

	// Middleware usage
	http.Handle("/api/test", middleware.RateLimiter(handler))

	// for loop (continuous server readiness example)
	for i := 1; i <= 1; i++ {
		fmt.Println("Server listening on port", config.ServerPort)
	}

	http.ListenAndServe(config.ServerPort, nil)
}
