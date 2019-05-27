package mbta

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestGetRoute(t *testing.T) {
	expected := &Route{
		ID:                    "66",
		Color:                 "FFC72C",
		Description:           "Key Bus",
		DirectionDestinations: []string{"Harvard", "Dudley"},
		DirectionNames:        []string{"Outbound", "Inbound"},
		LongName:              "Harvard - Dudley via Allston",
		SortOrder:             50660,
		TextColor:             "000000",
		Type:                  RouteTypeBus,
		ShortName:             "66",
	}
	server := httptest.NewServer(handlerForServer(t, fmt.Sprintf("%s/%s", routesAPIPath, "66")))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Routes.GetRoute("66", GetRouteRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}

func TestGetAllRoutes(t *testing.T) {
	expected := []*Route{
		&Route{
			ID:                    "66",
			Color:                 "FFC72C",
			Description:           "Key Bus",
			DirectionDestinations: []string{"Harvard", "Dudley"},
			DirectionNames:        []string{"Outbound", "Inbound"},
			LongName:              "Harvard - Dudley via Allston",
			SortOrder:             50660,
			TextColor:             "000000",
			Type:                  RouteTypeBus,
			ShortName:             "66",
		},
		&Route{
			ID:                    "39",
			Color:                 "FFC72C",
			Description:           "Key Bus",
			DirectionDestinations: []string{"Forest Hills", "Back Bay Station"},
			DirectionNames:        []string{"Outbound", "Inbound"},
			LongName:              "Forest Hills - Back Bay Station",
			SortOrder:             50390,
			TextColor:             "000000",
			Type:                  RouteTypeBus,
			ShortName:             "39",
		},
	}
	server := httptest.NewServer(handlerForServer(t, routesAPIPath))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Routes.GetAllRoutes(GetAllRoutesRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}
