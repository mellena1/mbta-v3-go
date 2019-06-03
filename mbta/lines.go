package mbta

const linesAPIPath = "/lines"

type LineService service

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

type GetAllLinesRequestConfig Struct {
	PageOffset        string                 `url:"page[offset],omitempty"`         // Offset (0-based) of first element in the page// Offset (0-based) of first element in the page
	PageLimit         string                 `url:"page[limit],omitempty"`          // Max number of elements to return// Max number of elements to return
	Sort              GetAllLinesSortByType  `url:"sort,omitempty"`                 // Results can be sorted by the id or any GetAllRoutesSortByType
	Fields            []string               `url:"fields[line],comma,omitempty"`   // Fields to include with the response. Note that fields can also be selected for included data types// Fields to include with the response. Multiple fields MUST be a comma-separated (U+002C COMMA, “,”) list. Note that fields can also be selected for included data types
	Include			  []LineInclude		     `url:"include,comma,omitempty"`        // Include extra data in response
	FilterIDs         []string               `url:"filter[id],comma,omitempty"`     // Filter by multiple IDs
}

type GetLineRequestConfig struct {
	Fields            []string               `url:"fields[line],comma,omitempty"`   // Fields to include with the response. Note that fields can also be selected for included data types// Fields to include with the response. Multiple fields MUST be a comma-separated (U+002C COMMA, “,”) list. Note that fields can also be selected for included data types
	Include			  []LineInclude		     `url:"include,comma,omitempty"`        // Include extra data in response
}