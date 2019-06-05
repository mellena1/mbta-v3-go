package mbta

import (
	"net/http/httptest"
	"testing"

	"golang.org/x/xerrors"
)

func Test_GetAllSchedules(t *testing.T) {
	parsedArrivalTime1, _ := parseISO8601Time("2019-06-03T05:01:00-04:00")
	parsedDepartureTime1, _ := parseISO8601Time("2019-06-03T05:01:00-04:00")

	parsedArrivalTime2, _ := parseISO8601Time("2019-06-03T05:02:00-04:00")
	parsedDepartureTime2, _ := parseISO8601Time("2019-06-03T05:02:00-04:00")
	expected := []*Schedule{
		&Schedule{
			ID:            "schedule-39988449-20:30-NewtonHighlandsRiverside-70238-180",
			ArrivalTime:   timeToTimeISO8601(parsedArrivalTime1),
			DepartureTime: timeToTimeISO8601(parsedDepartureTime1),
			DirectionID:   1,
			DropOffType:   SchedulePickupNotAvailable,
			PickupType:    SchedulePickupRegular,
			StopSequence:  180,
			Timepoint:     ScheduleTimepointEstimates,
			Route:         &Route{ID: "Green-C"},
			Stop:          &Stop{ID: "70238"},
			Trip:          &Trip{ID: "39988449-20:30-NewtonHighlandsRiverside"},
		},
		&Schedule{
			ID:            "schedule-39988449-20:30-NewtonHighlandsRiverside-70236-190",
			ArrivalTime:   timeToTimeISO8601(parsedArrivalTime2),
			DepartureTime: timeToTimeISO8601(parsedDepartureTime2),
			DirectionID:   1,
			DropOffType:   SchedulePickupRegular,
			PickupType:    SchedulePickupRegular,
			StopSequence:  190,
			Timepoint:     ScheduleTimepointEstimates,
			Route:         &Route{ID: "Green-C"},
			Stop:          &Stop{ID: "70236"},
			Trip:          &Trip{ID: "39988449-20:30-NewtonHighlandsRiverside"},
		},
	}
	opts := &GetAllSchedulesRequestConfig{FilterRouteIDs: []string{"Green-C"}}
	fullPath, _ := addOptions(schedulesAPIPath, opts)
	server := httptest.NewServer(handlerForServer(t, fullPath))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Schedules.GetAllSchedules(opts)
	ok(t, err)
	equals(t, expected, actual)
}

func Test_GetAllSchedulesFail(t *testing.T) {
	server := httptest.NewServer(handlerForServer(t, schedulesAPIPath))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	_, _, err := mbtaClient.Schedules.GetAllSchedules(&GetAllSchedulesRequestConfig{})
	equals(t, true, xerrors.Is(err, ErrInvalidConfig))
}
