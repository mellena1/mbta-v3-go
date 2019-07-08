package mbta

import (
	"context"
	"fmt"
	"net/http"
)

const tripsAPIPath = "/trips"

// TripService service handling all of the trip related API calls
type TripService service

// Trip holds all info about a given MBTA Trip
type Trip struct {
	ID                   string                 `jsonapi:"primary,trip"`
	WheelchairAccessible WheelchairBoardingType `jsonapi:"attr,wheelchair_accessible"` // Indicator of wheelchair accessibility
	Name                 string                 `jsonapi:"attr,name"`                  // The text that appears in schedules and sign boards to identify the trip to passengers
	Headsign             string                 `jsonapi:"attr,headsign"`              // The text that appears on a sign that identifies the trip’s destination to passengers
	DirectionID          int                    `jsonapi:"attr,direction_id"`          // Direction in which trip is traveling: 0 or 1.
	BlockID              string                 `jsonapi:"attr,block_id"`              // ID used to group sequential trips with the same vehicle for a given service_id
	BikesAllowed         BikesAllowedType       `jsonapi:"attr,bikes_allowed"`         // Indicator of whether or not bikes are allowed on this trip
	Route                *Route                 `jsonapi:"relation,route"`             // Trip that the current trip is linked with. Only includes id by default, use Include config option to get all data
	RoutePattern         *RoutePattern          `jsonapi:"relation,route_pattern"`
	Service              *Service               `jsonapi:"relation,service"`
	Shape                *Shape                 `jsonapi:"relation,shape"`
}

// TripInclude all of the includes for a trip request
type TripInclude string

const (
	TripIncludeRoute        TripInclude = includeRoute
	TripIncludeVehicle      TripInclude = includeVehicle
	TripIncludeService      TripInclude = includeService
	TripIncludeShape        TripInclude = includeShape
	TripIncludeRoutePattern TripInclude = includeRoutePattern
	TripIncludePredictions  TripInclude = includePredictions
)

// GetAllTripsSortByType all of the possible ways to sort by for a GetAllTrips request
type GetAllTripsSortByType string

const (
	GetAllTripsSortByBikesAllowedAscending          GetAllTripsSortByType = "bikes_allowed"
	GetAllTripsSortByBikesAllowedDescending         GetAllTripsSortByType = "-bikes_allowed"
	GetAllTripsSortByBlockIDAscending               GetAllTripsSortByType = "block_id"
	GetAllTripsSortByBlockIDDescending              GetAllTripsSortByType = "-block_id"
	GetAllTripsSortByDirectionIDAscending           GetAllTripsSortByType = "direction_id"
	GetAllTripsSortByDirectionIDDescending          GetAllTripsSortByType = "-direction_id"
	GetAllTripsSortByHeadsignAscending              GetAllTripsSortByType = "headsign"
	GetAllTripsSortByHeadsignDescending             GetAllTripsSortByType = "-headsign"
	GetAllTripsSortByNameAscending                  GetAllTripsSortByType = "name"
	GetAllTripsSortByNameDescending                 GetAllTripsSortByType = "-name"
	GetAllTripsSortByWheelchairAccessibleAscending  GetAllTripsSortByType = "wheelchair_accessible"
	GetAllTripsSortByWheelchairAccessibleDescending GetAllTripsSortByType = "-wheelchair_accessible"
)

// GetAllTripsRequestConfig extra options for the GetAllTrips request
type GetAllTripsRequestConfig struct {
	PageOffset            string                `url:"page[offset],omitempty"`                // Offset (0-based) of first element in the page
	PageLimit             string                `url:"page[limit],omitempty"`                 // Max number of elements to return
	Sort                  GetAllTripsSortByType `url:"sort,omitempty"`                        // Results can be sorted by the id or any GetAllTripsSortByType
	Fields                []string              `url:"fields[trip],comma,omitempty"`          // Fields to include with the response. Note that fields can also be selected for included data types
	Include               []TripInclude         `url:"include,comma,omitempty"`               // Include extra data in response (route, vehicle, service, shape, route_pattern, predictions)
	FilterDate            *TimeISO8601          `url:"filter[date],omitempty"`                // Filter by trips on a particular date The active date is the service date. Trips that begin between midnight and 3am are considered part of the previous service day
	FilterDirectionID     string                `url:"filter[direction_id],omitempty"`        // Filter by direction of travel along the route
	FilterRouteIDs        []string              `url:"filter[route],comma,omitempty"`         // Filter by route id(s)
	FilterRoutePatternIDs []string              `url:"filter[route_pattern],comma,omitempty"` // Filter by route pattern id(s)
	FilterIDs             []string              `url:"filter[id],comma,omitempty"`            // Filter by id(s)
	FilterNames           []string              `url:"filter[name],comma,omitempty"`          // Filter by names
}

// GetAllTrips returns all vehicles from the mbta API
func (s *TripService) GetAllTrips(config GetAllTripsRequestConfig) ([]*Trip, *http.Response, error) {
	return s.GetAllTripsWithContext(context.Background(), config)
}

// GetAllTripsWithContext returns all vehicles from the mbta API given a context
func (s *TripService) GetAllTripsWithContext(ctx context.Context, config GetAllTripsRequestConfig) ([]*Trip, *http.Response, error) {
	u, err := addOptions(tripsAPIPath, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	untypedTrips, resp, err := s.client.doManyPayload(req, &Trip{})
	trips := make([]*Trip, len(untypedTrips))
	for i := 0; i < len(untypedTrips); i++ {
		trips[i] = untypedTrips[i].(*Trip)
	}
	return trips, resp, err
}

// GetTripRequestConfig extra options for the GetTrip request
type GetTripRequestConfig struct {
	Fields  []string      `url:"fields[trip],comma,omitempty"` // Fields to include with the response. Note that fields can also be selected for included data types
	Include []TripInclude `url:"include,comma,omitempty"`      // Include extra data in response (route, vehicle, service, shape, route_pattern, predictions)
}

// GetTrip returns a vehicle from the mbta API
func (s *TripService) GetTrip(id string, config GetTripRequestConfig) (*Trip, *http.Response, error) {
	return s.GetTripWithContext(context.Background(), id, config)
}

// GetTripWithContext returns a vehicle from the mbta API given a context
func (s *TripService) GetTripWithContext(ctx context.Context, id string, config GetTripRequestConfig) (*Trip, *http.Response, error) {
	if id == "" {
		return nil, nil, ErrMustSpecifyID
	}

	path := fmt.Sprintf("%s/%s", tripsAPIPath, id)
	u, err := addOptions(path, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	var trip Trip
	resp, err := s.client.doSinglePayload(req, &trip)
	return &trip, resp, err
}
