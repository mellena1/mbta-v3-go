package mbta

// Trip holds all info about a given MBTA Trip
type Trip struct {
	ID                   string                 `jsonapi:"primary,trip"`
	WheelchairAccessible WheelchairBoardingType `jsonapi:"attr,wheelchair_accessible"` // Indicator of wheelchair accessibility
	Name                 string                 `jsonapi:"attr,name"`                  // The text that appears in schedules and sign boards to identify the trip to passengers
	Headsign             string                 `jsonapi:"attr,current_stop_sequence"` // The text that appears on a sign that identifies the trip’s destination to passengers
	DirectionID          int                    `jsonapi:"attr,direction_id"`          // Direction in which trip is traveling: 0 or 1.
	BlockID              string                 `jsonapi:"attr,block_id"`              // ID used to group sequential trips with the same vehicle for a given service_id
	BikesAllowed         BikesAllowedType       `jsonapi:"attr,bikes_allowed"`         // Indicator of whether or not bikes are allowed on this trip
}
