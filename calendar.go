package daytimer

import (
	"fmt"
	"path/filepath"
	"strings"

	calendar "google.golang.org/api/calendar/v3"
)

//===========================================================================
// Calendar Structures
//===========================================================================

// Calendars is a collection of Calendar items
type Calendars map[string]*Calendar

// Calendar is a wrapper for a Google Calendar List Entry that provides extra
// methods for use with the daytimer applciation.
type Calendar struct {
	Item   *calendar.CalendarListEntry // The Google API Calendar Entry
	active bool                        // If the user has marked this calendar as active
}

//===========================================================================
// Calendars Methods
//===========================================================================

// Primary returns the calendar listed as primary or nil
func (c Calendars) Primary() *Calendar {
	for _, cal := range c {
		if cal.Item.Primary {
			return cal
		}
	}
	return nil
}

// Active filters out all calendars that are not active, returning a new
// Calendars map collection with the active items.
func (c Calendars) Active() Calendars {
	cals := make(Calendars)
	for k, cal := range c {
		if cal.active {
			cals[k] = cal
		}
	}
	return cals
}

// Activate a specific calendar by id. Returns an error if the calendar is
// not in the collection. This method also dumps active calendars to disk.
func (c Calendars) Activate(cid string) error {
	cal, ok := c[cid]
	if !ok {
		return fmt.Errorf("'%s' is not a valid calendar id", cid)
	}

	cal.active = true
	return c.dumpActive()
}

func (c Calendars) String() string {
	reprs := make([]string, 0, len(c))
	for _, item := range c {
		reprs = append(reprs, item.String())
	}

	return strings.Join(reprs, "\n")
}

// Edit the calendars file.
func (c Calendars) Edit() error {
	path, err := c.configPath()
	if err != nil {
		return err
	}

	if err := EditFile(path); err != nil {
		return fmt.Errorf("%s: please edit %s directly", err, path)
	}

	return nil
}

// Get the path to the active calendars config
func (c Calendars) configPath() (string, error) {
	confDir, err := configDirectory()
	if err != nil {
		return "", err
	}

	return filepath.Join(confDir, "calendars.txt"), nil
}

// Load all active calendars from the config.
func (c Calendars) loadActive() error {
	path, err := c.configPath()
	if err != nil {
		return err
	}

	reader, err := readLines(path, true)
	if err != nil {
		return err
	}

	for line := range reader {
		line = strings.TrimSpace(line)
		cal, ok := c[line]
		if ok {
			cal.active = true
		}
	}

	return nil
}

// Dump all active calendars to the config.
func (c Calendars) dumpActive() error {
	path, err := c.configPath()
	if err != nil {
		return err
	}

	active := make([]string, 0, len(c))
	for id, cal := range c {
		if cal.active {
			active = append(active, id)
		}
	}

	return writeLines(active, path)
}

//===========================================================================
// Calendar Methods
//===========================================================================

func (c *Calendar) String() string {
	if c.Item == nil {
		return ""
	}

	title := c.Summary()

	var repr string
	if c.Item.Location != "" {
		repr = fmt.Sprintf("%s (%s) in %s", title, c.Item.Id, c.Item.Location)
	} else {
		repr = fmt.Sprintf("%s (%s)", title, c.Item.Id)
	}

	if c.Item.Primary {
		repr = "☆ " + repr
	}

	return repr
}

// Summary returns the summary of the calendar
func (c *Calendar) Summary() string {
	if c.Item == nil {
		return ""
	}

	if c.Item.SummaryOverride != "" {
		return c.Item.SummaryOverride
	} else if c.Item.Summary != "" {
		return c.Item.Summary
	}

	return "Google Calendar"
}
