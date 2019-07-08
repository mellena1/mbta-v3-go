package mbta

import (
	"context"
	"net/http"

	"golang.org/x/xerrors"
)

const schedulesAPIPath = "/schedules"

// ScheduleService service handling all of the schedule related API calls
type ScheduleService service

// ScheduleTimepoint whether times are exact or estimates
type ScheduleTimepoint bool

const (
	ScheduleTimepointExact     ScheduleTimepoint = true
	ScheduleTimepointEstimates ScheduleTimepoint = false
)

// SchedulePickupType whether times are exact or estimates
type SchedulePickupType int

const (
	SchedulePickupRegular SchedulePickupType = iota
	SchedulePickupNotAvailable
	SchedulePickupMustPhoneAgency
	SchedulePickupMustCoordinateWithDriver
)

// Schedule holds all info about a given MBTA schedule
type Schedule struct {
	ID            string             `jsonapi:"primary,schedule"`
	ArrivalTime   TimeISO8601        `jsonapi:"attr,arrival_time"`   // Time when the trip arrives at the given stop
	DepartureTime TimeISO8601        `jsonapi:"attr,departure_time"` // Time when the trip departs the given stop
	DirectionID   int                `jsonapi:"attr,direction_id"`   // Direction in which trip is traveling: 0 or 1.
	DropOffType   SchedulePickupType `jsonapi:"attr,drop_off_type"`  // How the vehicle arrives at stop_id
	PickupType    SchedulePickupType `jsonapi:"attr,pickup_type"`    // How the vehicle departs from stop_id.
	StopSequence  int                `jsonapi:"attr,stop_sequence"`  // The sequence the stop_id is arrived at during the trip_id. The stop sequence is monotonically increasing along the trip, but the stop_sequence along the trip_id are not necessarily consecutive
	Timepoint     ScheduleTimepoint  `jsonapi:"attr,timepoint"`      // whether the given times are exact or estimates
	Route         *Route             `jsonapi:"relation,route"`      // Route that the current schedule is linked with. Only includes id by default, use Include config option to get all data
	Stop          *Stop              `jsonapi:"relation,stop"`       // Stop that the schedule is linked with. Only includes id by default, use Include config option to get all data
	Trip          *Trip              `jsonapi:"relation,trip"`       // Trip that the current schedule is linked with. Only includes id by default, use Include config option to get all data
	Prediction    *Prediction        `jsonapi:"relation,prediction"`
}

// ScheduleInclude all of the includes for a schedule request
type ScheduleInclude string

const (
	ScheduleIncludePrediction ScheduleInclude = includePrediction
	ScheduleIncludeRoute      ScheduleInclude = includeRoute
	ScheduleIncludeStop       ScheduleInclude = includeStop
	ScheduleIncludeTrip       ScheduleInclude = includeTrip
)

// SchedulesSortByType all of the possible ways to sort by for a GetAllSchedules request
type SchedulesSortByType string

const (
	SchedulesSortByArrivalTimeAscending    SchedulesSortByType = "arrival_time"
	SchedulesSortByArrivalTimeDescending   SchedulesSortByType = "-arrival_time"
	SchedulesSortByDepartureTimeAscending  SchedulesSortByType = "departure_time"
	SchedulesSortByDepartureTimeDescending SchedulesSortByType = "-departure_time"
	SchedulesSortByDirectionIDAscending    SchedulesSortByType = "direction_id"
	SchedulesSortByDirectionIDDescending   SchedulesSortByType = "-direction_id"
	SchedulesSortByDropOffTypeAscending    SchedulesSortByType = "drop_off_type"
	SchedulesSortByDropOffTypeDescending   SchedulesSortByType = "-drop_off_type"
	SchedulesSortByPickupTypeAscending     SchedulesSortByType = "pickup_type"
	SchedulesSortByPickupTypeDescending    SchedulesSortByType = "-pickup_type"
	SchedulesSortByStopSequenceAscending   SchedulesSortByType = "stop_sequence"
	SchedulesSortByStopSequenceDescending  SchedulesSortByType = "-stop_sequence"
	SchedulesSortByTimepointAscending      SchedulesSortByType = "timepoint"
	SchedulesSortByTimepointDescending     SchedulesSortByType = "-timepoint"
)

// GetAllSchedulesRequestConfig extra options for the GetAllSchedules request
type GetAllSchedulesRequestConfig struct {
	PageOffset         string              `url:"page[offset],omitempty"`           // Offset (0-based) of first element in the page
	PageLimit          string              `url:"page[limit],omitempty"`            // Max number of elements to return
	Sort               SchedulesSortByType `url:"sort,omitempty"`                   // Results can be sorted by the id or any SchedulesSortByType
	Fields             []string            `url:"fields[schedule],comma,omitempty"` // Fields to include with the response. Note that fields can also be selected for included data types
	Include            []ScheduleInclude   `url:"include,comma,omitempty"`          // Include extra data in response (trip, stop, prediction, or route)
	FilterDates        []TimeISO8601       `url:"filter[date],comma,omitempty"`     // Filter by multiple dates
	FilterDirectionID  string              `url:"filter[direction_id],omitempty"`   // Filter by Direction ID (Either "0" or "1")
	FilterMinTime      []string            `url:"filter[min_time],comma,omitempty"` // Time before which schedule should not be returned. To filter times after midnight use more than 24 hours. For example, min_time=24:00 will return schedule information for the next calendar day, since that service is considered part of the current service day. Additionally, min_time=00:00&max_time=02:00 will not return anything. The time format is HH:MM.
	FilterMaxTime      []string            `url:"filter[max_time],comma,omitempty"` // Time after which schedule should not be returned. To filter times after midnight use more than 24 hours. For example, min_time=24:00 will return schedule information for the next calendar day, since that service is considered part of the current service day. Additionally, min_time=00:00&max_time=02:00 will not return anything. The time format is HH:MM.
	FilterRouteIDs     []string            `url:"filter[route],comma,omitempty"`    // Filter by route IDs
	FilterStopIDs      []string            `url:"filter[stop],comma,omitempty"`     // Filter by stop IDs
	FilterTripIDs      []string            `url:"filter[trip],comma,omitempty"`     // Filter by trip IDs
	FilterStopSequence string              `url:"filter[stop_sequence],omitempty"`  // Filter by the index of the stop in the trip. Symbolic values `first` and `last` can be used instead of numeric sequence number too.
}

// GetAllSchedules returns all schedules for a particular route, stop or trip from the mbta API
// NOTE: filter[route], filter[stop], or filter[trip] MUST be present for any schedules to be returned.
func (s *ScheduleService) GetAllSchedules(config *GetAllSchedulesRequestConfig) ([]*Schedule, *http.Response, error) {
	return s.GetAllSchedulesWithContext(context.Background(), config)
}

// GetAllSchedulesWithContext returns all schedules for a particular route, stop or trip from the mbta API given a context
// NOTE: filter[route], filter[stop], or filter[trip] MUST be present for any schedules to be returned.
func (s *ScheduleService) GetAllSchedulesWithContext(ctx context.Context, config *GetAllSchedulesRequestConfig) ([]*Schedule, *http.Response, error) {
	if len(config.FilterRouteIDs) == 0 && len(config.FilterStopIDs) == 0 && len(config.FilterTripIDs) == 0 {
		return nil, nil, xerrors.Errorf("Must filter by one of: RouteIDs, StopIDs, TripIDs: %w", ErrInvalidConfig)
	}

	u, err := addOptions(schedulesAPIPath, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	untypedSchedules, resp, err := s.client.doManyPayload(req, &Schedule{})
	schedules := make([]*Schedule, len(untypedSchedules))
	for i := 0; i < len(untypedSchedules); i++ {
		schedules[i] = untypedSchedules[i].(*Schedule)
	}
	return schedules, resp, err
}
