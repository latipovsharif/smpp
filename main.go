package main

import (
	"context"
	"fmt"
	"os"
	"smpp/ent"
	"smpp/f_base"
	"smpp/rabbit"
	"smpp/smsc"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	_ "github.com/lib/pq"
)

const logFilePath = "logs/smpp.log"

//const appVersion = "0.0.1"

func main() {

	lumberjackLogRotate := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    2,   // Max megabytes before log is rotated
		MaxBackups: 500, // Max number of old log files to keep
		MaxAge:     60,  // Max number of days to retain log files
		Compress:   true,
	}
	log.SetOutput(lumberjackLogRotate)

	client, err := ent.Open("postgres", "postgres://postgres:123@localhost:5432/testdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	sigs := make(chan os.Signal)
	//signal.Notify(sigs, os.Interrupt, os.Kill)
	messages := make(chan ent.Messages)

	s := smsc.NewSession(client)

	rs, cacheMap, err := rabbit.NewSession(client)
	if err != nil {
		log.Fatalf("cannot get rabbit session")
	}
	err = f_base.FBaseCon(cacheMap)
	if err != nil {
		log.Fatalf("cannot connection FireBase: %v", err)
	}

	go s.SendAndReceiveSMS()
	go rs.Consume(messages, cacheMap)
	go s.SubmitSM(messages)
	go rs.SendingMessage(messages, cacheMap)

	fmt.Println("awaiting signal")

	<-sigs
	s.Close()
	rs.Close()

}
