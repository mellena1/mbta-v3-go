package mbta

import (
	"context"
	"net/http"

	"golang.org/x/xerrors"
)

const predictionsAPIPath = "/predictions"

// PredictionScheduleRelationshipType possible values for the ScheduleRelationship field in a Prediction
type PredictionScheduleRelationshipType string

const (
	ScheduleRelationshipAdded       PredictionScheduleRelationshipType = "ADDED"
	ScheduleRelationshipCancelled   PredictionScheduleRelationshipType = "CANCELLED"
	ScheduleRelationshipNoData      PredictionScheduleRelationshipType = "NO_DATA"
	ScheduleRelationshipSkipped     PredictionScheduleRelationshipType = "SKIPPED"
	ScheduleRelationshipUnscheduled PredictionScheduleRelationshipType = "UNSCHEDULED"
)

// PredictionService service handling all of the prediction related API calls
type PredictionService service

// Prediction holds all info about a given MBTA prediction
type Prediction struct {
	ID                   string                              `jsonapi:"primary,prediction"`
	ArrivalTime          *TimeISO8601                        `jsonapi:"attr,arrival_time"`          // Time when the trip arrives at the given stop
	DepartureTime        *TimeISO8601                        `jsonapi:"attr,departure_time"`        // Time when the trip departs the given stop
	DirectionID          int                                 `jsonapi:"attr,direction_id"`          // Direction in which trip is traveling: 0 or 1.
	ScheduleRelationship *PredictionScheduleRelationshipType `jsonapi:"attr,schedule_relationship"` // How the predicted stop relates to the Model.Schedule.t stops.
	Status               *string                             `jsonapi:"attr,status"`                // Status of the schedule
	StopSequence         int                                 `jsonapi:"attr,stop_sequence"`         // The sequence the stop_id is arrived at during the trip_id. The stop sequence is monotonically increasing along the trip, but the stop_sequence along the trip_id are not necessarily consecutive
	Route                *Route                              `jsonapi:"relation,route"`             // Route that the prediction is linked with. Only includes id by default, use Include config option to get all data
	Schedule             *Schedule                           `jsonapi:"relation,schedule"`          // Schedule that the prediction is linked with. Only includes id by default, use Include config option to get all data
	Stop                 *Stop                               `jsonapi:"relation,stop"`              // Stop that the prediction is linked with. Only includes id by default, use Include config option to get all data
	Trip                 *Trip                               `jsonapi:"relation,trip"`              // Trip that the prediction is linked with. Only includes id by default, use Include config option to get all data
	Vehicle              *Vehicle                            `jsonapi:"relation,vehicle"`           // Vehicle that the prediction is linked with. Only includes id by default, use Include config option to get all data
	Alerts               []*Alert                            `jsonapi:"relation,alerts"`
}

// PredictionInclude all of the includes for a prediction request
type PredictionInclude string

const (
	PredictionIncludeSchedule PredictionInclude = includeSchedule
	PredictionIncludeStop     PredictionInclude = includeStop
	PredictionIncludeRoute    PredictionInclude = includeRoute
	PredictionIncludeTrip     PredictionInclude = includeTrip
	PredictionIncludeVehicle  PredictionInclude = includeVehicle
	PredictionIncludeAlerts   PredictionInclude = includeAlerts
)

// GetAllPredictionsSortByType all of the possible ways to sort by for a GetAllPredictions request
type GetAllPredictionsSortByType string

const (
	GetAllPredictionsSortByArrivalTimeAscending           GetAllPredictionsSortByType = "arrival_time"
	GetAllPredictionsSortByArrivalTimeDescending          GetAllPredictionsSortByType = "-arrival_time"
	GetAllPredictionsSortByDepartureTimeAscending         GetAllPredictionsSortByType = "departure_time"
	GetAllPredictionsSortByDepartureTimeDescending        GetAllPredictionsSortByType = "-departure_time"
	GetAllPredictionsSortByDirectionIDAscending           GetAllPredictionsSortByType = "direction_id"
	GetAllPredictionsSortByDirectionIDDescending          GetAllPredictionsSortByType = "-direction_id"
	GetAllPredictionsSortByScheduleRelationshipAscending  GetAllPredictionsSortByType = "schedule_relationship"
	GetAllPredictionsSortByScheduleRelationshipDescending GetAllPredictionsSortByType = "-schedule_relationship"
	GetAllPredictionsSortByStatusAscending                GetAllPredictionsSortByType = "status"
	GetAllPredictionsSortByStatusDescending               GetAllPredictionsSortByType = "-status"
	GetAllPredictionsSortByStopSequenceAscending          GetAllPredictionsSortByType = "stop_sequence"
	GetAllPredictionsSortByStopSequenceDescending         GetAllPredictionsSortByType = "-stop_sequence"
)

// GetAllPredictionsRequestConfig extra options for the GetAllPredictions request
type GetAllPredictionsRequestConfig struct {
	PageOffset        string                      `url:"page[offset],omitempty"`             // Offset (0-based) of first element in the page
	PageLimit         string                      `url:"page[limit],omitempty"`              // Max number of elements to return
	Sort              GetAllPredictionsSortByType `url:"sort,omitempty"`                     // Results can be sorted by the id or any GetAllPredictionsSortByType
	Fields            []string                    `url:"fields[prediction],comma,omitempty"` // Fields to include with the response. Note that fields can also be selected for included data types
	Include           []PredictionInclude         `url:"include,comma,omitempty"`            // Include extra data in response
	FilterLatitude    string                      `url:"filter[latitude],omitempty"`         // Latitude/Longitude must be both present or both absent
	FilterLongitude   string                      `url:"filter[longitude],omitempty"`        // Latitude/Longitude must be both present or both absent
	FilterRadius      string                      `url:"filter[radius],omitempty"`           // Radius accepts a floating point number, and the default is 0.01. For example, if you query for: latitude: 42, longitude: -71, radius: 0.05 then you will filter between latitudes 41.95 and 42.05, and longitudes -70.95 and -71.05
	FilterDirectionID string                      `url:"filter[direction_id],omitempty"`     // Filter by Direction ID (Either "0" or "1")
	FilterRouteType   []string                    `url:"filter[route_type],comma,omitempty"` // Filter by route_type
	FilterRouteIDs    []string                    `url:"filter[route],comma,omitempty"`      // Filter by route IDs
	FilterStopIDs     []string                    `url:"filter[stop],comma,omitempty"`       // Filter by stop IDs
	FilterTripIDs     []string                    `url:"filter[trip],comma,omitempty"`       // Filter by trip IDs
}

// GetAllPredictions returns all predictions from the mbta API
// NOTE: A filter MUST be present for any predictions to be returned.
func (s *PredictionService) GetAllPredictions(config *GetAllPredictionsRequestConfig) ([]*Prediction, *http.Response, error) {
	return s.GetAllPredictionsWithContext(context.Background(), config)
}

// GetAllPredictionsWithContext returns all predictions from the mbta API given a context
// NOTE: A filter MUST be present for any predictions to be returned.
func (s *PredictionService) GetAllPredictionsWithContext(ctx context.Context, config *GetAllPredictionsRequestConfig) ([]*Prediction, *http.Response, error) {
	isEmptyString := func(s string) bool { return s == "" }
	isEmptySlice := func(s []string) bool { return len(s) == 0 }
	// Ensure that a filter is set
	if isEmptySlice(config.FilterRouteType) && isEmptySlice(config.FilterRouteIDs) && isEmptySlice(config.FilterStopIDs) && isEmptySlice(config.FilterTripIDs) && isEmptyString(config.FilterLatitude) && isEmptyString(config.FilterLongitude) && isEmptyString(config.FilterRadius) && isEmptyString(config.FilterDirectionID) {
		return nil, nil, xerrors.Errorf("A filter must be present for any predictions to be returned: %w", ErrInvalidConfig)
	}

	u, err := addOptions(predictionsAPIPath, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	untypedPredictions, resp, err := s.client.doManyPayload(req, &Prediction{})
	predictions := make([]*Prediction, len(untypedPredictions))
	for i := 0; i < len(untypedPredictions); i++ {
		predictions[i] = untypedPredictions[i].(*Prediction)
	}
	return predictions, resp, err
}
