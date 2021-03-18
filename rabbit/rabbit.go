package rabbit

import (
	"context"
	"encoding/json"
	"os"
	"smpp/ent"
	"smpp/ent/user"
	"time"

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
func NewSession(db *ent.Client) (*Session, error) {
	urlAmqp := os.Getenv("RABBITMQ_HOST")
	conn, err := amqp.Dial(urlAmqp)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to rabbitmq")
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "cannot create new channel")
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

	return s, nil
}

// Consume start consuming from SMSChannel
func (s *Session) Consume(c chan<- ent.Messages) {
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
		ctx := context.Background()
		message := ent.Messages{}
		if err = json.Unmarshal(d.Body, &message); err != nil {
			log.Errorf("cannot unmarshal message: %v", err)
			if err := d.Nack(false, true); err != nil {
				log.Errorf("cannot nack message: %v", err)
			}
			continue
		}

		if s.chekBalans(ctx, message.UserId) {
			message.State = int(StateNew)
		}

		if _, err = s.db.Messages.Create().Save(ctx); err != nil {
			log.Errorf("cannot insert message: %v", err)
			d.Nack(false, true)
		}
		s.createUserMessage(ctx, message.UserId, message.ProviderId)
		if err = d.Ack(false); err != nil {
			log.Error("cannot ack message")
		}
		if message.State == 0 {
			continue
		}
		c <- message
	}

	<-s.done
}

// Close the session and cleanup
func (s *Session) Close() {
	s.channel.Close()
	s.connection.Close()
	s.done <- true
}

func (s *Session) chekBalans(ctx context.Context, userID uuid.UUID) bool {
	usr, _ := s.db.User.Query().
		Where(user.IDEQ(userID)).
		First(ctx)

	messagePrice, _ := s.db.User.Query().Where(user.IDEQ(userID)).QueryRateID().QueryIDPrice().First(ctx)
	if usr.Balance > messagePrice.Price {
		if _, err := s.db.User.Update().Where(user.ID(usr.ID)).
			SetBalance(usr.Balance - messagePrice.Price).
			SetUpdateAt(time.Now()).
			Save(ctx); err != nil {
			log.Errorf("cannot Update user balans: %v", err)
			return false
		}

		return true
	}
	return false
}

func (s *Session) createUserMessage(ctx context.Context, userID uuid.UUID, providerID uuid.UUID) {
	if _, err := s.db.UserMonthMessage.Create().
		SetUserIDID(userID).
		SetProviderIDID(providerID).Save(ctx); err != nil {
		log.Errorf("cannot insert User Month Message: %v", err)
	}
}
