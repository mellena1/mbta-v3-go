package mbta

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// VehicleService service handling all of the vehicle related API calls
type VehicleService service

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
	PageOffset        string                   // Offset (0-based) of first element in the page
	PageLimit         string                   // Max number of elements to return
	Sort              GetAllVehiclesSortByType // Results can be sorted by the id or any GetAllVehiclesSortByType
	IncludeTrip       bool                     // Include Trip object in response
	IncludeStop       bool                     // Include Stop object in response
	IncludeRoute      bool                     // Include Route object in response
	FilterIDs         []string                 // Filter by multiple IDs
	FilterTripIDs     []string                 // Filter by trip IDs
	FilterLabels      []string                 // Filter by label
	FilterRouteIDs    []string                 // Filter by route IDs. If the vehicle is on a multi-route trip, it will be returned for any of the routes
	FilterDirectionID string                   // Filter by Direction ID (Either "0" or "1")
	FilterRouteTypes  []string                 // Filter by route type(s)
}

func (config *GetAllVehiclesRequestConfig) addHTTPParamsToRequest(req *http.Request) {
	// Add includes params to request
	includes := GetVehicleRequestConfig{
		IncludeTrip:  config.IncludeTrip,
		IncludeStop:  config.IncludeStop,
		IncludeRoute: config.IncludeRoute,
	}
	includes.addHTTPParamsToRequest(req)

	q := req.URL.Query()
	if config.PageOffset != "" {
		q.Add("page[offset]", config.PageOffset)
	}
	if config.PageLimit != "" {
		q.Add("page[limit]", config.PageLimit)
	}
	if config.Sort != "" {
		q.Add("sort", string(config.Sort))
	}
	if config.FilterDirectionID != "" {
		q.Add("filter[direction_id]", config.FilterDirectionID)
	}

	addCommaSeparatedList := func(key string, l []string) {
		if len(l) > 0 {
			q.Add(key, strings.Join(l, ","))
		}
	}
	addCommaSeparatedList("filter[id]", config.FilterIDs)
	addCommaSeparatedList("filter[trip]", config.FilterTripIDs)
	addCommaSeparatedList("filter[label]", config.FilterLabels)
	addCommaSeparatedList("filter[route]", config.FilterRouteIDs)
	addCommaSeparatedList("filter[route_type]", config.FilterRouteTypes)

	req.URL.RawQuery = q.Encode()
}

// GetAllVehicles returns all vehicles from the mbta API
func (s *VehicleService) GetAllVehicles(config GetAllVehiclesRequestConfig) ([]Vehicle, error) {
	return s.GetAllVehiclesContext(context.Background(), config)
}

// GetAllVehiclesContext returns all vehicles from the mbta API given a context
func (s *VehicleService) GetAllVehiclesContext(ctx context.Context, config GetAllVehiclesRequestConfig) ([]Vehicle, error) {
	req, err := s.client.newRequest("GET", "/vehicles", nil)
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, err
	}

	var data struct {
		Vehicles []Vehicle `json:"data"`
	}
	_, err = s.client.do(req, &data)
	return data.Vehicles, err
}

// GetVehicleRequestConfig extra options for the GetVehicle request
// TODO: Figure out how to give the included objs back to the user
type GetVehicleRequestConfig struct {
	IncludeTrip  bool
	IncludeStop  bool
	IncludeRoute bool
}

func (config *GetVehicleRequestConfig) addHTTPParamsToRequest(req *http.Request) {
	q := req.URL.Query()

	includes := []string{}
	if config.IncludeTrip {
		includes = append(includes, "trip")
	}
	if config.IncludeStop {
		includes = append(includes, "stop")
	}
	if config.IncludeRoute {
		includes = append(includes, "route")
	}
	if len(includes) > 0 {
		includesString := strings.Join(includes, ",")
		q.Add("include", includesString)
		req.URL.RawQuery = q.Encode()
	}
}

// GetVehicle returns a vehicle from the mbta API
func (s *VehicleService) GetVehicle(id string, config GetVehicleRequestConfig) (Vehicle, error) {
	return s.GetVehicleContext(context.Background(), id, config)
}

// GetVehicleContext returns a vehicle from the mbta API given a context
func (s *VehicleService) GetVehicleContext(ctx context.Context, id string, config GetVehicleRequestConfig) (Vehicle, error) {
	path := fmt.Sprintf("/vehicles/%s", id)
	req, err := s.client.newRequest("GET", path, nil)
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)
	if err != nil {
		return Vehicle{}, err
	}

	var data struct {
		Vehicle Vehicle `json:"data"`
	}
	_, err = s.client.do(req, &data)
	return data.Vehicle, err
}
