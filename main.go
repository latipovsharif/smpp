package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"smpp/rabbit"
	"time"

	"github.com/linxGnu/gosmpp"
	"github.com/linxGnu/gosmpp/data"
	"github.com/linxGnu/gosmpp/pdu"
)

// Session is main session
type Session struct {
	trans *gosmpp.TransceiverSession
	c     chan *pdu.SubmitSM
}

// NewSession returns new session object
func NewSession(auth gosmpp.Auth) *Session {
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

		OnPDU: handlePDU(),

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
	}
}

func main() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt, os.Kill)

	s := NewSession(gosmpp.Auth{
		SMSC:       os.Getenv("SMSC_HOST"),     // "127.0.0.1:2775",
		SystemID:   os.Getenv("SMSC_LOGIN"),    // "522241",
		Password:   os.Getenv("SMSC_PASSWORD"), // "UUDHWB",
		SystemType: "",
	})

	go s.querySM()

	rs, err := rabbit.NewSession()
	if err != nil {
		log.Fatalf("cannot get rabbit session")
	}

	go rs.Consume()

	fmt.Println("awaiting signal")
	<-sigs
	s.trans.Close()
	rs.Close()
}

func (s *Session) sendingAndReceiveSMS() {
	for n := range s.c {
		if err := s.trans.Transceiver().Submit(n); err != nil {
			fmt.Println(err)
		}
	}
}

func handlePDU() func(pdu.PDU, bool) {
	return func(p pdu.PDU, responded bool) {
		switch pd := p.(type) {
		case *pdu.SubmitSMResp:
			fmt.Printf("SubmitSMResp:%+v\n", pd)
			fmt.Printf("SubmitSMResp:%+v sequence number \n", pd.GetSequenceNumber())

		case *pdu.GenerickNack:
			fmt.Println("GenericNack Received")

		case *pdu.EnquireLinkResp:
			fmt.Println("EnquireLinkResp Received")

		case *pdu.DataSM:
			fmt.Printf("DataSM:%+v\n", pd)

		case *pdu.DeliverSM:
			fmt.Print("how many times i were here ?")
			// fmt.Println(pd)
			// fmt.Println(p)
			// fmt.Println("============")
			// fmt.Println(pd.GetHeader().SequenceNumber)
			// fmt.Println(pd.GetResponse())
			// fmt.Printf("DeliverSM:%+v\n", pd)
			fmt.Println(pd.Message.GetMessage())
			// fmt.Println("============")
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

func (s *Session) newSubmitSM() {
	for i := 0; i < 10; i++ {
		// build up submitSM
		srcAddr := pdu.NewAddress()
		srcAddr.SetTon(5)
		srcAddr.SetNpi(0)
		_ = srcAddr.SetAddress("Abr")

		destAddr := pdu.NewAddress()
		destAddr.SetTon(1)
		destAddr.SetNpi(1)
		_ = destAddr.SetAddress("992927599997")

		submitSM := pdu.NewSubmitSM().(*pdu.SubmitSM)
		submitSM.SourceAddr = srcAddr
		submitSM.DestAddr = destAddr
		_ = submitSM.Message.SetMessageWithEncoding("ҷҷҷҷҷӣӣӣқққҳҳҳҳҳҳҳҳҳҳҳҳҳҳққййййёёёӣӣӣъҷҷҷҷҷҷ", data.UCS2)
		submitSM.ProtocolID = 0
		submitSM.RegisteredDelivery = 1
		submitSM.ReplaceIfPresentFlag = 0
		submitSM.EsmClass = 0
		submitSM.SetSequenceNumber(int32(i))
		s.c <- submitSM
		// fmt.Println(submitSM.GetSequenceNumber())
	}
}

func (s *Session) querySM() {
	q := pdu.NewQuerySM()
	q.(*pdu.QuerySM).MessageID = "8ba749dfda7b2cbe75d42581065c037ffdf0f5ea46f0c9a32885709a9e244c3e"
	if err := s.trans.Transceiver().Submit(q); err != nil {
		log.Fatalf("cannot send query sm %v", err)
	}
}
