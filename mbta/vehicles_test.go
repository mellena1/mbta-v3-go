package mbta

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func Test_GetVehicle(t *testing.T) {
	parsedTime, _ := parseISO8601Time("2019-05-14T16:05:53-04:00")
	expected := &Vehicle{
		ID:                  "y1772",
		Bearing:             270.0,
		CurrentStatus:       InTransitTo,
		CurrentStopSequence: 1,
		DirectionID:         1,
		Label:               "1772",
		Latitute:            42.349491119384766,
		Longitude:           -71.07652282714844,
		Speed:               nil,
		UpdatedAt:           timeToTimeISO8601(parsedTime),
		Stop:                &Stop{ID: "178"},
		Trip:                &Trip{ID: "39915343"},
	}
	server := httptest.NewServer(handlerForServer(t, fmt.Sprintf("%s/%s", vehiclesAPIPath, "y1772")))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, err := mbtaClient.Vehicles.GetVehicle("y1772", GetVehicleRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}

func Test_GetAllVehicles(t *testing.T) {
	parsedTime1, _ := parseISO8601Time("2019-05-14T17:25:37-04:00")
	parsedTime2, _ := parseISO8601Time("2019-05-14T17:25:36-04:00")
	expected := []*Vehicle{
		&Vehicle{
			ID:                  "y1772",
			Bearing:             194.0,
			CurrentStatus:       InTransitTo,
			CurrentStopSequence: 12,
			DirectionID:         1,
			Label:               "1772",
			Latitute:            42.335472106933594,
			Longitude:           -71.0453109741211,
			Speed:               nil,
			UpdatedAt:           timeToTimeISO8601(parsedTime1),
			Stop:                &Stop{ID: "46"},
			Trip:                &Trip{ID: "39915358"},
		},
		&Vehicle{
			ID:                  "y1869",
			Bearing:             231.0,
			CurrentStatus:       InTransitTo,
			CurrentStopSequence: 24,
			DirectionID:         1,
			Label:               "1869",
			Latitute:            42.331825256347656,
			Longitude:           -71.07601165771484,
			Speed:               nil,
			UpdatedAt:           timeToTimeISO8601(parsedTime2),
			Stop:                &Stop{ID: "10100"},
			Trip:                &Trip{ID: "39914092"},
		},
	}
	server := httptest.NewServer(handlerForServer(t, vehiclesAPIPath))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, err := mbtaClient.Vehicles.GetAllVehicles(GetAllVehiclesRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}
