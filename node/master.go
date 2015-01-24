package node

import (
	"log"
	"os"
	"os/signal"

	"github.com/gdamore/mangos"
	"github.com/gdamore/mangos/protocol/rep"
	"github.com/gdamore/mangos/protocol/surveyor"
	"github.com/gdamore/mangos/transport/tcp"
)

type Master struct {
	Node
	RegistrationURL  string
	RegistrationSock mangos.Socket
	EventSock        mangos.Socket
}

func NewMaster(registrationURL string, eventURL string) *Master {
	m := Master{}
	m.URL = eventURL
	m.RegistrationURL = registrationURL
	return &m
}

func (m *Master) Start() {
	c := make(chan os.Signal, 1)
	regChan := make(chan []byte, 1)
	pubChan := make(chan []byte, 1)

	signal.Notify(c, os.Interrupt, os.Kill)

	var err error

	if m.RegistrationSock, err = rep.NewSocket(); err != nil {
		log.Fatalf("Couldn't get registration socket - %s\n", err.Error())
	}
	m.RegistrationSock.AddTransport(tcp.NewTransport())

	if m.EventSock, err = surveyor.NewSocket(); err != nil {
		log.Fatalf("Couldn't get publish socket - %s\n", err.Error())
	}
	m.EventSock.AddTransport(tcp.NewTransport())

	err = m.RegistrationSock.Listen(m.RegistrationURL)
	if err != nil {
		log.Fatalf("Couldn't get the registration port open - %s\n", err.Error())
	}

	err = m.EventSock.Listen(m.URL)
	if err != nil {
		log.Fatalf("Couldn't get the publish port up - %s\n", err.Error())
	}

	go func() {
		for {
			resp, _ := m.RegistrationSock.Recv()
			regChan <- resp
		}
	}()

	go func() {
		for {
			resp, _ := m.EventSock.Recv()
			pubChan <- resp
		}
	}()

	for {
		select {
		case r := <-regChan:
			log.Printf("Got registration - %s", string(r))
			m.RegistrationSock.Send([]byte("ACK"))

		case s := <-c:
			log.Printf("Got signal to quit. Bye! - %s\n", s)
			close(regChan)
			close(pubChan)
			m.RegistrationSock.Close()
			m.EventSock.Close()
			os.Exit(0)
		}
	}
}
