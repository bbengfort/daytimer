package daytimer

import "time"

// Layouts for various date and time serialization
const (
	DateLayout = "2006-01-02"
)

// ParseDate converts a datestamp in the form of YYYY/MM/DD into a date. If
// the string is an empty string, it returns today's date.
func ParseDate(stamp string) (time.Time, error) {
	var err error
	var ts time.Time

	if stamp == "" {
		ts = time.Now()
	} else {
		ts, err = time.Parse(DateLayout, stamp)
		if err != nil {
			return time.Time{}, err
		}
	}

	year, month, day := ts.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, ts.Location()), nil
}

// DayRange returns two timestamps in RFC3339 format, midnight of the day
// supplied and midnight of the next day (e.g. 1 day), keeping the timezone
// of the original location.
func DayRange(date time.Time) (string, string) {
	midnight := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	tomorrow := midnight.Add(time.Hour * 24)
	return midnight.Format(time.RFC3339), tomorrow.Format(time.RFC3339)
}

// Localize a time.Time to a specific location without changing any values.
// Combining multiple jobs into one, this function accepts a string of the
// time zone location and loads it directly.
func Localize(ts time.Time, loc *time.Location) time.Time {
	return time.Date(
		ts.Year(), ts.Month(), ts.Day(),
		ts.Hour(), ts.Minute(), ts.Second(), ts.Nanosecond(),
		loc,
	)
}
