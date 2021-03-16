package rabbit

import "time"

type state int

// Message states
const (
	StateNew state = iota + 1
	StateDelivered
	StateNotDelivered
)

// Message format to share between sender and consumer
type Message struct {
	ID              int32  `json:"id" binding:"required"`
	ExternalID      string `json:"external_id"`
	Dst             string `json:"receiver" binding:"required"`
	Message         string `json:"message" binding:"required"`
	Src             string `json:"sender" binding:"required"`
	State           state
	CreatedAt       *time.Time
	LastUpdatedDate *time.Time
	SMSCMessageID   string
}
