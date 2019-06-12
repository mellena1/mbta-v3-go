package mbta

import (
	"context"
	"fmt"
	"net/http"
)

const servicesAPIPath = "/services"

type ServicesService service

// Weekday used to represent `valid_dates` in response
type Weekday int

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// ScheduleTypicality Describes how well this schedule represents typical service for the listed schedule_type
type Typicality int

const (
	Undefined Typicality = iota
	Typical
	SupplementalExtra
	ReducedHoliday
	PlannedMajorChange
	UnplannedMajorReduction
)

// Service holds all the info about a given MBTA Service
type Service struct {
	ID         string        `jsonapi:"primary,service"`
	AddedDates []TimeISO8601 `jsonapi:"attr,added_dates"`
	// AddedDatesNotes    []string      `jsonapi:"attr,added_dates_notes"`
	Description        string        `jsonapi:"attr,description"`
	EndDate            TimeISO8601   `jsonapi:"attr,end_date"`
	RemovedDates       []TimeISO8601 `jsonapi:"attr,removed_dates"`
	RemovedDatesNotes  []string      `jsonapi:"attr,removed_dates_notes"`
	ScheduleName       string        `jsonapi:"attr,schedule_name"`
	ScheduleType       string        `jsonapi:"attr,schedule_type"`
	ScheduleTypicality Typicality    `jsonapi:"attr,schedule_typicality"`
	StartDate          TimeISO8601   `jsonapi:"attr,start_date"`
	// ValidDays          []Weekday     `jsonapi:"attr,valid_days"`
}

// GetAllServicesSortByType all possible ways to sort /services request
type GetAllServicesSortByType string

const (
	GetAllServicesSortAddedDatesByAscending          GetAllServicesSortByType = "added_dates"
	GetAllServicesSortAddedDatesByDescending         GetAllServicesSortByType = "-added_dates"
	GetAllServicesSortAddedDatesNotesByAscending     GetAllServicesSortByType = "added_dates_notes"
	GetAllServicesSortAddedDatesNotesByDescending    GetAllServicesSortByType = "-added_dates_notes"
	GetAllServicesSortDescriptionByAscending         GetAllServicesSortByType = "description"
	GetAllServicesSortDescriptionByDescending        GetAllServicesSortByType = "-description"
	GetAllServicesSortEndDateByAscending             GetAllServicesSortByType = "end_date"
	GetAllServicesSortEndDateByDescending            GetAllServicesSortByType = "-end_date"
	GetAllServicesSortRemovedDatesByAscending        GetAllServicesSortByType = "removed_dates"
	GetAllServicesSortRemovedDatesByDescending       GetAllServicesSortByType = "-removed_dates"
	GetAllServicesSortRemovedDatesNotesByAscending   GetAllServicesSortByType = "removed_dates_notes"
	GetAllServicesSortRemovedDatesNotesByDescending  GetAllServicesSortByType = "-removed_dates_notes"
	GetAllServicesSortScheduleNameByAscending        GetAllServicesSortByType = "schedule_name"
	GetAllServicesSortScheduleNameByDescending       GetAllServicesSortByType = "-schedule_name"
	GetAllServicesSortScheduleTypeByAscending        GetAllServicesSortByType = "schedule_type"
	GetAllServicesSortScheduleTypeByDescending       GetAllServicesSortByType = "-schedule_type"
	GetAllServicesSortScheduleTypicalityByAscending  GetAllServicesSortByType = "schedule_typicality"
	GetAllServicesSortScheduleTypicalityByDescending GetAllServicesSortByType = "-schedule_typicality"
	GetAllServicesSortStartDateByAscending           GetAllServicesSortByType = "start_date"
	GetAllServicesSortStartDateByDescending          GetAllServicesSortByType = "-start_date"
	GetAllServicesSortValidDaysByAscending           GetAllServicesSortByType = "valid_days"
	GetAllServicesSortValidDaysByDescending          GetAllServicesSortByType = "-valid_days"
)

// GetAllServicesRequestConfig extra options for GetAllServices Request
type GetAllServicesRequestConfig struct {
	PageOffset   string                   `url:"page[offset],omitempty"`          // Offset (0-based) of first element in the page// Offset (0-based) of first element in the page
	PageLimit    string                   `url:"page[limit],omitempty"`           // Max number of elements to return// Max number of elements to return
	Sort         GetAllServicesSortByType `url:"sort,omitempty"`                  // Results can be sorted by the id or any GetAllRoutesSortByType
	Fields       []string                 `url:"fields[service],comma,omitempty"` // Fields to include with the response. Note that fields can also be selected for included data types// Fields to include with the response. Multiple fields MUST be a comma-separated (U+002C COMMA, “,”) list. Note that fields can also be selected for included data types
	FilterIDs    []string                 `url:"filter[id],comma,omitempty"`      // Filter by multiple IDs
	FilterRoutes []Route                  `url:"filter[route],comma,omitempty"`   // Filter by Routes
}

// GetAllServices returns all services from the mbta API
func (s *ServicesService) GetAllServices(config *GetAllServicesRequestConfig) ([]*Service, *http.Response, error) {
	return s.GetAllServicesWithContext(context.Background(), config)
}

// GetAllServicesWithContext returns all services from the mbta API given a context
func (s *ServicesService) GetAllServicesWithContext(ctx context.Context, config *GetAllServicesRequestConfig) ([]*Service, *http.Response, error) {
	u, err := addOptions(servicesAPIPath, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	untypedServices, resp, err := s.client.doManyPayload(req, &Service{})
	services := make([]*Service, len(untypedServices))
	for i := 0; i < len(untypedServices); i++ {
		services[i] = untypedServices[i].(*Service)
	}
	return services, resp, err
}

// GetServiceRequestConfig extra options for GetService Request
type GetServiceRequestConfig struct {
	Fields []string `url:"fields[service],comma,omitempty"` // Fields to include with the response. Note that fields can also be selected for included data types// Fields to include with the response. Multiple fields MUST be a comma-separated (U+002C COMMA, “,”) list. Note that fields can also be selected for included data types
}

// GetService returns a service from the mbta API
func (s *ServicesService) GetService(id string, config *GetServiceRequestConfig) (*Service, *http.Response, error) {
	return s.GetServiceWithContext(context.Background(), id, config)
}

// GetServiceWithContext returns a service from the mbta API given a context
func (s *ServicesService) GetServiceWithContext(ctx context.Context, id string, config *GetServiceRequestConfig) (*Service, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", servicesAPIPath, id)
	u, err := addOptions(path, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	var service Service
	resp, err := s.client.doSinglePayload(req, &service)
	return &service, resp, err
}
