package mbta

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const facilitiesAPIPath = "/facilities"

// FacilityService service handling all of the facility related API calls
// Note: This spec is not yet finalized by the MBTA, so this may change/break depending on what the MBTA does.
type FacilityService service

// FacilityType enum for the possible facility types
type FacilityType string

const (
	FacilityBikeStorage                 FacilityType = "BIKE_STORAGE"
	FacilityBridgePlate                 FacilityType = "BRIDGE_PLATE"
	FacilityElectricCarChargers         FacilityType = "ELECTRIC_CAR_CHARGERS"
	FacilityElevatedSubplatform         FacilityType = "ELEVATED_SUBPLATFORM"
	FacilityElevator                    FacilityType = "ELEVATOR"
	FacilityEscalator                   FacilityType = "ESCALATOR"
	FacilityFareMediaAssistanceFacility FacilityType = "FARE_MEDIA_ASSISTANCE_FACILITY"
	FacilityFareMediaAssistant          FacilityType = "FARE_MEDIA_ASSISTANT"
	FacilityFareVendingMachine          FacilityType = "FARE_VENDING_MACHINE"
	FacilityFareVendingRetailer         FacilityType = "FARE_VENDING_RETAILER"
	FacilityFullyElevatedPlatform       FacilityType = "FULLY_ELEVATED_PLATFORM"
	FacilityOther                       FacilityType = "OTHER"
	FacilityParkingArea                 FacilityType = "PARKING_AREA"
	FacilityPickDrop                    FacilityType = "PICK_DROP"
	FacilityPortableBoardingLift        FacilityType = "PORTABLE_BOARDING_LIFT"
	FacilityRamp                        FacilityType = "RAMP"
	FacilityTaxiStand                   FacilityType = "TAXI_STAND"
	FacilityTicketWindow                FacilityType = "TICKET_WINDOW"
)

// FacilityProperty Name/value pair for additional facility information
type FacilityProperty struct {
	Name  string `json:"name"`  // The name of the property
	Value string `json:"value"` // The value of the property
}

// UnmarshalJSON unmarshal into FacilityProperty
func (f *FacilityProperty) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Name  string      `json:"name"`
		Value json.Number `json:"value"` // could be num or string
	}
	json.Unmarshal(b, &tmp)
	f.Name = tmp.Name
	f.Value = tmp.Value.String()
	return nil
}

// Facility holds all info about a given MBTA Facility
type Facility struct {
	ID         string             `jsonapi:"primary,facility"`
	Type       FacilityType       `jsonapi:"attr,type"`       // The type of the facility
	ShortName  string             `jsonapi:"attr,short_name"` // Short name of the facility
	Properties []FacilityProperty `jsonapi:"attr,properties"` // Name/value pair for additional facility information
	Name       string             `jsonapi:"attr,name"`       // Name of the facility
	Longitude  *float64           `jsonapi:"attr,longitude"`  // Longitude of the facility. Degrees East, in the WGS-84 coordinate system
	Latitude   *float64           `jsonapi:"attr,latitude"`   // Latitude of the facility. Degrees North, in the WGS-84 coordinate system
}

// FacilityInclude all of the includes for a facility request
type FacilityInclude string

const (
	FacilityIncludeStop FacilityInclude = includeStop
)

// GetAllFacilitiesSortByType all of the possible ways to sort by for a GetAllFacilities request
type GetAllFacilitiesSortByType string

const (
	GetAllFacilitiesSortByLatitudeAscending    GetAllFacilitiesSortByType = "latitude"
	GetAllFacilitiesSortByLatitudeDescending   GetAllFacilitiesSortByType = "-latitude"
	GetAllFacilitiesSortByLongitudeAscending   GetAllFacilitiesSortByType = "longitude"
	GetAllFacilitiesSortByLongitudeDescending  GetAllFacilitiesSortByType = "-longitude"
	GetAllFacilitiesSortByNameAscending        GetAllFacilitiesSortByType = "name"
	GetAllFacilitiesSortByNameDescending       GetAllFacilitiesSortByType = "-name"
	GetAllFacilitiesSortByPropertiesAscending  GetAllFacilitiesSortByType = "properties"
	GetAllFacilitiesSortByPropertiesDescending GetAllFacilitiesSortByType = "-properties"
	GetAllFacilitiesSortByShortNameAscending   GetAllFacilitiesSortByType = "short_name"
	GetAllFacilitiesSortByShortNameDescending  GetAllFacilitiesSortByType = "-short_name"
	GetAllFacilitiesSortByTypeAscending        GetAllFacilitiesSortByType = "type"
	GetAllFacilitiesSortByTypeDescending       GetAllFacilitiesSortByType = "-type"
)

// GetAllFacilitiesRequestConfig extra options for the GetAllFacilities request
type GetAllFacilitiesRequestConfig struct {
	PageOffset    string                     `url:"page[offset],omitempty"`           // Offset (0-based) of first element in the page
	PageLimit     string                     `url:"page[limit],omitempty"`            // Max number of elements to return
	Sort          GetAllFacilitiesSortByType `url:"sort,omitempty"`                   // Results can be sorted by the id or any GetAllStopsSortByType
	Fields        []string                   `url:"fields[facility],comma,omitempty"` // Fields to include with the response. Note that fields can also be selected for included data types
	Include       []FacilityInclude          `url:"include,comma,omitempty"`          // Include extra data in response
	FilterStopIDs []string                   `url:"filter[stop],comma,omitempty"`     // Filter by stop ID
	FilterTypes   []string                   `url:"filter[type],comma,omitempty"`     // Filter by multiple types
}

// GetAllFacilities returns all facilities from the mbta API
func (s *FacilityService) GetAllFacilities(config *GetAllFacilitiesRequestConfig) ([]*Facility, *http.Response, error) {
	return s.GetAllFacilitiesWithContext(context.Background(), config)
}

// GetAllFacilitiesWithContext returns all facilities from the mbta API given a context
func (s *FacilityService) GetAllFacilitiesWithContext(ctx context.Context, config *GetAllFacilitiesRequestConfig) ([]*Facility, *http.Response, error) {
	u, err := addOptions(facilitiesAPIPath, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	untypedFacilities, resp, err := s.client.doManyPayload(req, &Facility{})
	facilities := make([]*Facility, len(untypedFacilities))
	for i := 0; i < len(untypedFacilities); i++ {
		facilities[i] = untypedFacilities[i].(*Facility)
	}
	return facilities, resp, err
}

// GetFacilityRequestConfig extra options for the GetFacility request
type GetFacilityRequestConfig struct {
	Fields  []string      `url:"fields[facility],comma,omitempty"` // Fields to include with the response. Note that fields can also be selected for included data types
	Include []StopInclude `url:"include,comma,omitempty"`          // Include extra data in response (parentstation)
}

// GetFacility returns a facility from the mbta API
func (s *FacilityService) GetFacility(id string, config *GetFacilityRequestConfig) (*Facility, *http.Response, error) {
	return s.GetFacilityWithContext(context.Background(), id, config)
}

// GetFacilityWithContext returns a facility from the mbta API given a context
func (s *FacilityService) GetFacilityWithContext(ctx context.Context, id string, config *GetFacilityRequestConfig) (*Facility, *http.Response, error) {
	if id == "" {
		return nil, nil, ErrMustSpecifyID
	}

	path := fmt.Sprintf("%s/%s", facilitiesAPIPath, id)
	u, err := addOptions(path, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	var facility Facility
	resp, err := s.client.doSinglePayload(req, &facility)
	return &facility, resp, err
}
