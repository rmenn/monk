package main

import (
	"os"
	"log"

	_ "github.com/zeromq/gyre"
	"github.com/codegangsta/cli"
)

func gatherCommands() []cli.Command {
	return []cli.Command {
		{
			Name: "agent",
			Usage: "Use this node as a monk agent",
      		Flags: []cli.Flag {
				cli.StringFlag {
					Name: "master",
					Usage: "Location of the master monk node",
					EnvVar: "MONK_MASTER",
				},
			},
			Action: startAgent,
		},
		{
			Name: "master",
			Usage: "Use this node as a monk master",
			Action: startMaster,
		},
	}
}


func startMaster(c *cli.Context) {
	log.Printf("Starting Master\n")
}


func startAgent(c *cli.Context) {
	log.Printf("Starting agent\n")
}


func main() {
	app := cli.NewApp()
	app.Name = "monk"
	app.Usage = "Efficient configuration management"
	app.Version = "0.1"
	app.Commands = gatherCommands()
	app.Run(os.Args)
}