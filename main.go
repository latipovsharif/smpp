package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"smpp/rabbit"
	"smpp/smsc"

	"github.com/go-pg/pg/v9"
)

var db *pg.DB
var messages chan rabbit.Message

func main() {
	db = pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		Database: "messages",
		User:     "postgres",
		Password: "123",
	})

	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt, os.Kill)
	messages := make(chan rabbit.Message)

	s := smsc.NewSession(db)

	rs, err := rabbit.NewSession(db)
	if err != nil {
		log.Fatalf("cannot get rabbit session")
	}

	go s.SendAndReceiveSMS()
	go rs.Consume(messages)
	go s.SubmitSM(messages)

	fmt.Println("awaiting signal")

	<-sigs
	s.Close()
	rs.Close()
}
