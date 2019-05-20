package mbta

import (
	"context"
	"fmt"
)

const routesAPIPath = "/routes"

// RouteService handling all of the route related API calls
type RouteService service

// GetAllRoutes returns all routes from the mbta API
func (s *RouteService) GetAllRoutes(config GetAllRoutesRequestConfig) ([]*Route, error) {
	return s.GetAllRoutesWithContext(context.Background(), config)
}

// GetAllRoutesWithContext returns all routes from the mbta API given a context
func (s *RouteService) GetAllRoutesWithContext(ctx context.Context, config GetAllRoutesRequestConfig) ([]*Route, error) {
	req, err := s.client.newGETRequest(routesAPIPath)
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, err
	}

	untypedRoutes, _, err := s.client.doManyPayload(req, Route{})
	routes := make([]*Route, len(untypedRoutes))
	for i := 0; i < len(untypedRoutes); i++ {
		routes[i] = untypedRoutes[i].(*Route)
	}
	return routes, err
}

// GetRoute return a route from the mbta API
func (s *RouteService) GetRoute(id string, config GetRouteRequestConfig) (*Route, error) {
	return s.GetRouteWithContext(context.Background(), id, config)
}

// GetRouteWithContext return a route from the mbta API given a context
func (s *RouteService) GetRouteWithContext(ctx context.Context, id string, config GetRouteRequestConfig) (*Route, error) {
	path := fmt.Sprintf("%s/%s", routesAPIPath, id)
	req, err := s.client.newGETRequest(path)
	config.addHTTPParamsToRequest(req)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, err
	}

	var route Route
	_, err = s.client.doSinglePayload(req, &route)
	return &route, err

}
