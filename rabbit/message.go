package rabbit

type state int

// Message states
const (
	StateNew state = iota + 1
	StateDelivered
	StateNotDelivered
)

