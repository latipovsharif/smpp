package rabbit

// Message format to share between sender and consumer
type Message struct {
	Message string
	Store   string
}
