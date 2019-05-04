package mbta

import (
	"context"
	"fmt"
)

// GetAllVehicles returns all vehicles from the mbta API
func (c *Client) GetAllVehicles() ([]Vehicle, error) {
	return c.GetAllVehiclesContext(context.Background())
}

// GetAllVehiclesContext returns all vehicles from the mbta API given a context
func (c *Client) GetAllVehiclesContext(ctx context.Context) ([]Vehicle, error) {
	req, err := c.newRequest("GET", "/vehicles", nil)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, err
	}

	var data struct {
		Vehicles []Vehicle `json:"data"`
	}
	_, err = c.do(req, &data)
	return data.Vehicles, err
}

// GetVehicle returns a vehicle from the mbta API
func (c *Client) GetVehicle(id string) (Vehicle, error) {
	return c.GetVehicleContext(context.Background(), id)
}

// GetVehicleContext returns a vehicle from the mbta API given a context
func (c *Client) GetVehicleContext(ctx context.Context, id string) (Vehicle, error) {
	path := fmt.Sprintf("/vehicles/%s", id)
	req, err := c.newRequest("GET", path, nil)
	req = req.WithContext(ctx)
	if err != nil {
		return Vehicle{}, err
	}

	var data struct {
		Vehicle Vehicle `json:"data"`
	}
	_, err = c.do(req, &data)
	return data.Vehicle, err
}
