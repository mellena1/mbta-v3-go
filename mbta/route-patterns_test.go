package mbta

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestGetRoutePattern(t *testing.T) {
	expected := &RoutePattern{
		ID:          "Mattapan-_-0",
		DirectionID: 0,
		Name:        "Mattapan",
		SortOrder:   10011000,
		TimeDesc:    nil,
		Typicality:  RoutePatternTypicalityTypical,
		RepresentativeTrip: &Trip{
			ID: "39923059",
		},
		Route: &Route{
			ID: "Mattapan",
		},
	}
	server := httptest.NewServer(handlerForServer(t, fmt.Sprintf("%s/%s", routesPatternsAPIPath, "Mattapan-_-0")))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.RoutePatterns.GetRoutePattern("Mattapan-_-0", &GetRoutePatternRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}

func TestGetAllRoutePatterns(t *testing.T) {
	expected := []*RoutePattern{
		&RoutePattern{
			ID:          "Red-1-0",
			DirectionID: 0,
			Name:        "Ashmont",
			SortOrder:   10010051,
			TimeDesc:    nil,
			Typicality:  RoutePatternTypicalityTypical,
			RepresentativeTrip: &Trip{
				ID: "40132582-L",
			},
			Route: &Route{
				ID: "Red",
			},
		},
		&RoutePattern{
			ID:          "Red-3-0",
			DirectionID: 0,
			Name:        "Braintree",
			SortOrder:   10010052,
			TimeDesc:    nil,
			Typicality:  RoutePatternTypicalityTypical,
			RepresentativeTrip: &Trip{
				ID: "40132593-L",
			},
			Route: &Route{
				ID: "Red",
			},
		},
	}
	server := httptest.NewServer(handlerForServer(t, routesPatternsAPIPath))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.RoutePatterns.GetAllRoutePatterns(&GetAllRoutePatternsRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}
