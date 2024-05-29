package model

import "time"

const (
	Valid   Status = "valid"
	Invalid Status = "invalid"
)

type Status string

// Event represents an event that is sent to the subscriber topic
type Event struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"data"`
}

// ProcessedEvent represents an event that has been processed by the handler and is sent to the publish topic
type ProcessedEvent struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"data"`
	Status    Status    `json:"status"`
}
