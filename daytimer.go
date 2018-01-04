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
	calendars Calendars         // The calendars belonging to the user
	timezone  *time.Location    // The timezone being used for requests
}

// Agenda returns a listing of the events for a specific day and a specific
// set of calendars. If no calendars are supplied, uses the active calendars.
func (dt *Daytimer) Agenda(date time.Time, calendars []string) (*Agenda, error) {
	var err error

	// Validate Calendars
	calendars, err = dt.validateCalendars(calendars)
	if err != nil {
		return nil, err
	}

	// Set the timezone of the date (may be UTC, hopefully a calendar's TZ)
	date = Localize(date, dt.Timezone())

	// Create time range
	after, before := DayRange(date)

	// Create a new agenda to start adding calendar items to
	agenda := new(Agenda)
	agenda.Title = fmt.Sprintf("Daily Agenda for %s", date.Format("Monday, January 2, 2006"))
	agenda.Items = make([]*AgendaItem, 0)
	agenda.Daytimer = dt
	agenda.Counts = make(map[string]int)

	// For each calendar, add agenda items
	for _, cid := range calendars {
		// Construct the query for the calendar
		query := dt.gcal.Events.List(cid)
		query = query.ShowDeleted(false).SingleEvents(true).OrderBy("startTime")
		query = query.TimeMin(after).TimeMax(before)

		var events *calendar.Events
		events, err = query.Do()
		if err != nil {
			return nil, fmt.Errorf("could not retrieve events for calendar '%s' from %s to %s", cid, after, before)
		}

		for _, event := range events.Items {
			item := &AgendaItem{
				Calendar: dt.calendars[cid],
				Event:    event,
			}
			agenda.Items = append(agenda.Items, item)
		}

		agenda.Counts[cid] = len(events.Items)
	}

	// Sort the agenda items
	agenda.Sort()
	return agenda, err
}

// Upcoming returns the n next events on the specified calendar.
func (dt *Daytimer) Upcoming(n int64, cid string) ([]*AgendaItem, error) {
	var ok bool
	var err error
	var cal *Calendar

	cals, err := dt.Calendars()
	if err != nil {
		return nil, err
	}

	cal, ok = cals[cid]
	if !ok {
		cal = cals.Primary()
	}

	// Set up query parameters
	after := time.Now().Format(time.RFC3339)

	query := dt.gcal.Events.List(cal.Item.Id)
	query = query.ShowDeleted(false).SingleEvents(true).OrderBy("startTime")
	query = query.TimeMin(after).MaxResults(n)

	var events *calendar.Events
	events, err = query.Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve the %d upcoming user's events: %v", n, err)
	}

	upcoming := make([]*AgendaItem, 0, len(events.Items))

	for _, event := range events.Items {
		item := &AgendaItem{
			Calendar: dt.calendars[cid],
			Event:    event,
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
func (dt *Daytimer) Calendars() (Calendars, error) {
	if dt.calendars == nil {
		// Fetch calendars from Google
		cals, err := dt.gcal.CalendarList.List().Do()
		if err != nil {
			return nil, err
		}

		// Create the calendars mapping
		dt.calendars = make(map[string]*Calendar)
		for _, cal := range cals.Items {
			dt.calendars[cal.Id] = &Calendar{Item: cal}
		}

		// Mark calendars active from config otherwise set primary as the
		// active calendar and dump the active calendars file.
		if err := dt.calendars.loadActive(); err != nil {
			for _, cal := range dt.calendars {
				if cal.Item.Primary {
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

// Timezone returns the timezone set on the daytimer, if one is not specified
// it searches the calendars for a timezone, then defaults to UTC.
func (dt *Daytimer) Timezone() *time.Location {

	if dt.timezone == nil {
		// Search for the first calendar with a timezone
		for _, cal := range dt.calendars {
			if cal.Item.TimeZone != "" {
				loc, err := time.LoadLocation(cal.Item.TimeZone)
				if err == nil {
					dt.timezone = loc
					break
				}
			}
		}

		// If it's still nil, then default to UTC
		if dt.timezone == nil {
			dt.timezone, _ = time.LoadLocation("UTC")
		}
	}

	return dt.timezone
}

// returns valid calendars from the list of calendar id strings. If no
// calendar id strings are specified, then it returns all active calendars.
func (dt *Daytimer) validateCalendars(calendars []string) ([]string, error) {
	// Ensure calendars are loaded
	if _, err := dt.Calendars(); err != nil {
		return nil, err
	}

	if calendars == nil || len(calendars) == 0 {
		// Use active calendars
		active := dt.calendars.Active()
		acals := make([]string, 0, len(active))
		for cid := range active {
			acals = append(acals, cid)
		}
		return acals, nil
	}

	// Validate that calendars are in the calendars list
	for _, cid := range calendars {
		_, ok := dt.calendars[cid]
		if !ok {
			return nil, fmt.Errorf("'%s' is not an available calendar", cid)
		}
	}

	// Return the original calendar list if all checks pass
	return calendars, nil
}
