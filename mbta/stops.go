package mbta

import (
	"context"
	"fmt"
	"net/http"
)

const stopsAPIPath = "/stops"

// StopService service handling all of the stop related API calls
type StopService service

// StopLocationType enum for the possible stop location types
type StopLocationType int

const (
	// StopLocationStop A location where passengers board or disembark from a transit vehicle
	StopLocationStop StopLocationType = iota
	// StopLocationStation A physical structure or area that contains one or more stops
	StopLocationStation
	// StopLocationStationEntranceExit A location where passengers can enter or exit a station from the street
	StopLocationStationEntranceExit
)

// Stop holds all info about a given MBTA Stop
type Stop struct {
	ID                 string                 `jsonapi:"primary,stop"`
	Address            *string                `jsonapi:"attr,address"`             // A street address for the station
	Description        *string                `jsonapi:"attr,description"`         // Description of the stop
	Latitude           float64                `jsonapi:"attr,latitude"`            // Degrees North, in the WGS-84 coordinate system
	LocationType       StopLocationType       `jsonapi:"attr,location_type"`       // The type of the stop
	Longitude          float64                `jsonapi:"attr,longitude"`           // Degrees East, in the WGS-84 coordinate system
	Name               string                 `jsonapi:"attr,name"`                // Name of a stop or station in the local and tourist vernacular
	PlatformCode       *string                `jsonapi:"attr,platform_code"`       // A short code representing the platform/track (like a number or letter)
	PlatformName       *string                `jsonapi:"attr,platform_name"`       // A textual description of the platform or track
	WheelchairBoarding WheelchairBoardingType `jsonapi:"attr,wheelchair_boarding"` // Whether there are any vehicles with wheelchair boarding or paths to stops that are wheelchair acessible
	ParentStation      *Stop                  `jsonapi:"relation,parent_station"`  // The link to the parent station. Only includes id by default, use IncludeParentStation config option to get all data
}

// StopInclude all of the includes for a stop request
type StopInclude string

const (
	StopIncludeParentStation StopInclude = includeParentStation
)

// GetAllStopsSortByType all of the possible ways to sort by for a GetAllStops request
type GetAllStopsSortByType string

const (
	GetAllStopsSortByAddressAscending             GetAllStopsSortByType = "address"
	GetAllStopsSortByAddressDescending            GetAllStopsSortByType = "-address"
	GetAllStopsSortByDescriptionAscending         GetAllStopsSortByType = "description"
	GetAllStopsSortByDescriptionDescending        GetAllStopsSortByType = "-description"
	GetAllStopsSortByLatitudeAscending            GetAllStopsSortByType = "latitude"
	GetAllStopsSortByLatitudeDescending           GetAllStopsSortByType = "-latitude"
	GetAllStopsSortByLocationTypeAscending        GetAllStopsSortByType = "location_type"
	GetAllStopsSortByLocationTypeDescending       GetAllStopsSortByType = "-location_type"
	GetAllStopsSortByLongitudeAscending           GetAllStopsSortByType = "longitude"
	GetAllStopsSortByLongitudeDescending          GetAllStopsSortByType = "-longitude"
	GetAllStopsSortByNameAscending                GetAllStopsSortByType = "name"
	GetAllStopsSortByNameDescending               GetAllStopsSortByType = "-name"
	GetAllStopsSortByPlatformCodeAscending        GetAllStopsSortByType = "platform_code"
	GetAllStopsSortByPlatformCodeDescending       GetAllStopsSortByType = "-platform_code"
	GetAllStopsSortByPlatformNameAscending        GetAllStopsSortByType = "platform_name"
	GetAllStopsSortByPlatformNameDescending       GetAllStopsSortByType = "-platform_name"
	GetAllStopsSortByWheelchairBoardingAscending  GetAllStopsSortByType = "wheelchair_boarding"
	GetAllStopsSortByWheelchairBoardingDescending GetAllStopsSortByType = "-wheelchair_boarding"
	GetAllStopsSortByDistanceAscending            GetAllStopsSortByType = "distance"
	GetAllStopsSortByDistanceDescending           GetAllStopsSortByType = "-distance"
)

// GetAllStopsRequestConfig extra options for the GetAllStops request
type GetAllStopsRequestConfig struct {
	PageOffset         string                `url:"page[offset],omitempty"`                // Offset (0-based) of first element in the page
	PageLimit          string                `url:"page[limit],omitempty"`                 // Max number of elements to return
	Sort               GetAllStopsSortByType `url:"sort,omitempty"`                        // Results can be sorted by the id or any GetAllStopsSortByType
	Fields             []string              `url:"fields[stop],comma,omitempty"`          // Fields to include with the response. Note that fields can also be selected for included data types
	Include            []StopInclude         `url:"include,comma,omitempty"`               // Include extra data in response (parentstation)
	FilterDirectionID  string                `url:"filter[direction_id],omitempty"`        // Filter by Direction ID (Either "0" or "1")
	FilterLatitude     string                `url:"filter[latitude],omitempty"`            // Latitude in degrees North in the WGS-84 coordinate system to search filter[radius] degrees around with filter[longitude]
	FilterLongitude    string                `url:"filter[longitude],omitempty"`           // Longitude in degrees East in the WGS-84 coordinate system to search filter[radius] degrees around with filter[latitude]
	FilterRadius       string                `url:"filter[radius],omitempty"`              // The distance is in degrees as if latitude and longitude were on a flat 2D plane and normal Pythagorean distance was calculated. Over the region MBTA serves, 0.02 degrees is approximately 1 mile. Defaults to 0.01 degrees (approximately a half mile)
	FilterIDs          []string              `url:"filter[id],comma,omitempty"`            // Filter by multiple IDs
	FilterRouteTypes   []string              `url:"filter[route_type],comma,omitempty"`    // Filter by route type(s)
	FilterRouteIDs     []string              `url:"filter[route],comma,omitempty"`         // Filter by route IDs. If the vehicle is on a multi-route trip, it will be returned for any of the routes
	FilterLocationType []string              `url:"filter[location_type],comma,omitempty"` // Filter by location type
}

// GetAllStops returns all stops from the mbta API
func (s *StopService) GetAllStops(config *GetAllStopsRequestConfig) ([]*Stop, *http.Response, error) {
	return s.GetAllStopsWithContext(context.Background(), config)
}

// GetAllStopsWithContext returns all stops from the mbta API given a context
func (s *StopService) GetAllStopsWithContext(ctx context.Context, config *GetAllStopsRequestConfig) ([]*Stop, *http.Response, error) {
	u, err := addOptions(stopsAPIPath, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	untypedStops, resp, err := s.client.doManyPayload(req, &Stop{})
	stops := make([]*Stop, len(untypedStops))
	for i := 0; i < len(untypedStops); i++ {
		stops[i] = untypedStops[i].(*Stop)
	}
	return stops, resp, err
}

// GetStopRequestConfig extra options for the GetStop request
type GetStopRequestConfig struct {
	Fields  []string      `url:"fields[stop],comma,omitempty"` // Fields to include with the response. Note that fields can also be selected for included data types
	Include []StopInclude `url:"include,comma,omitempty"`      // Include extra data in response (parentstation)
}

// GetStop returns a stop from the mbta API
func (s *StopService) GetStop(id string, config *GetStopRequestConfig) (*Stop, *http.Response, error) {
	return s.GetStopWithContext(context.Background(), id, config)
}

// GetStopWithContext returns a stop from the mbta API given a context
func (s *StopService) GetStopWithContext(ctx context.Context, id string, config *GetStopRequestConfig) (*Stop, *http.Response, error) {
	if id == "" {
		return nil, nil, ErrMustSpecifyID
	}

	path := fmt.Sprintf("%s/%s", stopsAPIPath, id)
	u, err := addOptions(path, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	var stop Stop
	resp, err := s.client.doSinglePayload(req, &stop)
	return &stop, resp, err
}
