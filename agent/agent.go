package agent

import (
	"os"
	"os/signal"
	"log"
	
	"github.com/zeromq/gyre"
)


type Agent struct {
	Conn *gyre.Gyre
	Hostname string
}


func New() *Agent {
	a := Agent{}
	var err error
	a.Conn, err = gyre.New()
	if err != nil {
		log.Fatalf("Error when initalizing monk agent - %s\n", err.Error())
	}
	a.Hostname, err = os.Hostname()
	if err != nil {
		log.Fatalf("Couldn't get hostname for this agent - %s\n", err.Error())
	}
	return &a
}

func (agent *Agent) Start() {

	c := make(chan os.Signal, 1)
	// Block until SIGs are given
	signal.Notify(c, os.Interrupt, os.Kill)	
	
	err := agent.Conn.Start()
	if err != nil {
		log.Fatalf("Error when starting monk agent - %s\n", err.Error())
	}
	agent.Conn.Join("MONK_AGENT")
	
	log.Printf("Started monk agent %s at %s\n", agent.Conn.Name(), agent.Hostname)
	
	for {
		select {
		case e := <- agent.Conn.Events():
			switch e.Type() {
			case gyre.EventShout:
				log.Printf("Executing state - %s\n", string(e.Msg()))
			}			
		case s := <-c:
			log.Printf("Got %s. Bye!\n", s)
			os.Exit(0)
		}
	}
	
}