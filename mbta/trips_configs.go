package mbta

import (
	"net/http"
)

// GetAllTripsSortByType all of the possible ways to sort by for a GetAllTrips request
type GetAllTripsSortByType string

const (
	GetAllTripsSortByBikesAllowedAscending  GetAllTripsSortByType = "bikes_allowed"
	GetAllTripsSortByBikesAllowedDescending GetAllTripsSortByType = "-bikes_allowed"
	// TODO: Finish these
)

// GetAllTripsRequestConfig extra options for the GetAllTrips request
type GetAllTripsRequestConfig struct {
	// TODO: Fill out fields
}

func (config *GetAllTripsRequestConfig) addHTTPParamsToRequest(req *http.Request) {
	// TODO: implement
	q := req.URL.Query()

	req.URL.RawQuery = q.Encode()
}

// GetTripRequestConfig extra options for the GetTrip request
type GetTripRequestConfig struct {
	// TODO: Fill out fields
}

func (config *GetTripRequestConfig) addHTTPParamsToRequest(req *http.Request) {
	// TODO: implement
	q := req.URL.Query()

	req.URL.RawQuery = q.Encode()
}
