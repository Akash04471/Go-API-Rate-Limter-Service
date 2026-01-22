package ratelimiter

import (
	"sync"
	"time"
)

type Client struct {
	RequestCount int
	ResetTime    time.Time
}

var (
	clients = make(map[string]*Client)
	mu      sync.Mutex
)

// Get client ID
func GetClientID(remoteAddr string) string {
	return remoteAddr
}

// Core rate limiting logic
func AllowRequest(clientID string, limit int, window time.Duration) (bool, int, int) {
	mu.Lock()
	defer mu.Unlock()

	client, exists := clients[clientID]
	if !exists {
		client = &Client{}//zero values
		clients[clientID] = client
	}

	now := time.Now()

	if client.ResetTime.IsZero() || now.After(client.ResetTime) {
		client.RequestCount = 0
		client.ResetTime = now.Add(window)
	}

	if client.RequestCount >= limit {
		remaining := int(client.ResetTime.Sub(now).Seconds())
		return false, client.RequestCount, remaining
	}

	client.RequestCount++
	remaining := int(client.ResetTime.Sub(now).Seconds())
	return true, client.RequestCount, remaining
}
