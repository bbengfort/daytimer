package daytimer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"

	calendar "google.golang.org/api/calendar/v3"
)

//===========================================================================
// Agenda Types
//===========================================================================

// Agenda is a wrapper for a list of events.
type Agenda struct {
	Title    string
	Items    []*AgendaItem
	Counts   map[string]int // list of calendars and # of items per each
	Daytimer *Daytimer
}

// AgendaItem holds an event or printable item for an agenda.
// TODO: Make an interface to allow different agenda items.
type AgendaItem struct {
	Calendar  *Calendar
	Event     *calendar.Event
	startTime time.Time
	endTime   time.Time
}

//===========================================================================
// Agenda Methods
//===========================================================================

// Sort the agenda items by time
func (a *Agenda) Sort() {
	// Agenda item comparison function
	cmp := func(i, j int) bool {
		return a.Items[i].StartTime().Before(a.Items[j].StartTime())
	}

	if !sort.SliceIsSorted(a.Items, cmp) {
		sort.Slice(a.Items, cmp)
	}
}

func (a *Agenda) String() string {
	items := make([]string, 0, len(a.Items)+2)
	if a.Title != "" {
		items = append(items, a.Title)
		items = append(items, strings.Repeat("-", len(a.Title)))
	}

	for _, item := range a.Items {
		items = append(items, item.String(""))
	}

	return strings.Join(items, "\n")
}

// Count returns the number of items on the agenda
func (a *Agenda) Count() int {
	return len(a.Items)
}

// Version returns the daytimer version so that the template is updated.
func (a *Agenda) Version() string {
	return Version
}

// WriteHTML writes the HTML data to the specified path.
func (a *Agenda) WriteHTML(path string) error {
	// Create Agenda HTML message
	buffer := new(bytes.Buffer)
	template := MustLoadTemplate("templates/agenda.html")
	if err := template.Execute(buffer, &a); err != nil {
		return err
	}

	// Write the buffer to disk
	return ioutil.WriteFile(path, buffer.Bytes(), 0644)
}

//===========================================================================
// AgendaItem Methods
//===========================================================================

func (a *AgendaItem) String(timeFormat string) string {
	if timeFormat == "" {
		timeFormat = "15:04"
	}

	when := a.StartTime().Format(timeFormat)
	return fmt.Sprintf("%s %s", when, a.Event.Summary)
}

// StartTime returns the computed start time from the event
func (a *AgendaItem) StartTime() time.Time {
	var err error

	if a.startTime.IsZero() {
		if a.Event.Start.DateTime != "" {
			a.startTime, err = time.Parse(time.RFC3339, a.Event.Start.DateTime)
		} else {
			a.startTime, err = time.Parse("2006-01-02", a.Event.Start.Date)
		}
	}

	if err != nil {
		panic(err)
	}

	return a.startTime
}

// EndTime returns the computed end time from teh event
func (a *AgendaItem) EndTime() time.Time {
	var err error

	if a.endTime.IsZero() && !a.Event.EndTimeUnspecified {
		if a.Event.End.DateTime != "" {
			a.endTime, err = time.Parse(time.RFC3339, a.Event.End.DateTime)
		} else {
			a.endTime, err = time.Parse("2006-01-02", a.Event.End.Date)
		}
	}

	if err != nil {
		panic(err)
	}

	return a.endTime
}

// Color returns the calendar color
func (a *AgendaItem) Color(layer string) string {
	switch layer {
	case "bg":
		return a.Calendar.Item.BackgroundColor
	case "fg":
		return a.Calendar.Item.ForegroundColor
	default:
		return ""
	}
}

// Period returns a string representation of the time period of the event.
func (a *AgendaItem) Period() string {
	if a.Event.Start.DateTime == "" {
		return "All Day"
	}

	if !a.EndTime().IsZero() && a.EndTime().After(a.StartTime()) {
		return fmt.Sprintf("%s – %s", a.StartTime().Format("15:04"), a.EndTime().Format("15:04"))
	}

	return a.StartTime().Format("15:04")
}
