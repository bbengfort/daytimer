// This command implements the daytimer console utility that is the main
// entry point to access and use your Google calendar from the command line.
package main

import (
	"fmt"
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
			Name:   "auth",
			Usage:  "reauthenticate daytimer with Google",
			Action: auth,
		},
		{
			Name:   "agenda",
			Usage:  "print out the agenda for a specific day",
			Action: agenda,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "d, date",
					Usage: "specify the date as YYYY/MM/DD (default today)",
				},
			},
		},
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

// Forces a reauthentication with Google
func auth(c *cli.Context) error {
	auth := new(daytimer.Authentication)
	if err := auth.Authenticate(); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	return nil
}

// Computes the agenda for a given date
func agenda(c *cli.Context) error {
	dt, err := daytimer.New()
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	date, err := daytimer.ParseDate(c.String("date"))
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	agenda, err := dt.Agenda(date)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	fmt.Println(agenda.String())
	return nil
}

// Lists the n upcoming events
func upcoming(c *cli.Context) error {
	dt, err := daytimer.New()
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	events, err := dt.Upcoming(c.Int64("number"))
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
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
