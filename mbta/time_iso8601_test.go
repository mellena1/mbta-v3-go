package mbta

import (
	"testing"
	"time"
)

func Test_TimeISO8601_Format(t *testing.T) {
	testTime := TimeISO8601{
		Time: time.Date(1991, time.August, 13, 12, 01, 02, 0, time.FixedZone("UTC-8", -8*60*60)),
	}
	expected := "1991-08-13T12:01:02-08:00"
	actual := testTime.Format()

	equals(t, expected, actual)
}

func Test_TimeISO8601_FormatOnlyDate(t *testing.T) {
	testTime := TimeISO8601{
		Time: time.Date(1991, time.August, 13, 12, 01, 02, 0, time.FixedZone("UTC-8", -8*60*60)),
	}
	expected := "1991-08-13"
	actual := testTime.FormatOnlyDate()

	equals(t, expected, actual)
}

func Test_TimeISO8601_MarshalJSON(t *testing.T) {
	testTime := TimeISO8601{
		Time: time.Date(1991, time.August, 13, 12, 01, 02, 0, time.FixedZone("UTC-8", -8*60*60)),
	}
	expected := "\"1991-08-13T12:01:02-08:00\""
	actual, err := testTime.MarshalJSON()
	ok(t, err)
	equals(t, expected, string(actual))
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

func Test_parseISO8601Time(t *testing.T) {
	timeLocation, _ := time.LoadLocation("America/New_York")
	expected := timeToTimeISO8601(time.Date(2019, time.May, 14, 16, 05, 53, 0, timeLocation))

	actual, _ := parseISO8601Time("2019-05-14T16:05:53-04:00")

	assert(t, expected.Time.Equal(actual), "times not equal")
}
