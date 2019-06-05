package mbta

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func Test_GetStop(t *testing.T) {
	id := "55"

	expected := &Stop{
		ID:                 id,
		Address:            nil,
		Description:        strPtr("Washington St @ Massachusetts Ave - Silver Line - Dudley"),
		Latitude:           42.336361,
		LocationType:       StopLocationStop,
		Longitude:          -71.077214,
		Name:               "Washington St @ Massachusetts Ave",
		PlatformCode:       nil,
		PlatformName:       strPtr("Dudley"),
		WheelchairBoarding: WheelchairBoardingAccessible,
		ParentStation:      nil,
	}
	server := httptest.NewServer(handlerForServer(t, fmt.Sprintf("%s/%s", stopsAPIPath, id)))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Stops.GetStop(id, &GetStopRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}

func Test_GetAllStops(t *testing.T) {
	expected := []*Stop{
		&Stop{
			ID:                 "111146",
			Address:            nil,
			Description:        nil,
			Latitude:           42.303676,
			LocationType:       StopLocationStop,
			Longitude:          -70.919766,
			Name:               "Pemberton Point",
			PlatformCode:       nil,
			PlatformName:       nil,
			WheelchairBoarding: WheelchairBoardingNoInfo,
			ParentStation:      nil,
		},
		&Stop{
			ID:                 "9172",
			Address:            nil,
			Description:        nil,
			Latitude:           42.416769,
			LocationType:       StopLocationStop,
			Longitude:          -71.105122,
			Name:               "116 Riverside Ave",
			PlatformCode:       nil,
			PlatformName:       nil,
			WheelchairBoarding: WheelchairBoardingNoInfo,
			ParentStation:      nil,
		},
	}
	server := httptest.NewServer(handlerForServer(t, stopsAPIPath))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Stops.GetAllStops(&GetAllStopsRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}
