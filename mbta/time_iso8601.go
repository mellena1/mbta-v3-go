package mbta

import (
	"fmt"
	"strings"
	"time"
)

const (
	iso8601Format         = "2006-01-02T15:04:05-07:00"
	iso8601FormatDateOnly = "2006-01-02"
)

// TimeISO8601 wrapper for a time.Time struct so that the Unmarshal works
type TimeISO8601 struct {
	Time time.Time
}

// Format the time as ISO8601
func (t *TimeISO8601) Format() string {
	return t.Time.Format(iso8601Format)
}

// FormatOnlyDate format the time as ISO8601 with only the date
func (t *TimeISO8601) FormatOnlyDate() string {
	return t.Time.Format(iso8601FormatDateOnly)
}

// MarshalJSON marshal time.Time as ISO8601
func (t *TimeISO8601) MarshalJSON() ([]byte, error) {
	strTime := fmt.Sprintf("\"%s\"", t.Format())
	return []byte(strTime), nil
}

// UnmarshalJSON unmarshal time.Time as ISO8601
func (t *TimeISO8601) UnmarshalJSON(b []byte) error {
	strTime := strings.Trim(string(b), "\"")
	parsed, err := parseISO8601Time(strTime)
	if err != nil {
		return err
	}

	t.Time = parsed
	return nil
}

func parseISO8601Time(timeStr string) (time.Time, error) {
	return time.Parse(iso8601Format, timeStr)
}

func timeToTimeISO8601(t time.Time) TimeISO8601 {
	return TimeISO8601{Time: t}
}
