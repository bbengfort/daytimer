// This command implements the daytimer console utility that is the main
// entry point to access and use your Google calendar from the command line.
package main

import (
	"fmt"
	"os"
	"strings"

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
			Before: initDaytimer,
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
			Before: initDaytimer,
			Flags: []cli.Flag{
				cli.Int64Flag{
					Name:  "n, number",
					Usage: "number of events to view",
					Value: 10,
				},
			},
		},
		{
			Name:   "calendars",
			Usage:  "manage the calendars used to display activity",
			Action: calendars,
			Before: initDaytimer,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "l, list",
					Usage: "list all available calendars from Google",
				},
				cli.StringFlag{
					Name:  "a, activate",
					Usage: "activate specfied calendars (comma sep)",
				},
			},
		},
	}

	// Run the application
	app.Run(os.Args)
}

//===========================================================================
// Helper Functions
//===========================================================================

var dt *daytimer.Daytimer

func initDaytimer(c *cli.Context) (err error) {
	dt, err = daytimer.New()
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}

//===========================================================================
// Commands
//===========================================================================

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

// Manages the active calendars set by the user
func calendars(c *cli.Context) error {
	cals, err := dt.Calendars()
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	// List all availabe calendars and return
	if c.Bool("list") {
		fmt.Println(cals.String())
		return nil
	}

	// Set calendars as active in config
	if activate := c.String("activate"); activate != "" {
		for _, cid := range strings.Split(activate, ",") {
			cid = strings.TrimSpace(cid)
			if err := cals.Activate(cid); err != nil {
				return cli.NewExitError(err.Error(), 1)
			}
		}
	}

	// Filter active calendars and print
	fmt.Println(cals.Active().String())
	return nil
}
