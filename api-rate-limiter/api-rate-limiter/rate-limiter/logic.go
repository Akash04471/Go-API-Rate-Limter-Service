package ratelimiter

import (
	"time"
)

// ---------- UNIT 1 ----------

// Creating your own type (custom type)
type ClientID string

// ---------- UNIT 2 ----------

// Struct to store client data (zero values used automatically)
type Client struct {
	RequestCount int
	ResetTime    time.Time
}

// Embedded struct example
type ClientMeta struct {
	Client
	IsBlocked bool
}

// Map to store clients (make + map)
var Clients = make(map[ClientID]*Client)

// Slice to track blocked clients
var BlockedClients []ClientID

// Multi-dimensional slice (example: request history)
var RequestHistory [][]int

// ---------- Rate Limiting Logic ----------

func AllowRequest(id ClientID, limit int, window time.Duration) bool {

	// Short declaration operator (:=)
	client, exists := Clients[id]

	if !exists {
		// Zero values automatically applied
		client = &Client{
			ResetTime: time.Now().Add(window),
		}
		Clients[id] = client
	}

	// Control flow: conditional
	if time.Now().After(client.ResetTime) {
		client.RequestCount = 0
		client.ResetTime = time.Now().Add(window)
	}

	// if-else for allow / block
	if client.RequestCount < limit {
		client.RequestCount++
		return true
	}

	// Append to slice
	BlockedClients = append(BlockedClients, id)
	return false
}

// Delete from slice example
func RemoveBlocked(id ClientID) {
	for i, v := range BlockedClients {
		if v == id {
			BlockedClients = append(BlockedClients[:i], BlockedClients[i+1:]...)
			break
		}
	}
}
func GetClientID(remoteAddr string) ClientID {
    return ClientID(remoteAddr)
}
