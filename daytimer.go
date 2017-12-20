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
	auth      *Authentication   // OAuth2 Authentication
	client    *http.Client      // HTTP client for API interaction
	gcal      *calendar.Service // Google Calendar Service RPCs
	calendars Calendars         // The calendards beling to the user
}

// Agenda returns a listing of the events for a specific day
func (dt Daytimer) Agenda(date time.Time) (*Agenda, error) {
	var err error
	agenda := new(Agenda)
	agenda.items, err = dt.Upcoming(10)
	return agenda, err
}

// Upcoming returns the n next events on the specified calendar.
func (dt Daytimer) Upcoming(n int64) ([]*AgendaItem, error) {
	t := time.Now().Format(time.RFC3339)
	events, err := dt.gcal.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(t).MaxResults(n).OrderBy("startTime").Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve the %d upcoming user's events: %v", n, err)
	}

	upcoming := make([]*AgendaItem, 0, len(events.Items))
	for _, event := range events.Items {
		item := &AgendaItem{
			calendar: nil, // TODO: add calendar here
			event:    event,
		}
		upcoming = append(upcoming, item)
	}

	return upcoming, nil
}

// Calendars returns the Calendar objects belonging to the user. The first
// call to this method will request the calendars list from Google, then mark
// calendars as active or not based on the user's configuration. This is then
// cached for the reminder of the process, but can be fetched again by
// clearing daytimer cache.
func (dt Daytimer) Calendars() (Calendars, error) {
	if dt.calendars == nil {
		// Fetch calendars from Google
		cals, err := dt.gcal.CalendarList.List().Do()
		if err != nil {
			return nil, err
		}

		// Create the calendars mapping
		dt.calendars = make(map[string]*Calendar)
		for _, cal := range cals.Items {
			dt.calendars[cal.Id] = &Calendar{item: cal}
		}

		// Mark calendars active from config otherwise set primary as the
		// active calendar and dump the active calendars file.
		if err := dt.calendars.loadActive(); err != nil {
			for _, cal := range dt.calendars {
				if cal.item.Primary {
					cal.active = true
				}
			}
			if err := dt.calendars.dumpActive(); err != nil {
				return nil, err
			}
		}

	}

	return dt.calendars, nil
}
