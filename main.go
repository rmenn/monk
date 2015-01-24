package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"

	"github.com/sudharsh/monk/node"
)

func gatherCommands() []cli.Command {

	globalFlags := []cli.Flag{
		cli.IntFlag{
			Name:   "registration-port",
			Value:  5250,
			Usage:  "Registration port for monk agents",
			EnvVar: "MONK_REGISTRATION_PORT",
		},
	}

	commands := []cli.Command{
		{
			Name:  "pupil",
			Usage: "Use this node as a monk pupil",
			Flags: append(globalFlags, cli.StringFlag{
				Name:   "master",
				Value:  "127.0.0.1",
				Usage:  "Location of the master monk node",
				EnvVar: "MONK_MASTER",
			},
				cli.IntFlag{
					Name:   "port",
					Value:  5256,
					Usage:  "monk agent port",
					EnvVar: "MONK_AGENT_PORT",
				}),
			Action: startPupil,
		},
		{
			Name:  "master",
			Usage: "Use this node as a monk master",
			Flags: append(globalFlags,
				cli.IntFlag{
					Name:   "publish-port",
					Value:  5251,
					Usage:  "monk master port",
					EnvVar: "MONK_PUBLISH_PORT",
				}),
			Action: startMaster,
		},
	}

	return commands
}

func startMaster(c *cli.Context) {
	registrationURL := fmt.Sprintf("tcp://%s:%d", "0.0.0.0", c.Int("registration-port"))
	eventURL := fmt.Sprintf("tcp://%s:%d", "0.0.0.0", c.Int("event-port"))
	m := node.NewMaster(registrationURL, eventURL)
	m.Start()
}

func startPupil(c *cli.Context) {
	masterRegistrationURL := fmt.Sprintf("tcp://%s:%d", c.String("master"), c.Int("registration-port"))
	//masterEventURL := fmt.Sprintf("tcp://%s:%d", c.String("master"), c.Int("event-port"))
	listenURL := fmt.Sprintf("tcp://%s:%d", "0.0.0.0", c.Int("port"))
	p := node.NewPupil(listenURL, masterRegistrationURL)
	p.Start()
}

func main() {
	app := cli.NewApp()
	app.Name = "monk"
	app.Usage = "Efficient configuration management"
	app.Version = "0.1"
	app.Commands = gatherCommands()
	app.Run(os.Args)
}
