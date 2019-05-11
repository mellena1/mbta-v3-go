package mbta

// StopLocationType enum for the possible stop location types
type StopLocationType int

const (
	// StopLocationSTOP A location where passengers board or disembark from a transit vehicle
	StopLocationSTOP StopLocationType = iota
	// StopLocationSTATION A physical structure or area that contains one or more stops
	StopLocationSTATION
	// StopLocationSTATIONENTRANCEEXIT A location where passengers can enter or exit a station from the street
	StopLocationSTATIONENTRANCEEXIT
)

// StopWheelchairBoardingType enum for the possible wheelchair boarding types at a stop
type StopWheelchairBoardingType int

const (
	// StopWheelchairBoardingNOINFO No information
	StopWheelchairBoardingNOINFO StopWheelchairBoardingType = iota
	// StopWheelchairBoardingACCESSIBLE Accessible (if trip is wheelchair accessible)
	StopWheelchairBoardingACCESSIBLE
	// StopWheelchairBoardingINACCESSIBLE Inaccessible
	StopWheelchairBoardingINACCESSIBLE
)

// Stop holds all info about a given MBTA Stop
type Stop struct {
	ID                 string                     `jsonapi:"primary,stop"`
	Address            *string                    `jsonapi:"attr,address"`             // A street address for the station
	Description        *string                    `jsonapi:"attr,description"`         // Description of the stop
	Latitute           float64                    `jsonapi:"attr,latitude"`            // Degrees North, in the WGS-84 coordinate system
	LocationType       StopLocationType           `jsonapi:"attr,location_type"`       // The type of the stop
	Longitude          float64                    `jsonapi:"attr,longitude"`           // Degrees East, in the WGS-84 coordinate system
	Name               string                     `jsonapi:"attr,name"`                // Name of a stop or station in the local and tourist vernacular
	PlatformCode       *string                    `jsonapi:"attr,platform_code"`       // A short code representing the platform/track (like a number or letter)
	PlatformName       *string                    `jsonapi:"attr,platform_name"`       // A textual description of the platform or track
	WheelchairBoarding StopWheelchairBoardingType `jsonapi:"attr,wheelchair_boarding"` // Whether there are any vehicles with wheelchair boarding or paths to stops that are wheelchair acessible
	ParentStation      *Stop                      `jsonapi:"relation,parent_station"`  // The link to the parent station. Only includes id by default, use IncludeParentStation config option to get all data
}
