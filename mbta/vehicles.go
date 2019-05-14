package mbta

import (
	"context"
	"fmt"
)

const vehiclesAPIPath = "/vehicles"

// VehicleService service handling all of the vehicle related API calls
type VehicleService service

// GetAllVehicles returns all vehicles from the mbta API
func (s *VehicleService) GetAllVehicles(config GetAllVehiclesRequestConfig) ([]*Vehicle, error) {
	return s.GetAllVehiclesContext(context.Background(), config)
}

// GetAllVehiclesContext returns all vehicles from the mbta API given a context
func (s *VehicleService) GetAllVehiclesContext(ctx context.Context, config GetAllVehiclesRequestConfig) ([]*Vehicle, error) {
	req, err := s.client.newGETRequest(vehiclesAPIPath)
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, err
	}

	untypedVehicles, _, err := s.client.doManyPayload(req, &Vehicle{})
	vehicles := make([]*Vehicle, len(untypedVehicles))
	for i := 0; i < len(untypedVehicles); i++ {
		vehicles[i] = untypedVehicles[i].(*Vehicle)
	}
	return vehicles, err
}

// GetVehicle returns a vehicle from the mbta API
func (s *VehicleService) GetVehicle(id string, config GetVehicleRequestConfig) (*Vehicle, error) {
	return s.GetVehicleContext(context.Background(), id, config)
}

// GetVehicleContext returns a vehicle from the mbta API given a context
func (s *VehicleService) GetVehicleContext(ctx context.Context, id string, config GetVehicleRequestConfig) (*Vehicle, error) {
	path := fmt.Sprintf("%s/%s", vehiclesAPIPath, id)
	req, err := s.client.newGETRequest(path)
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, err
	}

	var vehicle Vehicle
	_, err = s.client.doSinglePayload(req, &vehicle)
	return &vehicle, err
}
