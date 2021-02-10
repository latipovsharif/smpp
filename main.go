package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"smpp/ent"
	"smpp/rabbit"
	"smpp/smsc"

	"github.com/go-pg/pg/v9"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/facebook/ent/dialect"
	entsql "github.com/facebook/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Open new connection
func Open(dbURL string) *ent.Client {
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	x0 := ent.NewClient(ent.Driver(drv))
	return x0
}

var db *pg.DB
var messages chan rabbit.Message

const logFilePath = "logs/smpp.log"
const appVersion = "0.0.1"

func main() {

	lumberjackLogRotate := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    2,   // Max megabytes before log is rotated
		MaxBackups: 500, // Max number of old log files to keep
		MaxAge:     60,  // Max number of days to retain log files
		Compress:   true,
	}
	log.SetOutput(lumberjackLogRotate)

	client := Open("postgresql://postgres:123@127.0.0.1/messages")
	
	// Your code. For example:
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	
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
