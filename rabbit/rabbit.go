package rabbit

import (
	"encoding/json"
	"os"
	"smpp/ent"
	"sync"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

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
	db         *ent.Client
}

// NewSession return new rabbitmq session
func NewSession(db *ent.Client) (*Session, *CacheMap, error) {
	urlAmqp := os.Getenv("RABBITMQ_HOST")
	conn, err := amqp.Dial(urlAmqp)
	if err != nil {
		return nil, nil, errors.Wrap(err, "cannot connect to rabbitmq")
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, errors.Wrap(err, "cannot create new channel")
	}

	q, _ := ch.QueueDeclare(
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
		db:         db,
	}

	cacheMap := &CacheMap{
		Hmap:  map[uuid.UUID][]ent.Messages{},
		Mutex: &sync.RWMutex{},
	}

	return s, cacheMap, nil
}

// Consume start consuming from SMSChannel
func (s *Session) Consume(c chan<- ent.Messages, cacheMap *CacheMap) {
	msgs, err := s.channel.Consume(
		s.queue.Name, // queue
		SMSConsumer,  // consumer
		false,        // auto-ack
		true,         // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	//lint:ignore SA4006 this value
	arrMes := []ent.Messages{}
	if err != nil {
		log.Fatalf("cannot consume from channel %v", err)
	}
	for d := range msgs {
		message := ent.Messages{}
		if err = json.Unmarshal(d.Body, &message); err != nil {
			log.Errorf("cannot unmarshal message: %v", err)
			if err := d.Nack(false, true); err != nil {
				log.Errorf("cannot nack message: %v", err)
			}
			continue
		}
		if err = d.Ack(false); err != nil {
			log.Error("cannot ack message")
		}
		cacheMap.Mutex.RLock()
		arrMes = cacheMap.Hmap[uuid.MustParse(message.UserID)]
		arrMes = append(arrMes, message)
		cacheMap.Hmap[uuid.MustParse(message.UserID)] = arrMes
		cacheMap.Mutex.RUnlock()
	}

	<-s.done
}

// Close the session and cleanup
func (s *Session) Close() {
	s.channel.Close()
	s.connection.Close()
	s.done <- true
}
