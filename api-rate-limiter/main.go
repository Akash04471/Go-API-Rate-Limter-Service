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

	fmt.Println("Server listening on port", config.ServerPort)

	err := http.ListenAndServe(config.ServerPort, nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
