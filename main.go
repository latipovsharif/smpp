package main

import (
	"os"
	"os/signal"
	"smpp/rabbit"
	"smpp/server"
	"smpp/smsc"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/go-pg/pg/v9"
	"gopkg.in/natefinch/lumberjack.v2"
)

var db *pg.DB
var messages chan rabbit.Message

const logFilePath = "logs/smpp.log"
const appVersion = "0.0.1"

func main() {

	log.Infof("starting application: %v", appVersion)

	lumberjackLogRotate := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    2,   // Max megabytes before log is rotated
		MaxBackups: 500, // Max number of old log files to keep
		MaxAge:     60,  // Max number of days to retain log files
		Compress:   true,
	}
	log.SetOutput(lumberjackLogRotate)

	db = pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		Database: "messages",
		User:     "postgres",
		Password: "123",
	})

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	messages = make(chan rabbit.Message)
	s := smsc.NewSession(db)

	go s.SendAndReceiveSMS()

	srv := server.Server{}
	go srv.Run(log.StandardLogger(), db, s)

	log.Error("enter ctrl+c to exit")

	<-sigs
	s.Close()
}
