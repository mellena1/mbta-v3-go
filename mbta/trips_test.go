package mbta

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func Test_GetTrip(t *testing.T) {
	id := "40119999-BraintreeQuincyCenterL"

	expected := &Trip{
		ID:                   id,
		Name:                 "",
		WheelchairAccessible: WheelchairBoardingAccessible,
		Headsign:             "Ashmont",
		DirectionID:          0,
		BikesAllowed:         BikesAllowedNoInfo,
		BlockID:              "S931_-5-0-L-0-BraintreeQuincyCenter",
		Route:                &Route{ID: "Red"},
		RoutePattern:         &RoutePattern{ID: "Red-1-0"},
		Service:              &Service{ID: "RTL22019-hms29016-Saturday-01-BraintreeQuincyCenterL"},
		Shape:                &Shape{ID: "931_0009"},
	}
	server := httptest.NewServer(handlerForServer(t, fmt.Sprintf("%s/%s", tripsAPIPath, id)))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Trips.GetTrip(id, GetTripRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}

func Test_GetAllTrips(t *testing.T) {
	expected := []*Trip{
		&Trip{
			ID:                   "40119998-BraintreeQuincyCenterL",
			Name:                 "",
			WheelchairAccessible: WheelchairBoardingAccessible,
			Headsign:             "Ashmont",
			DirectionID:          0,
			BikesAllowed:         BikesAllowedNoInfo,
			BlockID:              "S931_-4-0-L-0-BraintreeQuincyCenter",
			Route:                &Route{ID: "Red"},
			RoutePattern:         &RoutePattern{ID: "Red-1-0"},
			Service:              &Service{ID: "RTL22019-hms29016-Saturday-01-BraintreeQuincyCenterL"},
			Shape:                &Shape{ID: "931_0009"},
		},
		&Trip{
			ID:                   "40119998-L",
			Name:                 "",
			WheelchairAccessible: WheelchairBoardingAccessible,
			Headsign:             "Ashmont",
			DirectionID:          0,
			BikesAllowed:         BikesAllowedNoInfo,
			BlockID:              "S931_-4-0-L",
			Route:                &Route{ID: "Red"},
			RoutePattern:         &RoutePattern{ID: "Red-1-0"},
			Service:              &Service{ID: "RTL22019-hms29016-Saturday-01-L"},
			Shape:                &Shape{ID: "931_0009"},
		},
	}
	server := httptest.NewServer(handlerForServer(t, tripsAPIPath))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Trips.GetAllTrips(GetAllTripsRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}
