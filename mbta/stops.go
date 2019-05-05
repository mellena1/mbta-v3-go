package mbta

import (
	"context"
	"fmt"
)

// StopService service handling all of the stop related API calls
type StopService service

// GetAllStops returns all stops from the mbta API
func (s *StopService) GetAllStops() ([]Stop, error) {
	return s.GetAllStopsContext(context.Background())
}

// GetAllStopsContext returns all stops from the mbta API given a context
func (s *StopService) GetAllStopsContext(ctx context.Context) ([]Stop, error) {
	req, err := s.client.newRequest("GET", "/stops", nil)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, err
	}

	var data struct {
		Stops []Stop `json:"data"`
	}
	_, err = s.client.do(req, &data)
	return data.Stops, err
}

// GetStop returns a stop from the mbta API
func (s *StopService) GetStop(id string) (Stop, error) {
	return s.GetStopContext(context.Background(), id)
}

// GetStopContext returns a stop from the mbta API given a context
func (s *StopService) GetStopContext(ctx context.Context, id string) (Stop, error) {
	path := fmt.Sprintf("/stops/%s", id)
	req, err := s.client.newRequest("GET", path, nil)
	req = req.WithContext(ctx)
	if err != nil {
		return Stop{}, err
	}

	var data struct {
		Stop Stop `json:"data"`
	}
	_, err = s.client.do(req, &data)
	return data.Stop, err
}
