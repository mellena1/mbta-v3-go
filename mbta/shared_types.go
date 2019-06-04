package mbta

// WheelchairBoardingType enum for the possible wheelchair boarding types at a stop
type WheelchairBoardingType int

const (
	// WheelchairBoardingNOINFO No information
	WheelchairBoardingNOINFO WheelchairBoardingType = iota
	// WheelchairBoardingACCESSIBLE Accessible (if trip is wheelchair accessible)
	WheelchairBoardingACCESSIBLE
	// WheelchairBoardingINACCESSIBLE Inaccessible
	WheelchairBoardingINACCESSIBLE
)

// BikesAllowedType enum for whether or not bikes are allowed
type BikesAllowedType int

const (
	// BikesAllowedNOINFO No information
	BikesAllowedNOINFO BikesAllowedType = iota
	// BikesAllowedYES Vehicle being used on this particular trip can accommodate at least one bicycle
	BikesAllowedYES
	// BikesAllowedNO No bicycles are allowed on this trip
	BikesAllowedNO
)

const (
	includeLine          = "line"
	includeStop          = "stop"
	includeTrip          = "trip"
	includeRoute         = "route"
	includeParentStation = "parent_station"
	includeVehicle       = "vehicle"
	includeService       = "service"
	includeShape         = "shape"
	includePrediction    = "prediction"
	includePredictions   = "predictions"
	includeRoutePattern  = "route_pattern"
	includeRoutePatterns = "route_patterns"
)
