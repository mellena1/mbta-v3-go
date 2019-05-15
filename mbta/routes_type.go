package mbta

// RouteType enum for possible Route types
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
	ID          string
	AgencyID    string
	ShortName   string
	LongName    string
	Description string
	Type        RouteType
	Color       string
	TextColor   string
	SortOrder   int
}
