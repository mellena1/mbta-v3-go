package mbta

import "net/http"

// GetAllStopsSortByType all of the possible ways to sort by for a GetAllStops request
type GetAllStopsSortByType string

const (
	GetAllStopsSortByAddressAscending             GetAllStopsSortByType = "address"
	GetAllStopsSortByAddressDescending            GetAllStopsSortByType = "-address"
	GetAllStopsSortByDescriptionAscending         GetAllStopsSortByType = "description"
	GetAllStopsSortByDescriptionDescending        GetAllStopsSortByType = "-description"
	GetAllStopsSortByLatitudeAscending            GetAllStopsSortByType = "latitude"
	GetAllStopsSortByLatitudeDescending           GetAllStopsSortByType = "-latitude"
	GetAllStopsSortByLocationTypeAscending        GetAllStopsSortByType = "location_type"
	GetAllStopsSortByLocationTypeDescending       GetAllStopsSortByType = "-location_type"
	GetAllStopsSortByLongitudeAscending           GetAllStopsSortByType = "longitude"
	GetAllStopsSortByLongitudeDescending          GetAllStopsSortByType = "-longitude"
	GetAllStopsSortByNameAscending                GetAllStopsSortByType = "name"
	GetAllStopsSortByNameDescending               GetAllStopsSortByType = "-name"
	GetAllStopsSortByPlatformCodeAscending        GetAllStopsSortByType = "platform_code"
	GetAllStopsSortByPlatformCodeDescending       GetAllStopsSortByType = "-platform_code"
	GetAllStopsSortByPlatformNameAscending        GetAllStopsSortByType = "platform_name"
	GetAllStopsSortByPlatformNameDescending       GetAllStopsSortByType = "-platform_name"
	GetAllStopsSortByWheelchairBoardingAscending  GetAllStopsSortByType = "wheelchair_boarding"
	GetAllStopsSortByWheelchairBoardingDescending GetAllStopsSortByType = "-wheelchair_boarding"
	GetAllStopsSortByDistanceAscending            GetAllStopsSortByType = "distance"
	GetAllStopsSortByDistanceDescending           GetAllStopsSortByType = "-distance"
)

// GetAllStopsRequestConfig extra options for the GetAllStops request
type GetAllStopsRequestConfig struct {
	PageOffset           string                // Offset (0-based) of first element in the page
	PageLimit            string                // Max number of elements to return
	Sort                 GetAllStopsSortByType // Results can be sorted by the id or any GetAllStopsSortByType
	Fields               []string              // Fields to include with the response. Multiple fields MUST be a comma-separated (U+002C COMMA, “,”) list. Note that fields can also be selected for included data types
	IncludeParentStation bool                  // Include IncludeParentStation data in response
	FilterDirectionID    string                // Filter by Direction ID (Either "0" or "1")
	FilterLatitude       string                // Latitude in degrees North in the WGS-84 coordinate system to search filter[radius] degrees around with filter[longitude]
	FilterLongitude      string                // Longitude in degrees East in the WGS-84 coordinate system to search filter[radius] degrees around with filter[latitude]
	FilterRadius         string                // The distance is in degrees as if latitude and longitude were on a flat 2D plane and normal Pythagorean distance was calculated. Over the region MBTA serves, 0.02 degrees is approximately 1 mile. Defaults to 0.01 degrees (approximately a half mile)
	FilterIDs            []string              // Filter by multiple IDs
	FilterRouteTypes     []string              // Filter by route type(s)
	FilterRouteIDs       []string              // Filter by route IDs. If the vehicle is on a multi-route trip, it will be returned for any of the routes
	FilterLocationType   []string              // Filter by location type
}

func (config *GetAllStopsRequestConfig) addHTTPParamsToRequest(req *http.Request) {
	// Add fields and includes params to request
	includes := GetStopRequestConfig{
		IncludeParentStation: config.IncludeParentStation,
	}
	includes.addHTTPParamsToRequest(req)

	q := req.URL.Query()
	addToQuery(q, "page[offset]", config.PageOffset)
	addToQuery(q, "page[limit]", config.PageLimit)
	addToQuery(q, "sort", string(config.Sort))
	addToQuery(q, "filter[direction_id]", config.FilterDirectionID)
	addToQuery(q, "filter[latitude]", config.FilterLatitude)
	addToQuery(q, "filter[longitude]", config.FilterLongitude)
	addToQuery(q, "filter[radius]", config.FilterRadius)
	addCommaSeparatedListToQuery(q, "filter[id]", config.FilterIDs)
	addCommaSeparatedListToQuery(q, "filter[route_type]", config.FilterRouteTypes)
	addCommaSeparatedListToQuery(q, "filter[route]", config.FilterRouteIDs)
	addCommaSeparatedListToQuery(q, "filter[location_type]", config.FilterLocationType)

	req.URL.RawQuery = q.Encode()
}

// GetStopRequestConfig extra options for the GetStop request
type GetStopRequestConfig struct {
	Fields               []string // Fields to include with the response. Multiple fields MUST be a comma-separated (U+002C COMMA, “,”) list. Note that fields can also be selected for included data types
	IncludeParentStation bool     // Include IncludeParentStation data in response
}

func (config *GetStopRequestConfig) addHTTPParamsToRequest(req *http.Request) {
	q := req.URL.Query()

	if config.IncludeParentStation {
		q.Add("include", "parent_station")
	}
	addCommaSeparatedListToQuery(q, "fields[stop]", config.Fields)
	req.URL.RawQuery = q.Encode()
}
