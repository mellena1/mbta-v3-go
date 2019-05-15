package mbta

// RouteType enum for possible Route types (see https://github.com/google/transit/blob/master/gtfs/spec/en/reference.md#routestxt)
type RouteType int

const (
	// RouteTypeLightRail ...
	RouteTypeLightRail RouteType = iota
	// RouteTypeSubway ...
	RouteTypeSubway
	// RouteTypeRail ...
	RouteTypeRail
	// RouteTypeBus ...
	RouteTypeBus
	// RouteTypeFerry ...
	RouteTypeFerry
	// RouteTypeCableCar ...
	RouteTypeCableCar
	// RouteTypeGondola ...
	RouteTypeGondola
	// RouteTypeFunicular ...
	RouteTypeFunicular
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
