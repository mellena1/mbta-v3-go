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
	includeAlerts        = "alerts"
	includeLine          = "line"
	includeStop          = "stop"
	includeTrip          = "trip"
	includeRoute         = "route"
	includeParentStation = "parent_station"
	includeVehicle       = "vehicle"
	includeSchedule      = "schedule"
	includeService       = "service"
	includeShape         = "shape"
	includePrediction    = "prediction"
	includePredictions   = "predictions"
	includeRoutePattern  = "route_pattern"
	includeRoutePatterns = "route_patterns"
)
