package mbta

import (
	"context"
	"fmt"
	"net/http"
)

const vehiclesAPIPath = "/vehicles"

// VehicleService service handling all of the vehicle related API calls
type VehicleService service

// VehicleStatus enum for the possible vehicle statuses
type VehicleStatus string

const (
	// InTransitTo vehicleStatus for a vehicle in transit
	InTransitTo VehicleStatus = "IN_TRANSIT_TO"
	// StoppedAt vehicleStatus for a vehicle stopped at a station
	StoppedAt VehicleStatus = "STOPPED_AT"
	// IncomingAt vehicleStatus for a vehicle getting into a station
	IncomingAt VehicleStatus = "INCOMING_AT"
)

// Vehicle holds all info about a given MBTA vehicle
type Vehicle struct {
	ID                  string        `jsonapi:"primary,vehicle"`
	Bearing             float32       `jsonapi:"attr,bearing"`               // Bearing, in degrees, clockwise from True North, i.e., 0 is North and 90 is East
	CurrentStatus       VehicleStatus `jsonapi:"attr,current_status"`        // Status of vehicle relative to the stops
	CurrentStopSequence int           `jsonapi:"attr,current_stop_sequence"` // not sure on this one yet
	DirectionID         int           `jsonapi:"attr,direction_id"`          // Direction in which trip is traveling: 0 or 1.
	Label               string        `jsonapi:"attr,label"`                 // User visible label, such as the one of on the signage on the vehicle
	Latitute            float64       `jsonapi:"attr,latitude"`              // Degrees North, in the WGS-84 coordinate system
	Longitude           float64       `jsonapi:"attr,longitude"`             // Degrees East, in the WGS-84 coordinate system
	Speed               *float32      `jsonapi:"attr,speed"`                 // meters per second
	UpdatedAt           TimeISO8601   `jsonapi:"attr,updated_at"`            // Time at which vehicle information was last updated. Format is ISO8601
	Stop                *Stop         `jsonapi:"relation,stop"`              // Stop that the vehicle is at. Only includes id by default, use IncludeStop config option to get all data
	Trip                *Trip         `jsonapi:"relation,trip"`              // Trip that the current vehicle is on
	// TODO: Route *Route `jsonapi:"relation,route"`
}

// VehicleInclude all of the includes for a vehicle request
type VehicleInclude string

const (
	VehicleIncludeTrip  VehicleInclude = includeTrip
	VehicleIncludeStop  VehicleInclude = includeStop
	VehicleIncludeRoute VehicleInclude = includeRoute
)

// GetAllVehiclesSortByType all of the possible ways to sort by for a GetAllVehicles request
type GetAllVehiclesSortByType string

const (
	GetAllVehiclesSortByBearingAscending              GetAllVehiclesSortByType = "bearing"
	GetAllVehiclesSortByBearingDescending             GetAllVehiclesSortByType = "-bearing"
	GetAllVehiclesSortByCurrentStatusAscending        GetAllVehiclesSortByType = "current_status"
	GetAllVehiclesSortByCurrentStatusDescending       GetAllVehiclesSortByType = "-current_status"
	GetAllVehiclesSortByCurrentStopSequenceAscending  GetAllVehiclesSortByType = "current_stop_sequence"
	GetAllVehiclesSortByCurrentStopSequenceDescending GetAllVehiclesSortByType = "-current_stop_sequence"
	GetAllVehiclesSortByDirectionIDAscending          GetAllVehiclesSortByType = "direction_id"
	GetAllVehiclesSortByDirectionIDDescending         GetAllVehiclesSortByType = "-direction_id"
	GetAllVehiclesSortByLabelAscending                GetAllVehiclesSortByType = "label"
	GetAllVehiclesSortByLabelDescending               GetAllVehiclesSortByType = "-label"
	GetAllVehiclesSortByLatitudeAscending             GetAllVehiclesSortByType = "latitude"
	GetAllVehiclesSortByLatitudeDescending            GetAllVehiclesSortByType = "-latitude"
	GetAllVehiclesSortByLongitudeAscending            GetAllVehiclesSortByType = "longitude"
	GetAllVehiclesSortByLongitudeDescending           GetAllVehiclesSortByType = "-longitude"
	GetAllVehiclesSortBySpeedAscending                GetAllVehiclesSortByType = "speed"
	GetAllVehiclesSortBySpeedDescending               GetAllVehiclesSortByType = "-speed"
	GetAllVehiclesSortByUpdatedAtAscending            GetAllVehiclesSortByType = "updated_at"
	GetAllVehiclesSortByUpdatedAtDescending           GetAllVehiclesSortByType = "-updated_at"
)

// GetAllVehiclesRequestConfig extra options for the GetAllVehicles request
type GetAllVehiclesRequestConfig struct {
	PageOffset        string                   `url:"page[offset],omitempty"`             // Offset (0-based) of first element in the page
	PageLimit         string                   `url:"page[limit],omitempty"`              // Max number of elements to return
	Sort              GetAllVehiclesSortByType `url:"sort,omitempty"`                     // Results can be sorted by the id or any GetAllVehiclesSortByType
	Fields            []string                 `url:"fields[vehicle],comma,omitempty"`    // Fields to include with the response. Multiple fields MUST be a comma-separated (U+002C COMMA, “,”) list. Note that fields can also be selected for included data types
	Include           []VehicleInclude         `url:"include,comma,omitempty"`            // Include extra data in response (trip, stop, or route)
	FilterIDs         []string                 `url:"filter[id],comma,omitempty"`         // Filter by multiple IDs
	FilterTripIDs     []string                 `url:"filter[trip],comma,omitempty"`       // Filter by trip IDs
	FilterLabels      []string                 `url:"filter[label],comma,omitempty"`      // Filter by label
	FilterRouteIDs    []string                 `url:"filter[route],comma,omitempty"`      // Filter by route IDs. If the vehicle is on a multi-route trip, it will be returned for any of the routes
	FilterDirectionID string                   `url:"filter[direction_id],omitempty"`     // Filter by Direction ID (Either "0" or "1")
	FilterRouteTypes  []string                 `url:"filter[route_type],comma,omitempty"` // Filter by route type(s)
}

// GetAllVehicles returns all vehicles from the mbta API
func (s *VehicleService) GetAllVehicles(config *GetAllVehiclesRequestConfig) ([]*Vehicle, *http.Response, error) {
	return s.GetAllVehiclesWithContext(context.Background(), config)
}

// GetAllVehiclesWithContext returns all vehicles from the mbta API given a context
func (s *VehicleService) GetAllVehiclesWithContext(ctx context.Context, config *GetAllVehiclesRequestConfig) ([]*Vehicle, *http.Response, error) {
	u, err := addOptions(vehiclesAPIPath, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	untypedVehicles, resp, err := s.client.doManyPayload(req, &Vehicle{})
	vehicles := make([]*Vehicle, len(untypedVehicles))
	for i := 0; i < len(untypedVehicles); i++ {
		vehicles[i] = untypedVehicles[i].(*Vehicle)
	}
	return vehicles, resp, err
}

// GetVehicleRequestConfig extra options for the GetVehicle request
type GetVehicleRequestConfig struct {
	Fields  []string         `url:"fields[vehicle],comma,omitempty"` // Fields to include with the response. Multiple fields MUST be a comma-separated (U+002C COMMA, “,”) list. Note that fields can also be selected for included data types
	Include []VehicleInclude `url:"include,comma,omitempty"`         // Include extra data in response (trip, stop, or route)
}

// GetVehicle returns a vehicle from the mbta API
func (s *VehicleService) GetVehicle(id string, config *GetVehicleRequestConfig) (*Vehicle, *http.Response, error) {
	return s.GetVehicleWithContext(context.Background(), id, config)
}

// GetVehicleWithContext returns a vehicle from the mbta API given a context
func (s *VehicleService) GetVehicleWithContext(ctx context.Context, id string, config *GetVehicleRequestConfig) (*Vehicle, *http.Response, error) {
	if id == "" {
		return nil, nil, ErrMustSpecifyID
	}

	path := fmt.Sprintf("%s/%s", vehiclesAPIPath, id)
	u, err := addOptions(path, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	var vehicle Vehicle
	resp, err := s.client.doSinglePayload(req, &vehicle)
	return &vehicle, resp, err
}
