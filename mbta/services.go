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
type ScheduleTypicality int

const (
	Undefined ScheduleTypicality = iota
	Typical
	SupplementalExtra
	ReducedHoliday
	PlannedMajorChange
	UnplannedMajorReduction
)

// Service holds all the info about a given MBTA Service
type Service struct {
	ID                 string             `jsonapi:"primary,service"`
	AddedDates         []TimeISO8601      `jsonapi:"attr,added_dates"`
	AddedDatesNotes    []*string          `jsonapi:"attr,added_dates_notes"`
	Description        string             `jsonapi:"attr,description"`
	EndDate            TimeISO8601        `jsonapi:"attr,end_date"`
	RemovedDates       []TimeISO8601      `jsonapi:"attr,removed_dates"`
	RemovedDatesNotes  []string           `jsonapi:"attr,removed_dates_notes"`
	ScheduleName       string             `jsonapi:"attr,schedule_name"`
	ScheduleType       string             `jsonapi:"attr,schedule_type"`
	ScheduleTypicality ScheduleTypicality `jsonapi:"attr,schedule_typicality"`
	StartDate          TimeISO8601        `jsonapi:"attr,start_date"`
	ValidDays          []Weekday          `jsonapi:"attr,valid_days"`
}

// ServicesSortByType all possible ways to sort /services request
type ServicesSortByType string

const (
	ServicesSortAddedDatesByAscending          ServicesSortByType = "added_dates"
	ServicesSortAddedDatesByDescending         ServicesSortByType = "-added_dates"
	ServicesSortAddedDatesNotesByAscending     ServicesSortByType = "added_dates_notes"
	ServicesSortAddedDatesNotesByDescending    ServicesSortByType = "-added_dates_notes"
	ServicesSortDescriptionByAscending         ServicesSortByType = "description"
	ServicesSortDescriptionByDescending        ServicesSortByType = "-description"
	ServicesSortEndDateByAscending             ServicesSortByType = "end_date"
	ServicesSortEndDateByDescending            ServicesSortByType = "-end_date"
	ServicesSortRemovedDatesByAscending        ServicesSortByType = "removed_dates"
	ServicesSortRemovedDatesByDescending       ServicesSortByType = "-removed_dates"
	ServicesSortRemovedDatesNotesByAscending   ServicesSortByType = "removed_dates_notes"
	ServicesSortRemovedDatesNotesByDescending  ServicesSortByType = "-removed_dates_notes"
	ServicesSortScheduleNameByAscending        ServicesSortByType = "schedule_name"
	ServicesSortScheduleNameByDescending       ServicesSortByType = "-schedule_name"
	ServicesSortScheduleTypeByAscending        ServicesSortByType = "schedule_type"
	ServicesSortScheduleTypeByDescending       ServicesSortByType = "-schedule_type"
	ServicesSortScheduleTypicalityByAscending  ServicesSortByType = "schedule_typicality"
	ServicesSortScheduleTypicalityByDescending ServicesSortByType = "-schedule_typicality"
	ServicesSortStartDateByAscending           ServicesSortByType = "start_date"
	ServicesSortStartDateByDescending          ServicesSortByType = "-start_date"
	ServicesSortValidDaysByAscending           ServicesSortByType = "valid_days"
	ServicesSortValidDaysByDescending          ServicesSortByType = "-valid_days"
)

// GetAllServicesRequestConfig extra options for GetAllServices Request
type GetAllServicesRequestConfig struct {
	PageOffset   string             `url:"page[offset],omitempty"`          // Offset (0-based) of first element in the page
	PageLimit    string             `url:"page[limit],omitempty"`           // Max number of elements to return
	Sort         ServicesSortByType `url:"sort,omitempty"`                  // Results can be sorted by the id or any RoutesSortByType
	Fields       []string           `url:"fields[service],comma,omitempty"` // Fields to include with the response. Note that fields can also be selected for included data types
	FilterIDs    []string           `url:"filter[id],comma,omitempty"`      // Filter by multiple IDs
	FilterRoutes []Route            `url:"filter[route],comma,omitempty"`   // Filter by Routes
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
	Fields []string `url:"fields[service],comma,omitempty"` // Fields to include with the response. Note that fields can also be selected for included data types
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
