package rabbit

// Message format to share between sender and consumer
type Message struct {
	ID          int32
	ExternalID  string
	Destination string
	Message     string
	Store       string
}
