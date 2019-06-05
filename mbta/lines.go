package mbta

import (
	"context"
	"fmt"
	"net/http"
)

const linesAPIPath = "/lines"

// LineService handling all of the route related API calls
type LineService service

// Line holds all the info about a given MBTA Route
type Line struct {
	ID        string  `jsonapi:"primary,line"`
	Color     string  `jsonapi:"attr,color"`
	LongName  string  `jsonapi:"attr,long_name"`
	ShortName string  `jsonapi:"attr,short_name"`
	SortOrder int     `jsonapi:"attr,sort_order"`
	TextColor string  `jsonapi:"attr,text_color"`
	Routes    []Route `jsonapi:"relation,route`
}

type LineInclude string

const (
	LineIncludeRoutes LineInclude = includeRoutes
)

// GetAllLinesSortByType all of the possible ways to sort by for a GetAllLines request
type GetAllLinesSortByType string

const (
	GetAllLinesSortByColorAscending      GetAllLinesSortByType = "color"
	GetAllLinesSortByColorDescending     GetAllLinesSortByType = "-color"
	GetAllLinesSortByLongNameDesending   GetAllLinesSortByType = "long_name"
	GetAllLinesSortByLongNameDescending  GetAllLinesSortByType = "-long_name"
	GetAllLinesSortByShortNameAscending  GetAllLinesSortByType = "short_name"
	GetAllLinesSortByShortNameDescending GetAllLinesSortByType = "-short_name"
	GetAllLinesSortBySortOrderAscending  GetAllLinesSortByType = "sort_order"
	GetAllLinesSortBySortOrderDescending GetAllLinesSortByType = "-sort_order"
	GetAllLinesSortByTextColorAscending  GetAllLinesSortByType = "text_color"
	GetAllLinesSortByTextColorDescending GetAllLinesSortByType = "-text_color"
)

// GetAllLinesRequestConfig extra options for the GetAllLines request
type GetAllLinesRequestConfig struct {
	PageOffset string                `url:"page[offset],omitempty"`       // Offset (0-based) of first element in the page// Offset (0-based) of first element in the page
	PageLimit  string                `url:"page[limit],omitempty"`        // Max number of elements to return// Max number of elements to return
	Sort       GetAllLinesSortByType `url:"sort,omitempty"`               // Results can be sorted by the id or any GetAllRoutesSortByType
	Fields     []string              `url:"fields[line],comma,omitempty"` // Fields to include with the response. Note that fields can also be selected for included data types// Fields to include with the response. Multiple fields MUST be a comma-separated (U+002C COMMA, “,”) list. Note that fields can also be selected for included data types
	Include    []LineInclude         `url:"include,comma,omitempty"`      // Include extra data in response
	FilterIDs  []string              `url:"filter[id],comma,omitempty"`   // Filter by multiple IDs
}

// GetLineRequestConfig extra options for the GetAllLines request
type GetLineRequestConfig struct {
	Fields  []string      `url:"fields[line],comma,omitempty"` // Fields to include with the response. Note that fields can also be selected for included data types// Fields to include with the response. Multiple fields MUST be a comma-separated (U+002C COMMA, “,”) list. Note that fields can also be selected for included data types
	Include []LineInclude `url:"include,comma,omitempty"`      // Include extra data in response
}

// GetAllLines returns all lines from the mbta API
func (s *LineService) GetAllLines(config GetAllLinesRequestConfig) ([]*Line, *http.Response, error) {
	return s.GetAllLinesWithContext(context.Background(), config)
}

// GetAllLinesWithContext returns all lines from the mbta API given a context
func (s *LineService) GetAllLinesWithContext(ctx context.Context, config GetAllLinesRequestConfig) ([]*Line, *http.Response, error) {
	u, err := addOptions(linesAPIPath, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	untypedLines, resp, err := s.client.doManyPayload(req, &Line{})
	lines := make([]*Line, len(untypedLines))
	for i := 0; i < len(untypedLines); i++ {
		lines[i] = untypedLines[i].(*Line)
	}

	return lines, resp, err
}

// GetLine return a line from the mbta API
func (s *LineService) GetLine(id string, config GetLineRequestConfig) (*Line, *http.Response, error) {
	return s.GetLineWithContext(context.Background(), id, config)
}

// GetLineWithContext return a line from the mbta API given a context
func (s *LineService) GetLineWithContext(ctx context.Context, id string, config GetLineRequestConfig) (*Line, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", linesAPIPath, id)
	u, err := addOptions(path, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	var line Line
	resp, err := s.client.doSinglePayload(req, &line)
	return &line, resp, err
}
