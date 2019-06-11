package mbta

import (
	"context"
	"fmt"
	"net/http"
)

const alertsAPIPath = "/alerts"

// AlertService handling all of the alert related API calls
type AlertService service

// AlertLifecycleType Identifies whether alert is a new or old, in effect or upcoming
type AlertLifecycleType string

const (
	AlertLifecycleNew             AlertLifecycleType = "NEW"
	AlertLifecycleOngoing         AlertLifecycleType = "ONGOING"
	AlertLifecycleOngoingUpcoming AlertLifecycleType = "ONGOING_UPCOMING"
	AlertLifecycleUpcoming        AlertLifecycleType = "UPCOMING"
)

// AlertActivityType An activity affected by an alert
type AlertActivityType string

const (
	// AlertActivityBoard Boarding a vehicle. Any passenger trip includes boarding a vehicle and exiting from a vehicle
	AlertActivityBoard AlertActivityType = "BOARD"
	// AlertActivityBringingBike Bringing a bicycle while boarding or exiting
	AlertActivityBringingBike AlertActivityType = "BRINGING_BIKE"
	// AlertActivityExit Exiting from a vehicle (disembarking). Any passenger trip includes boarding a vehicle and exiting a vehicle
	AlertActivityExit AlertActivityType = "EXIT"
	// AlertActivityParkCar Parking a car at a garage or lot in a station
	AlertActivityParkCar AlertActivityType = "PARK_CAR"
	// AlertActivityRide Riding through a stop without boarding or exiting… Not every passenger trip will include this – a passenger may board at one stop and exit at the next stop
	AlertActivityRide AlertActivityType = "RIDE"
	// AlertActivityStoreBike Storing a bicycle at a station
	AlertActivityStoreBike AlertActivityType = "STORE_BIKE"
	// AlertActivityUsingEscalator Using an escalator while boarding or exiting (should only be used for customers who specifically want to avoid stairs.)
	AlertActivityUsingEscalator AlertActivityType = "USING_ESCALATOR"
	// AlertActivityUsingWheelchair Using a wheelchair while boarding or exiting. Note that this applies to something that specifically affects customers who use a wheelchair to board or exit; a delay should not include this as an affected activity unless it specifically affects customers using wheelchairs
	AlertActivityUsingWheelchair AlertActivityType = "USING_WHEELCHAIR"
	// AlertActivityFilterAll Filter by all Activities
	AlertActivityFilterAll AlertActivityType = "ALL"
)

// AlertEffectType The effect of this problem on the affected entity
type AlertEffectType string

const (
	AlertEffectAccessIssue       AlertEffectType = "ACCESS_ISSUE"
	AlertEffectAdditionalService AlertEffectType = "ADDITIONAL_SERVICE"
	AlertEffectAmberAlert        AlertEffectType = "AMBER_ALERT"
	AlertEffectBikeIssue         AlertEffectType = "BIKE_ISSUE"
	AlertEffectCancellation      AlertEffectType = "CANCELLATION"
	AlertEffectDelay             AlertEffectType = "DELAY"
	AlertEffectDetour            AlertEffectType = "DETOUR"
	AlertEffectDockClosure       AlertEffectType = "DOCK_CLOSURE"
	AlertEffectDockIssue         AlertEffectType = "DOCK_ISSUE"
	AlertEffectElevatorClosure   AlertEffectType = "ELEVATOR_CLOSURE"
	AlertEffectEscalatorClosure  AlertEffectType = "ESCALATOR_CLOSURE"
	AlertEffectExtraService      AlertEffectType = "EXTRA_SERVICE"
	AlertEffectFacilityIssue     AlertEffectType = "FACILITY_ISSUE"
	AlertEffectModifiedService   AlertEffectType = "MODIFIED_SERVICE"
	AlertEffectNoService         AlertEffectType = "NO_SERVICE"
	AlertEffectOtherEffect       AlertEffectType = "OTHER_EFFECT"
	AlertEffectParkingClosure    AlertEffectType = "PARKING_CLOSURE"
	AlertEffectParkingIssue      AlertEffectType = "PARKING_ISSUE"
	AlertEffectPolicyChange      AlertEffectType = "POLICY_CHANGE"
	AlertEffectScheduleChange    AlertEffectType = "SCHEDULE_CHANGE"
	AlertEffectServiceChange     AlertEffectType = "SERVICE_CHANGE"
	AlertEffectShuttle           AlertEffectType = "SHUTTLE"
	AlertEffectSnowRoute         AlertEffectType = "SNOW_ROUTE"
	AlertEffectStationClosure    AlertEffectType = "STATION_CLOSURE"
	AlertEffectStationIssue      AlertEffectType = "STATION_ISSUE"
	AlertEffectStopClosure       AlertEffectType = "STOP_CLOSURE"
	AlertEffectStopMove          AlertEffectType = "STOP_MOVE"
	AlertEffectStopMoved         AlertEffectType = "STOP_MOVED"
	AlertEffectSummary           AlertEffectType = "SUMMARY"
	AlertEffectSuspension        AlertEffectType = "SUSPENSION"
	AlertEffectTrackChange       AlertEffectType = "TRACK_CHANGE"
	AlertEffectUnknownEffect     AlertEffectType = "UNKNOWN_EFFECT"
)

// AlertCauseType What is causing the alert
type AlertCauseType string

const (
	AlertCauseAccident                   AlertCauseType = "ACCIDENT"
	AlertCauseAmtrak                     AlertCauseType = "AMTRAK"
	AlertCauseAnEarlierMechanicalProblem AlertCauseType = "AN_EARLIER_MECHANICAL_PROBLEM"
	AlertCauseAnEarlierSignalProblem     AlertCauseType = "AN_EARLIER_SIGNAL_PROBLEM"
	AlertCauseAutosImpedingService       AlertCauseType = "AUTOS_IMPEDING_SERVICE"
	AlertCauseCoastGuardRestriction      AlertCauseType = "COAST_GUARD_RESTRICTION"
	AlertCauseCongestion                 AlertCauseType = "CONGESTION"
	AlertCauseConstruction               AlertCauseType = "CONSTRUCTION"
	AlertCauseCrossingMalfunction        AlertCauseType = "CROSSING_MALFUNCTION"
	AlertCauseDemonstration              AlertCauseType = "DEMONSTRATION"
	AlertCauseDisabledBus                AlertCauseType = "DISABLED_BUS"
	AlertCauseDisabledTrain              AlertCauseType = "DISABLED_TRAIN"
	AlertCauseDrawbridgeBeingRaised      AlertCauseType = "DRAWBRIDGE_BEING_RAISED"
	AlertCauseElectricalWork             AlertCauseType = "ELECTRICAL_WORK"
	AlertCauseFire                       AlertCauseType = "FIRE"
	AlertCauseFog                        AlertCauseType = "FOG"
	AlertCauseFreightTrainInterference   AlertCauseType = "FREIGHT_TRAIN_INTERFERENCE"
	AlertCauseHazmatCondition            AlertCauseType = "HAZMAT_CONDITION"
	AlertCauseHeavyRidership             AlertCauseType = "HEAVY_RIDERSHIP"
	AlertCauseHighWinds                  AlertCauseType = "HIGH_WINDS"
	AlertCauseHoliday                    AlertCauseType = "HOLIDAY"
	AlertCauseHurricane                  AlertCauseType = "HURRICANE"
	AlertCauseIceInHarbor                AlertCauseType = "ICE_IN_HARBOR"
	AlertCauseMaintenance                AlertCauseType = "MAINTENANCE"
	AlertCauseMechanicalProblem          AlertCauseType = "MECHANICAL_PROBLEM"
	AlertCauseMedicalEmergency           AlertCauseType = "MEDICAL_EMERGENCY"
	AlertCauseParade                     AlertCauseType = "PARADE"
	AlertCausePoliceAction               AlertCauseType = "POLICE_ACTION"
	AlertCausePowerProblem               AlertCauseType = "POWER_PROBLEM"
	AlertCauseSevereWeather              AlertCauseType = "SEVERE_WEATHER"
	AlertCauseSignalProblem              AlertCauseType = "SIGNAL_PROBLEM"
	AlertCauseSlipperyRail               AlertCauseType = "SLIPPERY_RAIL"
	AlertCauseSnow                       AlertCauseType = "SNOW"
	AlertCauseSpecialEvent               AlertCauseType = "SPECIAL_EVENT"
	AlertCauseSpeedRestriction           AlertCauseType = "SPEED_RESTRICTION"
	AlertCauseSwitchProblem              AlertCauseType = "SWITCH_PROBLEM"
	AlertCauseTieReplacement             AlertCauseType = "TIE_REPLACEMENT"
	AlertCauseTrackProblem               AlertCauseType = "TRACK_PROBLEM"
	AlertCauseTrackWork                  AlertCauseType = "TRACK_WORK"
	AlertCauseTraffic                    AlertCauseType = "TRAFFIC"
	AlertCauseUnknownCause               AlertCauseType = "UNKNOWN_CAUSE"
	AlertCauseUnrulyPassenger            AlertCauseType = "UNRULY_PASSENGER"
	AlertCauseWeather                    AlertCauseType = "WEATHER"
)

// Alert holds all the info about a given MBTA Alert
type Alert struct {
	ID             string                `jsonapi:"primary,alert"`
	URL            *JSONURL              `jsonapi:"attr,url"`             // A URL for extra details, such as outline construction or maintenance plans
	UpdatedAt      TimeISO8601           `jsonapi:"attr,updated_at"`      // Date/Time alert last updated
	Timeframe      *string               `jsonapi:"attr,timeframe"`       // Summarizes when an alert is in effect.
	ShortHeader    string                `jsonapi:"attr,short_header"`    // A shortened version of */attributes/header.
	Severity       int                   `jsonapi:"attr,severity"`        // How severe the alert it from least (0) to most (10) severe.
	ServiceEffect  string                `jsonapi:"attr,service_effect"`  // Summarizes the service and the impact to that service
	Lifecycle      AlertLifecycleType    `jsonapi:"attr,lifecycle"`       // Identifies whether alert is a new or old, in effect or upcoming
	InformedEntity []AlertInformedEntity `jsonapi:"attr,informed_entity"` // Object representing a particular part of the system affected by an alert
	Header         string                `jsonapi:"attr,header"`          // This plain-text string will be highlighted, for example in boldface
	Effect         AlertEffectType       `jsonapi:"attr,effect_name"`     // The effect of this problem on the affected entity
	Description    *string               `jsonapi:"attr,description"`     // This plain-text string will be formatted as the body of the alert (or shown on an explicit “expand” request by the user). The information in the description should add to the information of the header
	CreatedAt      TimeISO8601           `jsonapi:"attr,created_at"`      // Date/Time alert created
	Cause          AlertCauseType        `jsonapi:"attr,cause"`           // What is causing the alert
	Banner         *string               `jsonapi:"attr,banner"`          // Set if alert is meant to be displayed prominently, such as the top of every page
	ActivePeriod   []AlertActivePeriod   `jsonapi:"attr,active_period"`   // Date/Time ranges when alert is active
}

// AlertInformedEntity Object representing a particular part of the system affected by an alert
type AlertInformedEntity struct {
	TripID      *string             `json:"trip"`
	StopID      *string             `json:"stop"`
	RouteType   *RouteType          `json:"route_type"`
	RouteID     *string             `json:"route"`
	FacilityID  *string             `json:"facility"`
	DirectionID *int                `json:"direction_id"`
	Activities  []AlertActivityType `json:"activities"`
}

// AlertActivePeriod Date/Time ranges when alert is active
type AlertActivePeriod struct {
	Start TimeISO8601  `json:"start"` // Start Date
	End   *TimeISO8601 `json:"end"`   // End Date
}

// AlertInclude all of the includes for a alert request
type AlertInclude string

const (
	AlertIncludeStops      AlertInclude = includeStops
	AlertIncludeRoutes     AlertInclude = includeRoutes
	AlertIncludeTrips      AlertInclude = includeTrips
	AlertIncludeFacilities AlertInclude = includeFacilities
)

// GetAllAlertsSortByType all of the possible ways to sort by for a GetAllAlerts request
type GetAllAlertsSortByType string

const (
	GetAllAlertsSortByActivePeriodAscending    GetAllAlertsSortByType = "active_period"
	GetAllAlertsSortByActivePeriodDescending   GetAllAlertsSortByType = "-active_period"
	GetAllAlertsSortByBannerDesending          GetAllAlertsSortByType = "banner"
	GetAllAlertsSortByBannerDescending         GetAllAlertsSortByType = "-banner"
	GetAllAlertsSortByCauseAscending           GetAllAlertsSortByType = "cause"
	GetAllAlertsSortByCauseDescending          GetAllAlertsSortByType = "-cause"
	GetAllAlertsSortByCreatedAtAscending       GetAllAlertsSortByType = "created_at"
	GetAllAlertsSortByCreatedAtDescending      GetAllAlertsSortByType = "-created_at"
	GetAllAlertsSortByDescriptionAscending     GetAllAlertsSortByType = "description"
	GetAllAlertsSortByDescriptionDescending    GetAllAlertsSortByType = "-description"
	GetAllAlertsSortByEffectAscending          GetAllAlertsSortByType = "effect"
	GetAllAlertsSortByEffectDescending         GetAllAlertsSortByType = "-effect"
	GetAllAlertsSortByHeaderAscending          GetAllAlertsSortByType = "header"
	GetAllAlertsSortByHeaderDescending         GetAllAlertsSortByType = "-header"
	GetAllAlertsSortByInformedEntityAscending  GetAllAlertsSortByType = "informed_entity"
	GetAllAlertsSortByInformedEntityDescending GetAllAlertsSortByType = "-informed_entity"
	GetAllAlertsSortByLifecycleAscending       GetAllAlertsSortByType = "lifecycle"
	GetAllAlertsSortByLifecycleDescending      GetAllAlertsSortByType = "-lifecycle"
	GetAllAlertsSortByServiceEffectAscending   GetAllAlertsSortByType = "service_effect"
	GetAllAlertsSortByServiceEffectDescending  GetAllAlertsSortByType = "-service_effect"
	GetAllAlertsSortBySeverityAscending        GetAllAlertsSortByType = "severity"
	GetAllAlertsSortBySeverityDescending       GetAllAlertsSortByType = "-severity"
	GetAllAlertsSortByShortHeaderAscending     GetAllAlertsSortByType = "short_header"
	GetAllAlertsSortByShortHeaderDescending    GetAllAlertsSortByType = "-short_header"
	GetAllAlertsSortByTimeframeAscending       GetAllAlertsSortByType = "timeframe"
	GetAllAlertsSortByTimeframeDescending      GetAllAlertsSortByType = "-timeframe"
	GetAllAlertsSortByUpdatedAtAscending       GetAllAlertsSortByType = "updated_at"
	GetAllAlertsSortByUpdatedAtDescending      GetAllAlertsSortByType = "-updated_at"
	GetAllAlertsSortByURLAscending             GetAllAlertsSortByType = "url"
	GetAllAlertsSortByURLDescending            GetAllAlertsSortByType = "-url"
)

// GetAllAlertsRequestConfig extra options for the GetAllAlerts request
type GetAllAlertsRequestConfig struct {
	PageOffset        string                 `url:"page[offset],omitempty"`             // Offset (0-based) of first element in the page// Offset (0-based) of first element in the page
	PageLimit         string                 `url:"page[limit],omitempty"`              // Max number of elements to return// Max number of elements to return
	Sort              GetAllAlertsSortByType `url:"sort,omitempty"`                     // Results can be sorted by the id or any GetAllRoutesSortByType
	Fields            []string               `url:"fields[alert],comma,omitempty"`      // Fields to include with the response. Note that fields can also be selected for included data types
	Include           []AlertInclude         `url:"include,comma,omitempty"`            // Include extra data in response
	FilterActivity    []AlertActivityType    `url:"filter[activity],comma,omitempty"`   // Filter to alerts for only those activities If the filter is not given OR it is empty, then defaults to ["BOARD", "EXIT", “RIDE”]. If the value AlertActivityFilterAll is used then all alerts will be returned, not just those with the default activities
	FilterRouteType   []RouteType            `url:"filter[route_type],comma,omitempty"` // Filter by route_type
	FilterDirectionID string                 `url:"filter[direction_id],omitempty"`     // Filter by direction of travel along the route
	FilterRouteIDs    []string               `url:"filter[route],omitempty"`            // Filter by route IDs
	FilterStopIDs     []string               `url:"filter[stop],omitempty"`             // Filter by stop IDs
	FilterTripIDs     []string               `url:"filter[trip],omitempty"`             // Filter by trip IDs
	FilterFacilityIDs []string               `url:"filter[facility],omitempty"`         // Filter by facility IDs
	FilterIDs         []string               `url:"filter[id],comma,omitempty"`         // Filter by multiple IDs
	FilterBanner      string                 `url:"filter[banner],comma,omitempty"`     // When combined with other filters, filters by alerts with or without a banner. MUST be “true” or "false"
	FilterDateTime    *TimeISO8601           `url:"filter[datetime],omitempty"`         // Filter to alerts that are active at a given time. Additionally, set `TimeISO8601.Now = true` to filter to alerts that are currently active.
	FilterLifecycle   []string               `url:"filter[lifecycle],comma,omitempty"`  // Filters by an alert’s lifecycle
	FilterSeverity    []string               `url:"filter[severity],comma,omitempty"`   // Filters alerts by list of severities
}

// GetAllAlerts returns all alerts from the mbta API
func (s *AlertService) GetAllAlerts(config *GetAllAlertsRequestConfig) ([]*Alert, *http.Response, error) {
	return s.GetAllAlertsWithContext(context.Background(), config)
}

// GetAllAlertsWithContext returns all alerts from the mbta API given a context
func (s *AlertService) GetAllAlertsWithContext(ctx context.Context, config *GetAllAlertsRequestConfig) ([]*Alert, *http.Response, error) {
	u, err := addOptions(alertsAPIPath, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	untypedAlerts, resp, err := s.client.doManyPayload(req, &Alert{})
	alerts := make([]*Alert, len(untypedAlerts))
	for i := 0; i < len(untypedAlerts); i++ {
		alerts[i] = untypedAlerts[i].(*Alert)
	}

	return alerts, resp, err
}

// GetAlertRequestConfig extra options for the GetAlert request
type GetAlertRequestConfig struct {
	Fields  []string       `url:"fields[alert],comma,omitempty"` // Fields to include with the response. Note that fields can also be selected for included data types
	Include []AlertInclude `url:"include,comma,omitempty"`       // Include extra data in response
}

// GetAlert return an alert from the mbta API
func (s *AlertService) GetAlert(id string, config *GetAlertRequestConfig) (*Alert, *http.Response, error) {
	return s.GetAlertWithContext(context.Background(), id, config)
}

// GetAlertWithContext return an alert from the mbta API given a context
func (s *AlertService) GetAlertWithContext(ctx context.Context, id string, config *GetAlertRequestConfig) (*Alert, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", alertsAPIPath, id)
	u, err := addOptions(path, config)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.newGETRequest(u)
	if err != nil {
		return nil, nil, err
	}
	req = req.WithContext(ctx)

	var alert Alert
	resp, err := s.client.doSinglePayload(req, &alert)
	return &alert, resp, err
}
