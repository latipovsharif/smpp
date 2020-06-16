package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"smpp/rabbit"
	"smpp/smsc"
)

var messages chan rabbit.Message

func main() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt, os.Kill)
	messages := make(chan rabbit.Message)

	s := smsc.NewSession()

	rs, err := rabbit.NewSession()
	if err != nil {
		log.Fatalf("cannot get rabbit session")
	}

	go s.SendAndReceiveSMS()
	go rs.Consume(messages)
	go s.SubmitSM(messages)
	go s.QuerySM()

	fmt.Println("awaiting signal")

	<-sigs
	s.Close()
	rs.Close()
}
