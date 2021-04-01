package rabbit

type state int

// Message states
const (
	InvalidBalance state = iota
	StateNew       
	StateDelivered
	StateNotDelivered
)
