package mbta

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func Test_GetService(t *testing.T) {
	addedDates, _ := parseISO8601TimeDateOnlySlice([]string{"2019-05-28", "2019-05-29", "2019-05-30", "2019-05-31"})
	endDate, _ := parseISO8601TimeDateOnly("2019-06-21")
	startDate, _ := parseISO8601TimeDateOnly("2019-05-27")
	removedDates, _ := parseISO8601TimeDateOnlySlice([]string{"2019-05-27"})
	expected := &Service{
		ID:                 "BUS22019-hbb29011-Weekday-02",
		AddedDates:         timeSliceToTimeISO8601Slice(addedDates),
		AddedDatesNotes:    []*string{nil, nil, nil, nil},
		Description:        "Weekday schedule",
		EndDate:            timeToTimeISO8601(endDate),
		RemovedDates:       timeSliceToTimeISO8601Slice(removedDates),
		RemovedDatesNotes:  []string{"Memorial Day"},
		ScheduleName:       "Weekday",
		ScheduleType:       "Weekday",
		ScheduleTypicality: 1,
		StartDate:          timeToTimeISO8601(startDate),
		ValidDays:          []Weekday{Monday, Tuesday, Wednesday, Thursday, Friday},
	}
	server := httptest.NewServer(handlerForServer(t, fmt.Sprintf("%s/%s", servicesAPIPath, "BUS22019-hbb29011-Weekday-02")))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Services.GetService("BUS22019-hbb29011-Weekday-02", &GetServiceRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}

func Test_GetAllServices(t *testing.T) {
	addedDates1, _ := parseISO8601TimeDateOnlySlice([]string{"2019-05-28", "2019-05-29", "2019-05-30", "2019-05-31"})
	endDate1, _ := parseISO8601TimeDateOnly("2019-06-21")
	startDate1, _ := parseISO8601TimeDateOnly("2019-05-27")
	removedDates1, _ := parseISO8601TimeDateOnlySlice([]string{"2019-05-27"})

	addedDates2, _ := parseISO8601TimeDateOnlySlice([]string{"2019-06-01"})
	endDate2, _ := parseISO8601TimeDateOnly("2019-06-22")
	startDate2, _ := parseISO8601TimeDateOnly("2019-05-27")
	expected := []*Service{
		&Service{
			ID:                 "BUS22019-hbb29011-Weekday-02",
			AddedDates:         timeSliceToTimeISO8601Slice(addedDates1),
			AddedDatesNotes:    []*string{nil, nil, nil, nil},
			Description:        "Weekday schedule",
			EndDate:            timeToTimeISO8601(endDate1),
			RemovedDates:       timeSliceToTimeISO8601Slice(removedDates1),
			RemovedDatesNotes:  []string{"Memorial Day"},
			ScheduleName:       "Weekday",
			ScheduleType:       "Weekday",
			ScheduleTypicality: 1,
			StartDate:          timeToTimeISO8601(startDate1),
			ValidDays:          []Weekday{Monday, Tuesday, Wednesday, Thursday, Friday},
		},
		&Service{
			ID:                 "BUS22019-hbb29016-Saturday-02",
			AddedDates:         timeSliceToTimeISO8601Slice(addedDates2),
			AddedDatesNotes:    []*string{nil},
			Description:        "Saturday schedule",
			EndDate:            timeToTimeISO8601(endDate2),
			RemovedDates:       nil,
			RemovedDatesNotes:  []string{},
			ScheduleName:       "Saturday",
			ScheduleType:       "Saturday",
			ScheduleTypicality: 1,
			StartDate:          timeToTimeISO8601(startDate2),
			ValidDays:          []Weekday{Saturday},
		},
	}
	server := httptest.NewServer(handlerForServer(t, servicesAPIPath))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Services.GetAllServices(&GetAllServicesRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}
