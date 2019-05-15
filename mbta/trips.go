package mbta

import (
	"context"
	"fmt"
)

const tripsAPIPath = "/trips"

// TripService service handling all of the trip related API calls
type TripService service

// GetAllTrips returns all vehicles from the mbta API
func (s *TripService) GetAllTrips(config GetAllTripsRequestConfig) ([]*Trip, error) {
	return s.GetAllTripsContext(context.Background(), config)
}

// GetAllTripsContext returns all vehicles from the mbta API given a context
func (s *TripService) GetAllTripsContext(ctx context.Context, config GetAllTripsRequestConfig) ([]*Trip, error) {
	req, err := s.client.newGETRequest(tripsAPIPath)
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, err
	}

	untypedTrips, _, err := s.client.doManyPayload(req, &Trip{})
	trips := make([]*Trip, len(untypedTrips))
	for i := 0; i < len(untypedTrips); i++ {
		trips[i] = untypedTrips[i].(*Trip)
	}
	return trips, err
}

// GetTrip returns a vehicle from the mbta API
func (s *TripService) GetTrip(id string, config GetTripRequestConfig) (*Trip, error) {
	return s.GetTripContext(context.Background(), id, config)
}

// GetTripContext returns a vehicle from the mbta API given a context
func (s *TripService) GetTripContext(ctx context.Context, id string, config GetTripRequestConfig) (*Trip, error) {
	path := fmt.Sprintf("%s/%s", tripsAPIPath, id)
	req, err := s.client.newGETRequest(path)
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, err
	}

	var trip Trip
	_, err = s.client.doSinglePayload(req, &trip)
	return &trip, err
}
