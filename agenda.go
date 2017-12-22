package daytimer

import (
	"fmt"
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
	Title string
	Items []*AgendaItem
}

// AgendaItem holds an event or printable item for an agenda.
// TODO: Make an interface to allow different agenda items.
type AgendaItem struct {
	Calendar  *Calendar
	Event     *calendar.Event
	startTime time.Time
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
