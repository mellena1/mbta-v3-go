package mbta

import (
	"context"
	"fmt"
	"net/http"
)

const tripsAPIPath = "/trips"

// TripService service handling all of the trip related API calls
type TripService service

// GetAllTrips returns all vehicles from the mbta API
func (s *TripService) GetAllTrips(config GetAllTripsRequestConfig) ([]*Trip, *http.Response, error) {
	return s.GetAllTripsContext(context.Background(), config)
}

// GetAllTripsContext returns all vehicles from the mbta API given a context
func (s *TripService) GetAllTripsContext(ctx context.Context, config GetAllTripsRequestConfig) ([]*Trip, *http.Response, error) {
	req, err := s.client.newGETRequest(tripsAPIPath)
	if err != nil {
		return nil, nil, err
	}
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)

	untypedTrips, resp, err := s.client.doManyPayload(req, &Trip{})
	trips := make([]*Trip, len(untypedTrips))
	for i := 0; i < len(untypedTrips); i++ {
		trips[i] = untypedTrips[i].(*Trip)
	}
	return trips, resp, err
}

// GetTrip returns a vehicle from the mbta API
func (s *TripService) GetTrip(id string, config GetTripRequestConfig) (*Trip, *http.Response, error) {
	return s.GetTripContext(context.Background(), id, config)
}

// GetTripContext returns a vehicle from the mbta API given a context
func (s *TripService) GetTripContext(ctx context.Context, id string, config GetTripRequestConfig) (*Trip, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", tripsAPIPath, id)
	req, err := s.client.newGETRequest(path)
	if err != nil {
		return nil, nil, err
	}
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)

	var trip Trip
	resp, err := s.client.doSinglePayload(req, &trip)
	return &trip, resp, err
}
