package pgdb

type State string

const (
	StateNew           State = "NEW"
	StateProcessing    State = "PROCESSING"
	StateQueued        State = "QUEUED"
	StateSent          State = "SENT"
	StateDelivered     State = "Delivered"
	StateUndeliverable State = "UNDELIVERABLE"
)

// Message format to share between sender and consumer
type Message struct {
	tableName struct{} `pg:"sms_queue"` //nolint:golint,structcheck,unused
	Model
	Receiver    string // phone number of receiver
	Sender      string // alphanumeric or phone number of a sender
	Message     string // Final message that will be delivered to customer
	State       State  // Message states
	Attempts    int    // number of attempts to send message
	CreatedBy   string // uuid of a person who created this message
	SmsCenterID string // id given from sms center
}
