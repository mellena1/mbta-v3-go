package mbta

import (
	"context"
	"fmt"
)

// GetAllStops returns all stops from the mbta API
func (c *Client) GetAllStops() ([]Stop, error) {
	return c.GetAllStopsContext(context.Background())
}

// GetAllStopsContext returns all stops from the mbta API given a context
func (c *Client) GetAllStopsContext(ctx context.Context) ([]Stop, error) {
	req, err := c.newRequest("GET", "/stops", nil)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, err
	}

	var data struct {
		Stops []Stop `json:"data"`
	}
	_, err = c.do(req, &data)
	return data.Stops, err
}

// GetStop returns a stop from the mbta API
func (c *Client) GetStop(id string) (Stop, error) {
	return c.GetStopContext(context.Background(), id)
}

// GetStopContext returns a stop from the mbta API given a context
func (c *Client) GetStopContext(ctx context.Context, id string) (Stop, error) {
	path := fmt.Sprintf("/stops/%s", id)
	req, err := c.newRequest("GET", path, nil)
	req = req.WithContext(ctx)
	if err != nil {
		return Stop{}, err
	}

	var data struct {
		Stop Stop `json:"data"`
	}
	_, err = c.do(req, &data)
	return data.Stop, err
}
