package mbta

import (
	"context"
	"fmt"
)

const stopsAPIPath = "/stops"

// StopService service handling all of the stop related API calls
type StopService service

// GetAllStops returns all stops from the mbta API
func (s *StopService) GetAllStops(config GetAllStopsRequestConfig) ([]*Stop, error) {
	return s.GetAllStopsContext(context.Background(), config)
}

// GetAllStopsContext returns all stops from the mbta API given a context
func (s *StopService) GetAllStopsContext(ctx context.Context, config GetAllStopsRequestConfig) ([]*Stop, error) {
	req, err := s.client.newGETRequest(stopsAPIPath)
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, err
	}

	untypedStops, _, err := s.client.doManyPayload(req, Stop{})
	stops := make([]*Stop, len(untypedStops))
	for i := 0; i < len(untypedStops); i++ {
		stops[i] = untypedStops[i].(*Stop)
	}
	return stops, err
}

// GetStop returns a stop from the mbta API
func (s *StopService) GetStop(id string, config GetStopRequestConfig) (*Stop, error) {
	return s.GetStopContext(context.Background(), id, config)
}

// GetStopContext returns a stop from the mbta API given a context
func (s *StopService) GetStopContext(ctx context.Context, id string, config GetStopRequestConfig) (*Stop, error) {
	path := fmt.Sprintf("%s/%s", stopsAPIPath, id)
	req, err := s.client.newGETRequest(path)
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, err
	}

	var stop Stop
	_, err = s.client.doSinglePayload(req, &stop)
	return &stop, err
}
