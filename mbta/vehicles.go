package mbta

import (
	"context"
	"fmt"
)

const vehiclesAPIPath = "/vehicles"

// VehicleService service handling all of the vehicle related API calls
type VehicleService service

// GetAllVehicles returns all vehicles from the mbta API
func (s *VehicleService) GetAllVehicles(config GetAllVehiclesRequestConfig) ([]Vehicle, error) {
	return s.GetAllVehiclesContext(context.Background(), config)
}

// GetAllVehiclesContext returns all vehicles from the mbta API given a context
func (s *VehicleService) GetAllVehiclesContext(ctx context.Context, config GetAllVehiclesRequestConfig) ([]Vehicle, error) {
	req, err := s.client.newRequest("GET", vehiclesAPIPath, nil)
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, err
	}

	var vehicles []Vehicle
	_, err = s.client.do(req, &vehicles)
	return vehicles, err
}

// GetVehicle returns a vehicle from the mbta API
func (s *VehicleService) GetVehicle(id string, config GetVehicleRequestConfig) (Vehicle, error) {
	return s.GetVehicleContext(context.Background(), id, config)
}

// GetVehicleContext returns a vehicle from the mbta API given a context
func (s *VehicleService) GetVehicleContext(ctx context.Context, id string, config GetVehicleRequestConfig) (Vehicle, error) {
	path := fmt.Sprintf("/%s/%s", vehiclesAPIPath, id)
	req, err := s.client.newRequest("GET", path, nil)
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)
	if err != nil {
		return Vehicle{}, err
	}

	var vehicle Vehicle
	_, err = s.client.do(req, &vehicle)
	return vehicle, err
}
