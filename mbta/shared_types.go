package mbta

// WheelchairBoardingType enum for the possible wheelchair boarding types at a stop
type WheelchairBoardingType int

const (
	// WheelchairBoardingNoInfo No information
	WheelchairBoardingNoInfo WheelchairBoardingType = iota
	// WheelchairBoardingAccessible Accessible (if trip is wheelchair accessible)
	WheelchairBoardingAccessible
	// WheelchairBoardingInaccessible Inaccessible
	WheelchairBoardingInaccessible
)

// BikesAllowedType enum for whether or not bikes are allowed
type BikesAllowedType int

const (
	// BikesAllowedNoInfo No information
	BikesAllowedNoInfo BikesAllowedType = iota
	// BikesAllowedYes Vehicle being used on this particular trip can accommodate at least one bicycle
	BikesAllowedYes
	// BikesAllowedNo No bicycles are allowed on this trip
	BikesAllowedNo
)

const (
	includeLine               = "line"
	includeStop               = "stop"
  includeStops              = "stops"
	includeTrip               = "trip"
  includeTrips              = "trips"
	includeRoute              = "route"
	includeRoutes             = "routes"
	includeParentStation      = "parent_station"
	includeVehicle            = "vehicle"
	includeService            = "service"
	includeShape              = "shape"
	includePrediction         = "prediction"
	includePredictions        = "predictions"
	includeRepresentativeTrip = "representative_trip"
	includeRoutePattern       = "route_pattern"
	includeRoutePatterns      = "route_patterns"
  includeFacilities         = "facilities"
)
