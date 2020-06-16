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
	ID              int32
	ExternalID      string
	Dst             string
	Message         string
	Src             string
	State           state
	CreatedAt       *time.Time
	LastUpdatedDate *time.Time
	SMSCMessageID   string
}
