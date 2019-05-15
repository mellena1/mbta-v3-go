package mbta

// VehicleStatus enum for the possible vehicle statuses
type VehicleStatus string

const (
	// InTransitTo vehicleStatus for a vehicle in transit
	InTransitTo VehicleStatus = "IN_TRANSIT_TO"
	// StoppedAt vehicleStatus for a vehicle stopped at a station
	StoppedAt VehicleStatus = "STOPPED_AT"
	// IncomingAt vehicleStatus for a vehicle getting into a station
	IncomingAt VehicleStatus = "INCOMING_AT"
)

// Vehicle holds all info about a given MBTA vehicle
type Vehicle struct {
	ID                  string        `jsonapi:"primary,vehicle"`
	Bearing             float32       `jsonapi:"attr,bearing"`               // Bearing, in degrees, clockwise from True North, i.e., 0 is North and 90 is East
	CurrentStatus       VehicleStatus `jsonapi:"attr,current_status"`        // Status of vehicle relative to the stops
	CurrentStopSequence int           `jsonapi:"attr,current_stop_sequence"` // not sure on this one yet
	DirectionID         int           `jsonapi:"attr,direction_id"`          // Direction in which trip is traveling: 0 or 1.
	Label               string        `jsonapi:"attr,label"`                 // User visible label, such as the one of on the signage on the vehicle
	Latitute            float64       `jsonapi:"attr,latitude"`              // Degrees North, in the WGS-84 coordinate system
	Longitude           float64       `jsonapi:"attr,longitude"`             // Degrees East, in the WGS-84 coordinate system
	Speed               *float32      `jsonapi:"attr,speed"`                 // meters per second
	UpdatedAt           TimeISO8601   `jsonapi:"attr,updated_at"`            // Time at which vehicle information was last updated. Format is ISO8601
	Stop                *Stop         `jsonapi:"relation,stop"`              // Stop that the vehicle is at. Only includes id by default, use IncludeStop config option to get all data
	Trip                *Trip         `jsonapi:"relation,trip"`              // Trip that the current vehicle is on
	// TODO: Route *Route `jsonapi:"relation,route"`
}
