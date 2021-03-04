package pgdb

import (
	"github.com/go-pg/pg/v9"
)

// Session session for rabbitmq
type Session struct {
	connection connection
}

// NewSession return new rabbitmq session
func NewSession(db *pg.DB) *Session {
	s := &Session{}
	s.connection.syncAllConnections()

	return s
}

// Consume start consuming from SMSChannel
func (s *Session) Consume(c chan<- Message) {
	for _, con := range s.connection.DB {
		con.Model()
	}
}

// Close the session and cleanup
func (s *Session) Close() {
}
