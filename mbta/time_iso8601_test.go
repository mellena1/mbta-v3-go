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

func Test_TimeISO8601_UnmarshalJSON(t *testing.T) {
	timeLocation, _ := time.LoadLocation("America/New_York")
	tests := []struct {
		val      string
		expected time.Time
		valid    bool
	}{
		{"2019-05-14T16:05:53-04:00", time.Date(2019, time.May, 14, 16, 05, 53, 0, timeLocation), true},
		{"\"2019-05-14T16:05:53-04:00\"", time.Date(2019, time.May, 14, 16, 05, 53, 0, timeLocation), true},
		{"2019-05-14T-----16:05:53-04:00", time.Date(2019, time.May, 14, 16, 05, 53, 0, timeLocation), false},
	}

	tISO := TimeISO8601{}
	for _, test := range tests {
		err := tISO.UnmarshalJSON([]byte(test.val))
		if test.valid {
			ok(t, err)
			assert(t, test.expected.Equal(tISO.Time), "times not equal")
		} else {
			assert(t, err != nil, "this should have errored")
		}
	}
}
