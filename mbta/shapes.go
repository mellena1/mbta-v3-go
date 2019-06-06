package mbta

import (
	"context"
	"fmt"
	"net/http"
)

const shapesAPIPath = "/shapes"

// ShapeService handles all of the shape related API calls
type ShapeService service

// Shape holds all the info about an MBTA shape
type Shape struct {
	ID          string  `jsonapi:"primary,shape"`
	Priority    int     `jsonapi:"attr,priority"`
	Polyline    string  `jsonapi:"attr,polyline"`
	Name        string  `jsonapi:"attr,name"`
	DirectionID int     `jsonapi:"attr,direction_id"`
	Stops       []*Stop `jsonapi:"relation,stops"`
	Route       *Route  `jsonapi:"relation,route"`
}

// GetAllShapesSortByType is an enumerable for all the ways you can sort shapes
type GetAllShapesSortByType string

const (
	GetAllShapesSortByDirectionIDAscending GetAllShapesSortByType = "direction_id"
	GetAllShapesSortByDirectionIDDecending GetAllShapesSortByType = "-direction_id"
	GetAllShapesSortByNameAscending        GetAllShapesSortByType = "name"
	GetAllShapesSortByNameDecending        GetAllShapesSortByType = "-name"
	GetAllShapesSortByPolylineAscending    GetAllShapesSortByType = "polyline"
	GetAllShapesSortByPolylineDecending    GetAllShapesSortByType = "-polyline"
	GetAllShapesSortByPriorityAscending    GetAllShapesSortByType = "priority"
	GetAllShapesSortByPriorityDecending    GetAllShapesSortByType = "-priority"
)

// ShapeInclude is an enumerable for all the includes for shape
type ShapeInclude string

const (
	ShapeIncludeRoute ShapeInclude = includeRoute
	ShapeIncludeStop               = includeStop
)

// GetAllShapesRequestConfig holds the request info for the GetAllShapes function
type GetAllShapesRequestConfig struct {
	PageOffset        int                    `url:"page[offset],omitempty"`         // Offset (0-based) of first element in the page
	PageLimit         int                    `url:"page[limit],omitempty"`          // Max number of elements to return
	Sort              GetAllShapesSortByType `url:"sort,omitempty"`                 // Results can be sorted by the id or any /data/{index}/attributes key. Assumes ascending; may be prefixed with '-' for descending
	Fields            []string               `url:"fields[shape],comma,omitempty"`  // Fields to include with the response. Multiple fields MUST be a comma-separated (U+002C COMMA, “,”) list. Note that fields can also be selected for included data types: see the V3 API Best Practices (https://www.mbta.com/developers/v3-api/best-practices) for an example.
	Include           []ShapeInclude         `url:"include,comma,omitempty"`        // Can include choose to include route and stop.
	FilterRoute       []string               `url:"filter[route],comma,omitempty"`  // Filter by /data/{index}/relationships/route/data/id.
	FilterDirectionID string                 `url:"filter[direction_id],omitempty"` // Filter by direction of travel along the route.
}

// GetShapeRequestConfig holds the request info for the GetShape function
type GetShapeRequestConfig struct {
	Fields  []string `url:"fields[shape],comma,omitempty"` // Fields to include with the response. Multiple fields MUST be a comma-separated (U+002C COMMA, “,”) list. Note that fields can also be selected for included data types: see the V3 API Best Practices for an example.
	Include []string `url:"include,comma,omitempty"`       // Can include choose to include route and stop.
}

// GetAllShapes gets all the shapes based on the config info
func (s *ShapeService) GetAllShapes(config GetAllShapesRequestConfig) ([]*Shape, *http.Response, error) {
	return s.GetAllShapesWithContext(context.Background(), config)
}

// GetAllShapesWithContext gets all the shapes based on the config info and accepts a context
func (s *ShapeService) GetAllShapesWithContext(ctx context.Context, config GetAllShapesRequestConfig) ([]*Shape, *http.Response, error) {
	u, err := addOptions(shapesAPIPath, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	untypedShapes, resp, err := s.client.doManyPayload(req, &Shape{})
	shapes := make([]*Shape, len(untypedShapes))
	for i := 0; i < len(untypedShapes); i++ {
		shapes[i] = untypedShapes[i].(*Shape)
	}
	return shapes, resp, err
}

// GetShape gets the shape with the specified ID
func (s *ShapeService) GetShape(id string, config GetShapeRequestConfig) (*Shape, *http.Response, error) {
	return s.GetShapeWithContext(context.Background(), id, config)

}

// GetShapeWithContext gets the shape with the specified ID and accepts context
func (s *ShapeService) GetShapeWithContext(ctx context.Context, id string, config GetShapeRequestConfig) (*Shape, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", shapesAPIPath, id)
	u, err := addOptions(path, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	var shape Shape
	resp, err := s.client.doSinglePayload(req, &shape)
	return &shape, resp, err
}
