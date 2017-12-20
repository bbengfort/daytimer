package daytimer

import (
	"fmt"
	"strings"

	calendar "google.golang.org/api/calendar/v3"
)

//===========================================================================
// Agenda Types
//===========================================================================

// Agenda is a wrapper for a list of events.
type Agenda struct {
	items []*AgendaItem
}

// AgendaItem holds an event or printable item for an agenda.
// TODO: Make an interface to allow different agenda items.
type AgendaItem struct {
	calendar *calendar.CalendarListEntry
	event    *calendar.Event
}

//===========================================================================
// Agenda Methods
//===========================================================================

func (a *Agenda) String() string {
	items := make([]string, 0, len(a.items))
	for _, item := range a.items {
		items = append(items, item.String())
	}

	return strings.Join(items, "\n")
}

//===========================================================================
// AgendaItem Methods
//===========================================================================

func (a *AgendaItem) String() string {
	var when string

	// If the DateTime is an empty string, the event is an all day event
	if a.event.Start.DateTime != "" {
		when = a.event.Start.DateTime
	} else {
		when = a.event.Start.Date
	}
	return fmt.Sprintf("%s (%s)", a.event.Summary, when)
}
