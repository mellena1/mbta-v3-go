package mbta

import (
	"context"
	"fmt"
	"net/http"
)

const routesAPIPath = "/routes"

// RouteService handling all of the route related API calls
type RouteService service

// GetAllRoutes returns all routes from the mbta API
func (s *RouteService) GetAllRoutes(config GetAllRoutesRequestConfig) ([]*Route, *http.Response, error) {
	return s.GetAllRoutesWithContext(context.Background(), config)
}

// GetAllRoutesWithContext returns all routes from the mbta API given a context
func (s *RouteService) GetAllRoutesWithContext(ctx context.Context, config GetAllRoutesRequestConfig) ([]*Route, *http.Response, error) {
	req, err := s.client.newGETRequest(routesAPIPath)
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, nil, err
	}

	untypedRoutes, resp, err := s.client.doManyPayload(req, &Route{})
	routes := make([]*Route, len(untypedRoutes))
	for i := 0; i < len(untypedRoutes); i++ {
		routes[i] = untypedRoutes[i].(*Route)
	}
	return routes, resp, err
}

// GetRoute return a route from the mbta API
func (s *RouteService) GetRoute(id string, config GetRouteRequestConfig) (*Route, *http.Response, error) {
	return s.GetRouteWithContext(context.Background(), id, config)
}

// GetRouteWithContext return a route from the mbta API given a context
func (s *RouteService) GetRouteWithContext(ctx context.Context, id string, config GetRouteRequestConfig) (*Route, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", routesAPIPath, id)
	req, err := s.client.newGETRequest(path)
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, nil, err
	}

	var route Route
	resp, err := s.client.doSinglePayload(req, &route)
	return &route, resp, err

}
