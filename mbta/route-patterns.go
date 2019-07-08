package mbta

import (
	"context"
	"fmt"
	"net/http"
)

const routesPatternsAPIPath = "/route-patterns"

// RoutePatternsService handling all of the route-patterns related API calls
type RoutePatternsService service

// RoutePatternTypicalityType Explains how common the route pattern is. For the MBTA, this is within the context of the entire route
type RoutePatternTypicalityType int

const (
	// RoutePatternTypicalityNotDefined Not defined
	RoutePatternTypicalityNotDefined RoutePatternTypicalityType = iota
	// RoutePatternTypicalityTypical Typical. Pattern is common for the route. Most routes will have only one such pattern per direction. A few routes may have more than 1, such as the Red Line (with one branch to Ashmont and another to Braintree); routes with more than 2 are rare
	RoutePatternTypicalityTypical
	// RoutePatternTypicalityDeviation Pattern is a deviation from the regular route
	RoutePatternTypicalityDeviation
	// RoutePatternTypicalityHighlyAtypical Pattern represents a highly atypical pattern for the route, such as a special routing which only runs a handful of times per day
	RoutePatternTypicalityHighlyAtypical
	// RoutePatternTypicalityDiversionFromNormal Diversions from normal service, such as planned detours, bus shuttles, or snow routes
	RoutePatternTypicalityDiversionFromNormal
)

// RoutePattern holds all the info about a given MBTA route-pattern
type RoutePattern struct {
	ID                 string                     `jsonapi:"primary,route_pattern"`
	Typicality         RoutePatternTypicalityType `jsonapi:"attr,typicality"`              // Explains how common the route pattern is. For the MBTA, this is within the context of the entire route
	TimeDesc           *string                    `jsonapi:"attr,time_desc"`               // User-facing description of when the route pattern operate. Not all route patterns will include a time description
	SortOrder          int                        `jsonapi:"attr,sort_order"`              // Can be used to order the route patterns in a way which is ideal for presentation to customers. Route patterns with smaller sort_order values should be displayed before those with larger values
	Name               string                     `jsonapi:"attr,name"`                    // User-facing description of where trips on the route pattern serve. These names are published in the form `Destination, Destination via Street or Landmark, Origin - Destination, or Origin - Destination via Street or Landmark`. Note that names for bus and subway route patterns currently do not include the origin location, but will in the future
	DirectionID        int                        `jsonapi:"attr,direction_id"`            // Direction in which trip is traveling: 0 or 1
	RepresentativeTrip *Trip                      `jsonapi:"relation,representative_trip"` // A trip that can be considered a canonical trip for the route pattern. This trip can be used to deduce a pattern’s canonical set of stops and shape. Only includes id by default, use Include config option to get all data
	Route              *Route                     `jsonapi:"relation,route"`               // The route that this pattern belongs to. Only includes id by default, use Include config option to get all data
}

// RoutePatternInclude all of the includes for a route-pattern request
type RoutePatternInclude string

const (
	RoutePatternIncludeRoute              RoutePatternInclude = includeLine
	RoutePatternIncludeRepresentativeTrip RoutePatternInclude = includeRepresentativeTrip
)

// RoutePatternsSortByType all of the possible ways to sort by for a GetAllRoutePatterns request
type RoutePatternsSortByType string

const (
	RoutePatternsSortByDirectionIDAscending  RoutePatternsSortByType = "direction_id"
	RoutePatternsSortByDirectionIDDescending RoutePatternsSortByType = "-direction_id"
	RoutePatternsSortByNameAscending         RoutePatternsSortByType = "name"
	RoutePatternsSortByNameDescending        RoutePatternsSortByType = "-name"
	RoutePatternsSortBySortOrderAscending    RoutePatternsSortByType = "sort_order"
	RoutePatternsSortBySortOrderDescending   RoutePatternsSortByType = "-sort_order"
	RoutePatternsSortByTimeDescAscending     RoutePatternsSortByType = "time_desc"
	RoutePatternsSortByTimeDescDescending    RoutePatternsSortByType = "-time_desc"
	RoutePatternsSortByTypicalityAscending   RoutePatternsSortByType = "typicality"
	RoutePatternsSortByTypicalityDescending  RoutePatternsSortByType = "-typicality"
)

// GetAllRoutePatternsRequestConfig extra options for the GetAllRoutePatterns request
type GetAllRoutePatternsRequestConfig struct {
	PageOffset        string                  `url:"page[offset],omitempty"`         // Offset (0-based) of first element in the page
	PageLimit         string                  `url:"page[limit],omitempty"`          // Max number of elements to return
	Sort              RoutePatternsSortByType `url:"sort,omitempty"`                 // Results can be sorted by the id or any RoutesSortByType
	Include           []RoutePatternInclude   `url:"include,comma,omitempty"`        // Include extra data in response
	FilterDirectionID string                  `url:"filter[direction_id],omitempty"` // Filter by Direction ID (Either "0" or "1")
	FilterIDs         []string                `url:"filter[id],comma,omitempty"`     // Filter by multiple IDs
	FilterRouteIDs    []string                `url:"filter[route],comma,omitempty"`  // Filter by stops
}

// GetAllRoutePatterns returns all routes from the mbta API
func (s *RoutePatternsService) GetAllRoutePatterns(config *GetAllRoutePatternsRequestConfig) ([]*RoutePattern, *http.Response, error) {
	return s.GetAllRoutePatternsWithContext(context.Background(), config)
}

// GetAllRoutePatternsWithContext returns all routes from the mbta API given a context
func (s *RoutePatternsService) GetAllRoutePatternsWithContext(ctx context.Context, config *GetAllRoutePatternsRequestConfig) ([]*RoutePattern, *http.Response, error) {
	u, err := addOptions(routesPatternsAPIPath, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	untypedRoutePatterns, resp, err := s.client.doManyPayload(req, &RoutePattern{})
	routePatterns := make([]*RoutePattern, len(untypedRoutePatterns))
	for i := 0; i < len(untypedRoutePatterns); i++ {
		routePatterns[i] = untypedRoutePatterns[i].(*RoutePattern)
	}
	return routePatterns, resp, err
}

// GetRoutePatternRequestConfig extra options for GetRoutePattern request
type GetRoutePatternRequestConfig struct {
	Include []RouteInclude `url:"include,comma,omitempty"` // Include extra data in response
}

// GetRoutePattern return a route from the mbta API
func (s *RoutePatternsService) GetRoutePattern(id string, config *GetRoutePatternRequestConfig) (*RoutePattern, *http.Response, error) {
	return s.GetRoutePatternWithContext(context.Background(), id, config)
}

// GetRoutePatternWithContext return a route from the mbta API given a context
func (s *RoutePatternsService) GetRoutePatternWithContext(ctx context.Context, id string, config *GetRoutePatternRequestConfig) (*RoutePattern, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", routesPatternsAPIPath, id)
	u, err := addOptions(path, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	var routePattern RoutePattern
	resp, err := s.client.doSinglePayload(req, &routePattern)
	return &routePattern, resp, err
}
