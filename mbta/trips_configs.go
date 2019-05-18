package mbta

import (
	"net/http"
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
	PageOffset            string                // Offset (0-based) of first element in the page
	PageLimit             string                // Max number of elements to return
	Sort                  GetAllTripsSortByType // Results can be sorted by the id or any GetAllTripsSortByType
	Fields                []string              // Fields to include with the response. Multiple fields MUST be a comma-separated (U+002C COMMA, “,”) list. Note that fields can also be selected for included data types
	IncludeRoute          bool                  // Include Route data in response (The primary route for the trip)
	IncludeVehicle        bool                  // Include Vehicle data in response (The vehicle on this trip)
	IncludeService        bool                  // Include Service data in response (The service controlling when this trip is active)
	IncludeShape          bool                  // Include Shape data in response (The shape of the trip)
	IncludeRoutePattern   bool                  // Include RoutePattern data in response (The route pattern for the trip)
	IncludePredictions    bool                  // Include Predictions data in response (Predictions of when the vehicle on this trip will arrive at or depart from each stop on the route(s) on the trip)
	FilterDate            *TimeISO8601          // Filter by trips on a particular date The active date is the service date. Trips that begin between midnight and 3am are considered part of the previous service day
	FilterDirectionID     string                // Filter by direction of travel along the route
	FilterRouteIDs        []string              // Filter by route id(s)
	FilterRoutePatternIDs []string              // Filter by route pattern id(s)
	FilterIDs             []string              // Filter by id(s)
	FilterNames           []string              // Filter by names
}

func (config *GetAllTripsRequestConfig) addHTTPParamsToRequest(req *http.Request) {
	// Add fields and includes params to request
	includes := GetTripRequestConfig{
		Fields:              config.Fields,
		IncludeRoute:        config.IncludeRoute,
		IncludeVehicle:      config.IncludeVehicle,
		IncludeService:      config.IncludeService,
		IncludeShape:        config.IncludeShape,
		IncludeRoutePattern: config.IncludeRoutePattern,
		IncludePredictions:  config.IncludePredictions,
	}
	includes.addHTTPParamsToRequest(req)

	q := req.URL.Query()
	addToQuery(q, "page[offset]", config.PageOffset)
	addToQuery(q, "page[limit]", config.PageLimit)
	addToQuery(q, "sort", string(config.Sort))
	if config.FilterDate != nil {
		addToQuery(q, "filter[date]", config.FilterDate.FormatOnlyDate())
	}
	addToQuery(q, "filter[direction_id]", config.FilterDirectionID)
	addCommaSeparatedListToQuery(q, "filter[route]", config.FilterRouteIDs)
	addCommaSeparatedListToQuery(q, "filter[route_pattern]", config.FilterRoutePatternIDs)
	addCommaSeparatedListToQuery(q, "filter[id]", config.FilterIDs)
	addCommaSeparatedListToQuery(q, "filter[name]", config.FilterNames)
	req.URL.RawQuery = q.Encode()
}

// GetTripRequestConfig extra options for the GetTrip request
type GetTripRequestConfig struct {
	Fields              []string // Fields to include with the response. Multiple fields MUST be a comma-separated (U+002C COMMA, “,”) list. Note that fields can also be selected for included data types
	IncludeRoute        bool     // Include Route data in response (The primary route for the trip)
	IncludeVehicle      bool     // Include Vehicle data in response (The vehicle on this trip)
	IncludeService      bool     // Include Service data in response (The service controlling when this trip is active)
	IncludeShape        bool     // Include Shape data in response (The shape of the trip)
	IncludeRoutePattern bool     // Include RoutePattern data in response (The route pattern for the trip)
	IncludePredictions  bool     // Include Predictions data in response (Predictions of when the vehicle on this trip will arrive at or depart from each stop on the route(s) on the trip)
}

func (config *GetTripRequestConfig) addHTTPParamsToRequest(req *http.Request) {
	q := req.URL.Query()

	includes := []string{}
	if config.IncludeRoute {
		includes = append(includes, "route")
	}
	if config.IncludeVehicle {
		includes = append(includes, "vehicle")
	}
	if config.IncludeService {
		includes = append(includes, "service")
	}
	if config.IncludeShape {
		includes = append(includes, "shape")
	}
	if config.IncludePredictions {
		includes = append(includes, "predictions")
	}
	if config.IncludeRoutePattern {
		includes = append(includes, "route_pattern")
	}

	addCommaSeparatedListToQuery(q, "include", includes)
	addCommaSeparatedListToQuery(q, "fields[trip]", config.Fields)
	req.URL.RawQuery = q.Encode()
}
