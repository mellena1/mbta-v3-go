package mbta

import (
	"strings"
	"time"
)

const iso8601Format = "2006-01-02T15:04:05-07:00"

// TimeISO8601 wrapper for a time.Time struct so that the Unmarshal works
type TimeISO8601 struct {
	Time time.Time
}

// UnmarshalJSON unmarshal time.Time as ISO8601
func (t *TimeISO8601) UnmarshalJSON(b []byte) error {
	strTime := strings.Trim(string(b), "\"")
	parsed, err := time.Parse(time.RFC3339, strTime)
	if err != nil {
		return err
	}

	t.Time = parsed
	return nil
}
