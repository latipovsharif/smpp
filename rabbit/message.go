package rabbit

type state int

// Message states
const (
	StateNew state = iota + 1
	StateDelivered
	StateNotDelivered
)

// Message format to share between sender and consumer
// type Message struct {
// 	SequenceNumber  int32      `json:"sequenceNumber"`
// 	ExternalID      string     `json:"externalID"`
// 	Dst             string     `json:"dst"`
// 	Message         string     `json:"message"`
// 	Src             string     `json:"src"`
// 	State           state      `json:"state"`
// 	SMSCMessageID   string     `json:"smsCMessageID"`
// 	ProviderID      string     `json:"providerID"`
// 	UserID          string     `json:"userID"`
// 	CreatedAt       *time.Time `json:"createdAt"`
// 	LastUpdatedDate *time.Time `json:"lastUpdatedDate"`
// }
