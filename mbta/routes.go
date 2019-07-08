package mbta

import (
	"context"
	"fmt"
	"net/http"
)

const routesAPIPath = "/routes"

// RouteService handling all of the route related API calls
type RouteService service

// RouteType enum for possible Route types (see https://github.com/google/transit/blob/master/gtfs/spec/en/reference.md#routestxt)
type RouteType int

const (
	// RouteTypeLightRail ...
	RouteTypeLightRail RouteType = iota
	// RouteTypeHeavyRail ...
	RouteTypeHeavyRail
	// RouteTypeCommuterRail ...
	RouteTypeCommuterRail
	// RouteTypeBus ...
	RouteTypeBus
	// RouteTypeFerry ...
	RouteTypeFerry
)

// Route holds all the info about a given MBTA Route
type Route struct {
	ID                    string    `jsonapi:"primary,route"`
	Color                 string    `jsonapi:"attr,color"`
	Description           string    `jsonapi:"attr,description"`
	DirectionDestinations []string  `jsonapi:"attr,direction_destinations"`
	DirectionNames        []string  `jsonapi:"attr,direction_names"`
	LongName              string    `jsonapi:"attr,long_name"`
	SortOrder             int       `jsonapi:"attr,sort_order"`
	TextColor             string    `jsonapi:"attr,text_color"`
	Type                  RouteType `jsonapi:"attr,type"`
	ShortName             string    `jsonapi:"attr,short_name"`
	// Line				  Line	    `jsonapi:"relation,line"`
}

type RouteInclude string

const (
	RouteIncludeLine          RouteInclude = includeLine
	RouteIncludeStop          RouteInclude = includeStop
	RouteIncludeRoutePatterns RouteInclude = includeRoutePatterns
)

// GetAllRoutesSortByType all of the possible ways to sort by for a GetAllRoutes request
type GetAllRoutesSortByType string

const (
	GetAllRoutesSortByColorAscending                 GetAllRoutesSortByType = "color"
	GetAllRoutesSortByColorDescending                GetAllRoutesSortByType = "-color"
	GetAllRoutesSortByDescriptionAscending           GetAllRoutesSortByType = "description"
	GetAllRoutesSortByDescriptionDescending          GetAllRoutesSortByType = "-description"
	GetAllRoutesSortByDirectionDestinationAscending  GetAllRoutesSortByType = "direction_destinations"
	GetAllRoutesSortByDirectionDestinationDescending GetAllRoutesSortByType = "-direction_destinations"
	GetAllRoutesSortByDirectionNameAscending         GetAllRoutesSortByType = "direction_names"
	GetAllRoutesSortByDirectionNameDescending        GetAllRoutesSortByType = "-direction_names"
	GetAllRoutesSortByFareClassAscending             GetAllRoutesSortByType = "fare_class"
	GetAllRoutesSortByFareClassDescending            GetAllRoutesSortByType = "-fare_class"
	GetAllRoutesSortByLongNameDesending              GetAllRoutesSortByType = "long_name"
	GetAllRoutesSortByLongNameDescending             GetAllRoutesSortByType = "-long_name"
	GetAllRoutesSortByShortNameAscending             GetAllRoutesSortByType = "short_name"
	GetAllRoutesSortByShortNameDescending            GetAllRoutesSortByType = "-short_name"
	GetAllRoutesSortBySortOrderAscending             GetAllRoutesSortByType = "sort_order"
	GetAllRoutesSortBySortOrderDescending            GetAllRoutesSortByType = "-sort_order"
	GetAllRoutesSortByTextColorAscending             GetAllRoutesSortByType = "text_color"
	GetAllRoutesSortByTextColorDescending            GetAllRoutesSortByType = "-text_color"
	GetAllRoutesSortByTypeAscending                  GetAllRoutesSortByType = "type"
	GetAllRoutesSortByTypeDescending                 GetAllRoutesSortByType = "-type"
)

// GetAllRoutesRequestConfig extra options for the GetAllRoutes request
type GetAllRoutesRequestConfig struct {
	PageOffset        string                 `url:"page[offset],omitempty"`         // Offset (0-based) of first element in the page// Offset (0-based) of first element in the page
	PageLimit         string                 `url:"page[limit],omitempty"`          // Max number of elements to return// Max number of elements to return
	Sort              GetAllRoutesSortByType `url:"sort,omitempty"`                 // Results can be sorted by the id or any GetAllRoutesSortByType
	Include           []RouteInclude         `url:"include,comma,omitempty"`        // Include extra data in response
	Fields            []string               `url:"fields[route],comma,omitempty"`  // Fields to include with the response. Note that fields can also be selected for included data types// Fields to include with the response. Multiple fields MUST be a comma-separated (U+002C COMMA, “,”) list. Note that fields can also be selected for included data types
	FilterDirectionID string                 `url:"filter[direction_id],omitempty"` // Filter by Direction ID (Either "0" or "1")
	FilterDate        string                 `url:"filter[data],omitempty"`         // Filter by date that route is active
	FilterIDs         []string               `url:"filter[id],comma,omitempty"`     // Filter by multiple IDs
	FilterStop        string                 `url:"filter[stop],omitempty"`         // Filter by stops
	FilterRouteTypes  []RouteType            `url:"filter[type],comma,omitempty"`   // Filter by different route types
}

// GetRouteRequestConfig extra options for GetRoute request
type GetRouteRequestConfig struct {
	Fields  []string       `url:"fields[stop],comma,omitempty"` // Fields to include with the response. Note that fields can also be selected for included data types// Fields to include with the response. Multiple fields MUST be a comma-separated (U+002C COMMA, “,”) list. Note that fields can also be selected for included data types
	Include []RouteInclude `url:"include,comma,omitempty"`      // Include extra data in response
}

// GetAllRoutes returns all routes from the mbta API
func (s *RouteService) GetAllRoutes(config *GetAllRoutesRequestConfig) ([]*Route, *http.Response, error) {
	return s.GetAllRoutesWithContext(context.Background(), config)
}

// GetAllRoutesWithContext returns all routes from the mbta API given a context
func (s *RouteService) GetAllRoutesWithContext(ctx context.Context, config *GetAllRoutesRequestConfig) ([]*Route, *http.Response, error) {
	u, err := addOptions(routesAPIPath, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	untypedRoutes, resp, err := s.client.doManyPayload(req, &Route{})
	routes := make([]*Route, len(untypedRoutes))
	for i := 0; i < len(untypedRoutes); i++ {
		routes[i] = untypedRoutes[i].(*Route)
	}
	return routes, resp, err
}

// GetRoute return a route from the mbta API
func (s *RouteService) GetRoute(id string, config *GetRouteRequestConfig) (*Route, *http.Response, error) {
	return s.GetRouteWithContext(context.Background(), id, config)
}

// GetRouteWithContext return a route from the mbta API given a context
func (s *RouteService) GetRouteWithContext(ctx context.Context, id string, config *GetRouteRequestConfig) (*Route, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", routesAPIPath, id)
	u, err := addOptions(path, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	var route Route
	resp, err := s.client.doSinglePayload(req, &route)
	return &route, resp, err
}
