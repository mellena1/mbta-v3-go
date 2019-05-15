package mbta

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func Test_GetStop(t *testing.T) {
	expected := &Stop{
		ID:                 "55",
		Address:            nil,
		Description:        strPtr("Washington St @ Massachusetts Ave - Silver Line - Dudley"),
		Latitute:           42.336361,
		LocationType:       StopLocationSTOP,
		Longitude:          -71.077214,
		Name:               "Washington St @ Massachusetts Ave",
		PlatformCode:       nil,
		PlatformName:       strPtr("Dudley"),
		WheelchairBoarding: WheelchairBoardingACCESSIBLE,
		ParentStation:      nil,
	}
	server := httptest.NewServer(handlerForServer(t, fmt.Sprintf("%s/%s", stopsAPIPath, "55")))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Stops.GetStop("55", GetStopRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}

func Test_GetAllStops(t *testing.T) {
	expected := []*Stop{
		&Stop{
			ID:                 "111146",
			Address:            nil,
			Description:        nil,
			Latitute:           42.303676,
			LocationType:       StopLocationSTOP,
			Longitude:          -70.919766,
			Name:               "Pemberton Point",
			PlatformCode:       nil,
			PlatformName:       nil,
			WheelchairBoarding: WheelchairBoardingNOINFO,
			ParentStation:      nil,
		},
		&Stop{
			ID:                 "9172",
			Address:            nil,
			Description:        nil,
			Latitute:           42.416769,
			LocationType:       StopLocationSTOP,
			Longitude:          -71.105122,
			Name:               "116 Riverside Ave",
			PlatformCode:       nil,
			PlatformName:       nil,
			WheelchairBoarding: WheelchairBoardingNOINFO,
			ParentStation:      nil,
		},
	}
	server := httptest.NewServer(handlerForServer(t, stopsAPIPath))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Stops.GetAllStops(GetAllStopsRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}
