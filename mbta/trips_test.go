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
		WheelchairAccessible: WheelchairBoardingACCESSIBLE,
		Headsign:             "Ashmont",
		DirectionID:          0,
		BikesAllowed:         BikesAllowedNOINFO,
		BlockID:              "S931_-5-0-L-0-BraintreeQuincyCenter",
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
			WheelchairAccessible: WheelchairBoardingACCESSIBLE,
			Headsign:             "Ashmont",
			DirectionID:          0,
			BikesAllowed:         BikesAllowedNOINFO,
			BlockID:              "S931_-4-0-L-0-BraintreeQuincyCenter",
		},
		&Trip{
			ID:                   "40119998-L",
			Name:                 "",
			WheelchairAccessible: WheelchairBoardingACCESSIBLE,
			Headsign:             "Ashmont",
			DirectionID:          0,
			BikesAllowed:         BikesAllowedNOINFO,
			BlockID:              "S931_-4-0-L",
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
