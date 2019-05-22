package mbta

import "net/http"

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
	PageOffset           string                 // Offset (0-based) of first element in the page
	PageLimit            string                 // Max number of elements to return
	Sort                 GetAllRoutesSortByType // Results can be sorted by the id or any GetAllRoutesSortType
	IncludeStop          bool                   // Include Stop in response
	IncludeLine          bool                   // Include Line in response
	IncludeRoutePatterns bool                   // Include RoutePatterns in response
	FilterDirectionID    string                 // Filter by Direction ID (Either "0" or "1")
	FilterDate           string                 // Filter by date that route is active
	FilterIDs            []string               // Filter by multiple IDs
	FilterStop           string                 // Filter by stops
	FilterRouteTypes     []RouteType            // Filter by different route types
}

// GetRouteRequestConfig extra options for GetRoute request
type GetRouteRequestConfig struct {
	IncludeLine          bool // Include Line in response
	IncludeRoutePatterns bool // Include RoutePatterns in response
}

func (config *GetAllRoutesRequestConfig) addHTTPParamsToRequest(req *http.Request) {
	// Add include params to request
	includes := GetRouteRequestConfig{
		IncludeLine: config.IncludeLine,
	}
	includes.addHTTPParamsToRequest(req)

	q := req.URL.Query()
	addToQuery(q, "page[offset]", config.PageOffset)
	addToQuery(q, "page[limit]", config.PageLimit)
	addToQuery(q, "sort", string(config.Sort))
	addToQuery(q, "filter[direction_id]", config.FilterDirectionID)
	addToQuery(q, "filter[date]", config.FilterDate)
	addToQuery(q, "filter[stop]", config.FilterStop)
	addCommaSeparatedListToQuery(q, "filter[id]", config.FilterIDs)
	addCommaSeparatedListToQuery(q, "filter[type]", routeTypeSliceToStringSlice(config.FilterRouteTypes))

	req.URL.RawQuery = q.Encode()
}

func (config *GetRouteRequestConfig) addHTTPParamsToRequest(req *http.Request) {
	q := req.URL.Query()
	if config.IncludeLine {
		q.Add("include", "line")
		req.URL.RawQuery = q.Encode()
	}
}

func routeTypeSliceToStringSlice(routeTypes []RouteType) []string {
	filterRouteTypesStringSlice := []string{}
	for _, i := range routeTypes {
		filterRouteTypesStringSlice = append(filterRouteTypesStringSlice, string(i))
	}
	return filterRouteTypesStringSlice
}
