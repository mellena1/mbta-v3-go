package mbta

import (
	"encoding/json"
)

// Stop holds all info about a given MBTA Stop
type Stop struct {
	Attributes    StopAttributes    `json:"attributes"`
	Relationships StopRelationships `json:"relationships"`
	ID            string            `json:"id"`
}

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

// StopAttributes all attributes that a Stop has
type StopAttributes struct {
	Address            *string                    `json:"address"`             // A street address for the station
	Description        VehicleStatus              `json:"description"`         // Description of the stop
	Latitute           float64                    `json:"latitude"`            // Degrees North, in the WGS-84 coordinate system
	LocationType       StopLocationType           `json:"location_type"`       // The type of the stop
	Longitude          float64                    `json:"longitude"`           // Degrees East, in the WGS-84 coordinate system
	Name               string                     `json:"name"`                // Name of a stop or station in the local and tourist vernacular
	PlatformCode       *string                    `json:"platform_code"`       // A short code representing the platform/track (like a number or letter)
	PlatformName       *string                    `json:"platform_name"`       // A textual description of the platform or track
	WheelchairBoarding StopWheelchairBoardingType `json:"wheelchair_boarding"` // Whether there are any vehicles with wheelchair boarding or paths to stops that are wheelchair acessible
}

// StopRelationships all relationships that a Stop has
type StopRelationships struct {
	ParentStationID string
}

// UnmarshalJSON custom JSON unmarshalling for the StopRelationships
func (s *StopRelationships) UnmarshalJSON(b []byte) error {
	type relationship struct {
		Data *struct {
			ID string `json:"id"`
		}
	}
	type relationships struct {
		ParentStation relationship `json:"parent_station"`
	}

	var rels relationships
	err := json.Unmarshal(b, &rels)
	if err != nil {
		return err
	}
	if rels.ParentStation.Data != nil {
		s.ParentStationID = rels.ParentStation.Data.ID
	}

	return nil
}
