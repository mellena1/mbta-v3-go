package mbta

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetAlert(t *testing.T) {
	url, _ := url.Parse("https://www.mbta.com/bus-schedule-changes")
	updatedAt, _ := parseISO8601Time("2019-06-06T16:59:04-04:00")
	createdAt, _ := parseISO8601Time("2019-05-28T22:13:37-04:00")
	activePeriodStart, _ := parseISO8601Time("2019-06-23T04:30:00-04:00")
	activePeriodEndTime, _ := parseISO8601Time("2019-07-08T02:30:00-04:00")
	activePeriodEnd := timeToTimeISO8601(activePeriodEndTime)
	rTBus := RouteTypeBus
	expected := &Alert{
		ID:            "313120",
		URL:           &JSONURL{URL: url},
		UpdatedAt:     timeToTimeISO8601(updatedAt),
		Timeframe:     strPtr("starting June 23"),
		ShortHeader:   "Beginning Sunday, June 23, the Route 43 summer bus schedule will take effect with Saturday and Sunday schedule changes throughout the day.",
		Severity:      3,
		ServiceEffect: "Route 43 schedule change",
		Lifecycle:     AlertLifecycleUpcoming,
		InformedEntity: []AlertInformedEntity{
			AlertInformedEntity{
				Activities: []AlertActivityType{
					AlertActivityBoard,
					AlertActivityExit,
					AlertActivityRide,
				},
				RouteID:   strPtr("43"),
				RouteType: &rTBus,
			},
		},
		Header:      "Beginning Sunday, June 23, the Route 43 summer bus schedule will take effect with Saturday and Sunday schedule changes throughout the day.",
		Effect:      AlertEffectScheduleChange,
		Description: strPtr("Please find your route at mbta.com/bus-schedule-changes and check the June 23 schedule for specific changes. You can also ask your bus driver for a printed summer schedule."),
		CreatedAt:   timeToTimeISO8601(createdAt),
		Cause:       AlertCauseUnknownCause,
		Banner:      nil,
		ActivePeriod: []AlertActivePeriod{
			AlertActivePeriod{
				Start: timeToTimeISO8601(activePeriodStart),
				End:   &activePeriodEnd,
			},
		},
	}
	server := httptest.NewServer(handlerForServer(t, fmt.Sprintf("%s/%s", alertsAPIPath, "313120")))
	defer server.Close()
	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()
	actual, _, err := mbtaClient.Alerts.GetAlert("313120", &GetAlertRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}

func TestGetAllAlerts(t *testing.T) {
	updatedAt1, _ := parseISO8601Time("2019-06-08T16:35:24-04:00")
	createdAt1, _ := parseISO8601Time("2019-06-08T16:35:24-04:00")
	activePeriodStart1, _ := parseISO8601Time("2019-06-08T16:35:24-04:00")

	url2, _ := url.Parse("https://www.mbta.com/bus-schedule-changes")
	updatedAt2, _ := parseISO8601Time("2019-06-06T16:48:11-04:00")
	createdAt2, _ := parseISO8601Time("2019-05-28T22:28:16-04:00")
	activePeriodStart2, _ := parseISO8601Time("2019-06-23T04:30:00-04:00")
	activePeriodEnd2Time, _ := parseISO8601Time("2019-07-08T02:30:00-04:00")
	activePeriodEnd2 := timeToTimeISO8601(activePeriodEnd2Time)
	rTBus := RouteTypeBus
	expected := []*Alert{
		&Alert{
			ID:            "315463",
			URL:           nil,
			UpdatedAt:     timeToTimeISO8601(updatedAt1),
			Timeframe:     nil,
			ShortHeader:   "Porter Escalator 511 (Ashmont/Braintree platform to paid lobby) unavailable due to maintenance",
			Severity:      3,
			ServiceEffect: "Porter escalator unavailable",
			Lifecycle:     AlertLifecycleNew,
			InformedEntity: []AlertInformedEntity{
				AlertInformedEntity{
					Activities: []AlertActivityType{
						AlertActivityUsingEscalator,
					},
					FacilityID: strPtr("511"),
					StopID:     strPtr("70065"),
				},
				AlertInformedEntity{
					Activities: []AlertActivityType{
						AlertActivityUsingEscalator,
					},
					FacilityID: strPtr("511"),
					StopID:     strPtr("place-portr"),
				},
				AlertInformedEntity{
					Activities: []AlertActivityType{
						AlertActivityUsingEscalator,
					},
					FacilityID: strPtr("511"),
					StopID:     strPtr("70066"),
				},
			},
			Header:      "Porter Escalator 511 (Ashmont/Braintree platform to paid lobby) unavailable due to maintenance",
			Effect:      AlertEffectEscalatorClosure,
			Description: nil,
			CreatedAt:   timeToTimeISO8601(createdAt1),
			Cause:       AlertCauseMaintenance,
			Banner:      nil,
			ActivePeriod: []AlertActivePeriod{
				AlertActivePeriod{
					Start: timeToTimeISO8601(activePeriodStart1),
					End:   nil,
				},
			},
		},
		&Alert{
			ID:            "313136",
			URL:           &JSONURL{URL: url2},
			UpdatedAt:     timeToTimeISO8601(updatedAt2),
			Timeframe:     strPtr("starting June 23"),
			ShortHeader:   "Beginning Sun, Jun 23, the Route 78 summer bus schedule will take effect with weekday, Sat, and Sun schedule changes throughout the day",
			Severity:      3,
			ServiceEffect: "Route 78 schedule change",
			Lifecycle:     AlertLifecycleUpcoming,
			InformedEntity: []AlertInformedEntity{
				AlertInformedEntity{
					Activities: []AlertActivityType{
						AlertActivityBoard,
						AlertActivityExit,
						AlertActivityRide,
					},
					RouteID:   strPtr("78"),
					RouteType: &rTBus,
				},
			},
			Header:      "Beginning Sunday, June 23, the Route 78 summer bus schedule will take effect with weekday, Saturday, and Sunday schedule changes throughout the day.",
			Effect:      AlertEffectScheduleChange,
			Description: strPtr("Weekday, Saturday, and Sunday schedule changes throughout the day. Due to Harvard Busway Renovation, inbound service will operate via Brattle Street.\r\n\r\nPlease find your route at mbta.com/bus-schedule-changes and check the June 23 schedule for specific changes. You can also ask your bus driver for a printed summer schedule."),
			CreatedAt:   timeToTimeISO8601(createdAt2),
			Cause:       AlertCauseUnknownCause,
			Banner:      nil,
			ActivePeriod: []AlertActivePeriod{
				AlertActivePeriod{
					Start: timeToTimeISO8601(activePeriodStart2),
					End:   &activePeriodEnd2,
				},
			},
		},
	}
	server := httptest.NewServer(handlerForServer(t, alertsAPIPath))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Alerts.GetAllAlerts(&GetAllAlertsRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}
