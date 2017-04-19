// Package daytimer provides interactive functions for the daytimer command
// line utility, implementing a command line application that can access the
// Google Calendar API.
package daytimer

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/context"

	calendar "google.golang.org/api/calendar/v3"
)

// New creates a daytimer and initialize the internal state.
func New() (*Daytimer, error) {
	var err error

	// Initialize the daytimer and authentication
	daytimer := new(Daytimer)
	daytimer.auth = new(Authentication)

	// Create the daytimer client with a background context.
	ctx := context.Background()

	// Load the configuration from client_secret.json
	config, err := daytimer.auth.Config()
	if err != nil {
		return nil, err
	}

	// Load the token from the cache or force authentication
	token, err := daytimer.auth.Token()
	if err != nil {
		return nil, err
	}

	// Create the client with the conifguration
	daytimer.client = config.Client(ctx, token)

	// Create the google calendar service
	daytimer.gcal, err = calendar.New(daytimer.client)
	if err != nil {
		return nil, fmt.Errorf("could not create the google calendar service: %v", err)
	}

	return daytimer, nil
}

// Daytimer provides state and methods for interacting with Google Calendar.
type Daytimer struct {
	auth   *Authentication   // OAuth2 Authentication
	client *http.Client      // HTTP client for API interaction
	gcal   *calendar.Service // Google Calendar Service RPCs
}

// Agenda returns a listing of the events for a specific day
func (dt Daytimer) Agenda(date time.Time) (*Agenda, error) {
	var err error
	agenda := new(Agenda)
	agenda.events, err = dt.Upcoming(10)
	return agenda, err
}

// Upcoming returns the n next events on the calendar.
func (dt Daytimer) Upcoming(n int64) ([]string, error) {
	var upcoming []string

	t := time.Now().Format(time.RFC3339)
	events, err := dt.gcal.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(t).MaxResults(n).OrderBy("startTime").Do()
	if err != nil {
		return upcoming, fmt.Errorf("unable to retrieve the %d upcoming user's events: %v", n, err)
	}

	upcoming = make([]string, 0, len(events.Items))
	for _, i := range events.Items {
		var when string
		// If the DateTime is an empty string, the event is an all day event
		if i.Start.DateTime != "" {
			when = i.Start.DateTime
		} else {
			when = i.Start.Date
		}
		upcoming = append(upcoming, fmt.Sprintf("%s (%s)", i.Summary, when))
	}

	return upcoming, nil
}
