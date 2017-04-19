package daytimer

import "strings"

// Agenda is a wrapper for a list of events.
type Agenda struct {
	events []string
}

func (a *Agenda) String() string {
	return strings.Join(a.events, "\n")
}
