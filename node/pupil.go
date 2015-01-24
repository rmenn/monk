package node

import (
	"bytes"
	"log"
	"os"
	"os/signal"

	"github.com/gdamore/mangos"
	"github.com/gdamore/mangos/protocol/req"
	"github.com/gdamore/mangos/protocol/respondent"
	"github.com/gdamore/mangos/transport/tcp"

	"code.google.com/p/go-uuid/uuid"

	capn "github.com/glycerine/go-capnproto"

	"github.com/sudharsh/monk/messages"
)

type Pupil struct {
	Node
	UUID                  string
	MasterRegistrationURL string
	Sock                  mangos.Socket
}

func NewPupil(listenURL string, masterRegistrationURL string) *Pupil {
	p := Pupil{}
	p.UUID = uuid.New()
	p.MasterRegistrationURL = masterRegistrationURL
	//	p.MasterEventURL = masterEventURL
	p.URL = listenURL
	return &p
}

func (p *Pupil) Start() {
	c := make(chan os.Signal, 1)
	regChan := make(chan []byte, 1)
	eventChan := make(chan []byte, 1)

	signal.Notify(c, os.Interrupt, os.Kill)

	var err error
	if p.Sock, err = req.NewSocket(); err != nil {
		log.Fatalf("Couldn't get pair socket - %s\n", err.Error())
	}
	p.Sock.AddTransport(tcp.NewTransport())

	go p.register(regChan)

	for {
		select {
		case registrationResponse := <-regChan:
			go p.handleRegistration(registrationResponse, eventChan)
		case event := <-eventChan:
			log.Printf("Received event - %s\n", string(event))
		case s := <-c:
			log.Printf("Got signal to quit. Bye! - %s\n", s)
			os.Exit(0)
		}
	}
}

// Private funcs follow

// Register this pupil with the master by packing a bunch of facts
func (p *Pupil) register(regChan chan []byte) {
	var err error
	log.Printf("Registering with master at %s\n", p.MasterRegistrationURL)
	if err = p.Sock.Dial(p.MasterRegistrationURL); err != nil {
		log.Fatalf("Wasn't able to reach the monk master at %s - %s", p.MasterRegistrationURL, err.Error())
	}

	// We need to receive this only once
	regMessage := p.prepareRegistrationMessage()
	err = p.Sock.Send(regMessage)
	if err != nil {
		log.Fatalf("Couldn't send registration request\n")
	}

	var msg []byte
	msg, err = p.Sock.Recv()
	if err != nil {
		log.Fatalf("Error on registering with the monk master - %s", err.Error())
	}
	regChan <- msg
}

// FIXME: Use something json.Marshal interface to convert it into a capnp message format
func isSuccess(ackMessage []byte) bool {
	buf := bytes.NewBuffer(ackMessage)

	capMsg, err := capn.ReadFromStream(buf, nil)
	if err != nil {
		log.Printf("Error unpacking message - %s\n", err.Error())
		return false
	}
	s := messages.ReadRootResponse(capMsg)
	return s.Success()
}

func (p *Pupil) prepareRegistrationMessage() []byte {
	seg := capn.NewBuffer(nil)
	_p := messages.NewRootPupil(seg)
	_p.SetUuid(p.UUID)
	_p.SetUrl(p.URL)

	buf := bytes.Buffer{}
	seg.WriteTo(&buf)

	return buf.Bytes()
}

// Once the monk master responds saying the pupil registration is successful
// Close the socket to the registration port of the master
// and re-use the same socket to start this node as a respondent to the event
// socket in the master.
// This is done since registration should be done only once anyway
func (p *Pupil) handleRegistration(response []byte, eventChan chan []byte) {
	if !isSuccess(response) {
		log.Fatal("Couldn't register with the monk master")
	}
	log.Printf("pupil registered successfully")

	// Close this registration socket and open another one as
	// a respondent
	p.Sock.Close()

	var err error
	if p.Sock, err = respondent.NewSocket(); err != nil {
		log.Fatal("Couldn't open the event port")
	}
	p.Sock.AddTransport(tcp.NewTransport())

	if err = p.Sock.Listen(p.URL); err != nil {
		log.Fatalf("Error when listening on %s - %s\n", p.URL, err.Error())
	}

	log.Printf("pupil ready to receive events from the master")

	for {
		resp, err := p.Sock.Recv()
		if err != nil {
			log.Printf("WARNING: Error when trying to recv() - %s\n", err.Error())
		}
		eventChan <- resp
	}

}
