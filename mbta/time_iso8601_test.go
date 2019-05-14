package mbta

import (
	"testing"
	"time"
)

func Test_parseISO8601Time(t *testing.T) {
	timeLocation, _ := time.LoadLocation("America/New_York")
	expected := timeToTimeISO8601(time.Date(2019, time.May, 14, 16, 05, 53, 0, timeLocation))

	actual, _ := parseISO8601Time("2019-05-14T16:05:53-04:00")

	assert(t, expected.Time.Equal(actual), "times not equal")
}
