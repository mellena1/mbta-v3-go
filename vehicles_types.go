package mbta

import (
	"encoding/json"
	"time"
)

// Vehicle holds all info about a given MBTA vehicle
type Vehicle struct {
	Attributes    VehicleAttributes    `json:"attributes"`
	Relationships VehicleRelationships `json:"relationships"`
	ID            string               `json:"id"`
}

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

// VehicleAttributes all attributes that a Vehicle has
type VehicleAttributes struct {
	Bearing             float32       `json:"bearing"`               // Bearing, in degrees, clockwise from True North, i.e., 0 is North and 90 is East
	CurrentStatus       VehicleStatus `json:"current_status"`        // Status of vehicle relative to the stops
	CurrentStopSequence int           `json:"current_stop_sequence"` // not sure on this one yet
	DirectionID         int           `json:"direction_id"`          // Direction in which trip is traveling: 0 or 1.
	Label               string        `json:"label"`                 // User visible label, such as the one of on the signage on the vehicle
	Latitute            float64       `json:"latitude"`              // Degrees North, in the WGS-84 coordinate system
	Longitude           float64       `json:"longitude"`             // Degrees East, in the WGS-84 coordinate system
	Speed               *float32      `json:"speed"`                 // meters per second
	UpdatedAt           time.Time     `json:"updated_at"`            // Time at which vehicle information was last updated. Format is ISO8601
}

// VehicleRelationships all relationships that a Vehicle has
type VehicleRelationships struct {
	RouteID string
	StopID  string
	TripID  string
}

// UnmarshalJSON custom JSON unmarshalling for the VehicleRelationships
func (v *VehicleRelationships) UnmarshalJSON(b []byte) error {
	type relationship struct {
		Data struct {
			ID string `json:"id"`
		}
	}
	type relationships struct {
		Route relationship `json:"route"`
		Stop  relationship `json:"stop"`
		Trip  relationship `json:"trip"`
	}

	var rels relationships
	err := json.Unmarshal(b, &rels)
	if err != nil {
		return err
	}
	v.RouteID = rels.Route.Data.ID
	v.StopID = rels.Stop.Data.ID
	v.TripID = rels.Trip.Data.ID

	return nil
}
