package mbta

import (
	"context"
	"fmt"
	"net/http"
)

const vehiclesAPIPath = "/vehicles"

// VehicleService service handling all of the vehicle related API calls
type VehicleService service

// GetAllVehicles returns all vehicles from the mbta API
func (s *VehicleService) GetAllVehicles(config GetAllVehiclesRequestConfig) ([]*Vehicle, *http.Response, error) {
	return s.GetAllVehiclesContext(context.Background(), config)
}

// GetAllVehiclesContext returns all vehicles from the mbta API given a context
func (s *VehicleService) GetAllVehiclesContext(ctx context.Context, config GetAllVehiclesRequestConfig) ([]*Vehicle, *http.Response, error) {
	req, err := s.client.newGETRequest(vehiclesAPIPath)
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, nil, err
	}

	untypedVehicles, resp, err := s.client.doManyPayload(req, &Vehicle{})
	vehicles := make([]*Vehicle, len(untypedVehicles))
	for i := 0; i < len(untypedVehicles); i++ {
		vehicles[i] = untypedVehicles[i].(*Vehicle)
	}
	return vehicles, resp, err
}

// GetVehicle returns a vehicle from the mbta API
func (s *VehicleService) GetVehicle(id string, config GetVehicleRequestConfig) (*Vehicle, *http.Response, error) {
	return s.GetVehicleContext(context.Background(), id, config)
}

// GetVehicleContext returns a vehicle from the mbta API given a context
func (s *VehicleService) GetVehicleContext(ctx context.Context, id string, config GetVehicleRequestConfig) (*Vehicle, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", vehiclesAPIPath, id)
	req, err := s.client.newGETRequest(path)
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, nil, err
	}

	var vehicle Vehicle
	resp, err := s.client.doSinglePayload(req, &vehicle)
	return &vehicle, resp, err
}
