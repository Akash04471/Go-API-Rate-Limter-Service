package config

import "time"

// Configuration-level variables
// Using var so they are accessible across packages

var (
	// Maximum number of requests allowed per client
	RequestLimit int = 10

	// Time window for rate limiting
	TimeWindow time.Duration = 1 * time.Minute

	// Server port number
	ServerPort string = ":8080"
)
