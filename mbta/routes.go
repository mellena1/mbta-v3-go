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
	Line                  *Line     `jsonapi:"relation,line"`
}

// RouteInclude all of the includes for a route request
type RouteInclude string

const (
	RouteIncludeLine          RouteInclude = includeLine
	RouteIncludeStop          RouteInclude = includeStop
	RouteIncludeRoutePatterns RouteInclude = includeRoutePatterns
)

// RoutesSortByType all of the possible ways to sort by for a GetAllRoutes request
type RoutesSortByType string

const (
	RoutesSortByColorAscending                 RoutesSortByType = "color"
	RoutesSortByColorDescending                RoutesSortByType = "-color"
	RoutesSortByDescriptionAscending           RoutesSortByType = "description"
	RoutesSortByDescriptionDescending          RoutesSortByType = "-description"
	RoutesSortByDirectionDestinationAscending  RoutesSortByType = "direction_destinations"
	RoutesSortByDirectionDestinationDescending RoutesSortByType = "-direction_destinations"
	RoutesSortByDirectionNameAscending         RoutesSortByType = "direction_names"
	RoutesSortByDirectionNameDescending        RoutesSortByType = "-direction_names"
	RoutesSortByFareClassAscending             RoutesSortByType = "fare_class"
	RoutesSortByFareClassDescending            RoutesSortByType = "-fare_class"
	RoutesSortByLongNameDesending              RoutesSortByType = "long_name"
	RoutesSortByLongNameDescending             RoutesSortByType = "-long_name"
	RoutesSortByShortNameAscending             RoutesSortByType = "short_name"
	RoutesSortByShortNameDescending            RoutesSortByType = "-short_name"
	RoutesSortBySortOrderAscending             RoutesSortByType = "sort_order"
	RoutesSortBySortOrderDescending            RoutesSortByType = "-sort_order"
	RoutesSortByTextColorAscending             RoutesSortByType = "text_color"
	RoutesSortByTextColorDescending            RoutesSortByType = "-text_color"
	RoutesSortByTypeAscending                  RoutesSortByType = "type"
	RoutesSortByTypeDescending                 RoutesSortByType = "-type"
)

// GetAllRoutesRequestConfig extra options for the GetAllRoutes request
type GetAllRoutesRequestConfig struct {
	PageOffset        string           `url:"page[offset],omitempty"`         // Offset (0-based) of first element in the page
	PageLimit         string           `url:"page[limit],omitempty"`          // Max number of elements to return
	Sort              RoutesSortByType `url:"sort,omitempty"`                 // Results can be sorted by the id or any RoutesSortByType
	Include           []RouteInclude   `url:"include,comma,omitempty"`        // Include extra data in response
	Fields            []string         `url:"fields[route],comma,omitempty"`  // Fields to include with the response. Note that fields can also be selected for included data types
	FilterDirectionID string           `url:"filter[direction_id],omitempty"` // Filter by Direction ID (Either "0" or "1")
	FilterDate        string           `url:"filter[data],omitempty"`         // Filter by date that route is active
	FilterIDs         []string         `url:"filter[id],comma,omitempty"`     // Filter by multiple IDs
	FilterStop        string           `url:"filter[stop],omitempty"`         // Filter by stops
	FilterRouteTypes  []RouteType      `url:"filter[type],comma,omitempty"`   // Filter by different route types
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

// GetRouteRequestConfig extra options for GetRoute request
type GetRouteRequestConfig struct {
	Fields  []string       `url:"fields[route],comma,omitempty"` // Fields to include with the response. Note that fields can also be selected for included data types
	Include []RouteInclude `url:"include,comma,omitempty"`       // Include extra data in response
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
