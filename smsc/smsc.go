package smsc

import (
	"fmt"
	"log"
	"os"
	"smpp/rabbit"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/linxGnu/gosmpp"
	"github.com/linxGnu/gosmpp/data"
	"github.com/linxGnu/gosmpp/pdu"
)

// Session is main session
type Session struct {
	trans *gosmpp.TransceiverSession
	c     chan *pdu.SubmitSM
	db    *pg.DB
}

// NewSession returns new session object
func NewSession(db *pg.DB) *Session {
	auth := gosmpp.Auth{
		SMSC:       os.Getenv("SMSC_HOST"),     // "127.0.0.1:2775",
		SystemID:   os.Getenv("SMSC_LOGIN"),    // "522241",
		Password:   os.Getenv("SMSC_PASSWORD"), // "UUDHWB",
		SystemType: "",
	}

	trans, err := gosmpp.NewTransceiverSession(gosmpp.NonTLSDialer, auth, gosmpp.TransceiveSettings{
		EnquireLink: 5 * time.Second,

		OnSubmitError: func(p pdu.PDU, err error) {
			log.Fatal(err)
		},

		OnReceivingError: func(err error) {
			fmt.Println(err)
		},

		OnRebindingError: func(err error) {
			fmt.Println(err)
		},

		OnPDU: handlePDU(db),

		OnClosed: func(state gosmpp.State) {
			fmt.Println(state)
		},
	}, 5*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	return &Session{
		trans: trans,
		c:     make(chan *pdu.SubmitSM),
		db:    db,
	}
}

// SendAndReceiveSMS to smsc
func (s *Session) SendAndReceiveSMS() {
	for n := range s.c {
		if err := s.trans.Transceiver().Submit(n); err != nil {
			fmt.Println(err)
		}
	}
}

func handlePDU(db *pg.DB) func(pdu.PDU, bool) {
	return func(p pdu.PDU, responded bool) {
		switch pd := p.(type) {
		case *pdu.SubmitSMResp:
			if _, err := db.Model((*rabbit.Message)(nil)).
				Set("smsc_message_id = ?, state = ?, last_updated_date = ?", pd.MessageID, rabbit.StateDelivered, time.Now()).
				Where("id = ?", pd.GetSequenceNumber()).Update(); err != nil {
				fmt.Println(err)
			}
		case *pdu.GenerickNack:
			fmt.Printf("SubmitSMResp:%+v sequence number \n", pd.GetSequenceNumber())
			fmt.Println("GenericNack Received")
		case *pdu.EnquireLinkResp:
			fmt.Println("EnquireLinkResp Received")
		case *pdu.DataSM:
			fmt.Printf("DataSM:%+v\n", pd)
		case *pdu.DeliverSM:
			fmt.Printf("DeliverSM:%+v\n", pd)
		case *pdu.QuerySM:
			fmt.Println("QuerySM")
			fmt.Println(pd)
		case *pdu.QuerySMResp:
			fmt.Println("QuerySMResp")
			fmt.Println(pd)
		default:
			fmt.Println("Default ")
			fmt.Println(pd)
		}
	}
}

// SubmitSM submit new short message
func (s *Session) SubmitSM(c <-chan rabbit.Message) {
	for m := range c {
		srcAddr := pdu.NewAddress()
		srcAddr.SetTon(5)
		srcAddr.SetNpi(0)
		_ = srcAddr.SetAddress(m.Src)

		destAddr := pdu.NewAddress()
		destAddr.SetTon(1)
		destAddr.SetNpi(1)
		_ = destAddr.SetAddress(m.Dst)

		submitSM := pdu.NewSubmitSM().(*pdu.SubmitSM)
		submitSM.SourceAddr = srcAddr
		submitSM.DestAddr = destAddr
		_ = submitSM.Message.SetMessageWithEncoding(m.Message, data.UCS2)
		submitSM.ProtocolID = 0
		submitSM.RegisteredDelivery = 1
		submitSM.ReplaceIfPresentFlag = 0
		submitSM.EsmClass = 0
		submitSM.SetSequenceNumber(m.ID)
		s.c <- submitSM
	}
}

// QuerySM make query to smsc about state of sms by message id
func (s *Session) QuerySM() {
	q := pdu.NewQuerySM()
	q.(*pdu.QuerySM).MessageID = "0cacd4da82d8c2df76ebfd60a8a5ffa498d4f7b1259e17133ff5ed2db89befdc"
	a := pdu.NewAddress()
	a.SetTon(5)
	a.SetNpi(0)
	a.SetAddress("oasis")

	q.(*pdu.QuerySM).SourceAddr = a
	if err := s.trans.Transceiver().Submit(q); err != nil {
		log.Fatalf("cannot send query sm %v", err)
	}
}

// Close session
func (s *Session) Close() {
	s.trans.Close()
}
