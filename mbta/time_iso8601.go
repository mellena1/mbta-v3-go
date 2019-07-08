package mbta

import (
	"fmt"
	"net/url"
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
	Now  bool // Used for when "NOW" is an option in filters
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
		parsed, err = parseISO8601TimeDateOnly(strTime)
		if err != nil {
			return err
		}
	}

	t.Time = parsed
	return nil
}

// EncodeValues implement the "github.com/google/go-querystring/query" interface for encoding
func (t *TimeISO8601) EncodeValues(key string, v *url.Values) error {
	if t.Now {
		v.Add(key, "NOW")
		return nil
	}
	v.Add(key, t.FormatOnlyDate())
	return nil
}

func parseISO8601Time(timeStr string) (time.Time, error) {
	return time.Parse(iso8601Format, timeStr)
}

func parseISO8601TimeDateOnly(timeStr string) (time.Time, error) {
	return time.Parse(iso8601FormatDateOnly, timeStr)
}

func parseISO8601TimeDateOnlySlice(timeStrSlice []string) ([]time.Time, error) {
	var timeSlice = make([]time.Time, len(timeStrSlice))
	for i, str := range timeStrSlice {
		timeSlice[i], _ = parseISO8601TimeDateOnly(str)
	}
	return timeSlice, nil
}

func timeToTimeISO8601(t time.Time) TimeISO8601 {
	return TimeISO8601{Time: t}
}

func timeSliceToTimeISO8601Slice(timeSlice []time.Time) []TimeISO8601 {
	var timeISO8601Slice = make([]TimeISO8601, len(timeSlice))
	for i, t := range timeSlice {
		timeISO8601Slice[i] = TimeISO8601{Time: t}
	}
	return timeISO8601Slice
}
