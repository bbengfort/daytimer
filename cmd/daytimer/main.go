// This command implements the daytimer console utility that is the main
// entry point to access and use your Google calendar from the command line.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bbengfort/daytimer"
	"github.com/urfave/cli"
)

func main() {

	// Instantiate the command line application
	app := cli.NewApp()
	app.Name = "daytimer"
	app.Usage = "interact with your Google calendar from the command line"
	app.Version = "0.1"

	// Define commands available to the application
	app.Commands = []cli.Command{
		{
			Name:   "upcoming",
			Usage:  "view the upcoming events on your calendar",
			Action: upcoming,
			Flags: []cli.Flag{
				cli.Int64Flag{
					Name:  "n, number",
					Usage: "number of events to view",
					Value: 10,
				},
			},
		},
	}

	// Run the application
	app.Run(os.Args)
}

func upcoming(c *cli.Context) error {
	daytimer, err := daytimer.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	events, err := daytimer.Upcoming(c.Int64("number"))
	if err != nil {
		log.Fatal(err.Error())
	}

	if len(events) > 0 {
		fmt.Println("Upcoming events:")
		for _, event := range events {
			fmt.Println(event)
		}
	} else {
		fmt.Println("No upcoming events found.")
	}

	return nil
}
