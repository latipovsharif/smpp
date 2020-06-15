package rabbit

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

// Channel for sending and receiving sms's
const (
	SMSChannel  string = "SMS_CHANNEL"
	SMSConsumer string = "SMS_CONSUMER"
)

// Session session for rabbitmq
type Session struct {
	connection *amqp.Connection
	queue      amqp.Queue
	channel    *amqp.Channel
	done       chan bool
}

// NewSession return new rabbitmq session
func NewSession() (*Session, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to rabbitmq")
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "cannot create new channel")
	}

	q, err := ch.QueueDeclare(
		SMSChannel, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)

	s := &Session{
		connection: conn,
		channel:    ch,
		queue:      q,
		done:       make(chan bool),
	}

	return s, nil
}

// Consume start consuming from SMSChannel
func (s *Session) Consume() {
	msgs, err := s.channel.Consume(
		s.queue.Name, // queue
		SMSConsumer,  // consumer
		false,        // auto-ack
		true,         // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	if err != nil {
		log.Fatalf("cannot consume from channel %v", err)
	}

	for d := range msgs {
		message := &Message{}
		if err := json.Unmarshal(d.Body, message); err != nil {
			fmt.Println(err)
			if err := d.Nack(false, true); err != nil {
				log.Fatalf("cannot nack message")
			}
		}
		fmt.Println(message.Message)
		if err := d.Ack(false); err != nil {
			log.Fatalf("cannot ack message")
		}
	}
	<-s.done
}

// Close the session and cleanup
func (s *Session) Close() {
	s.channel.Close()
	s.connection.Close()
	s.done <- true
}
