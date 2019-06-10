package mbta

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetAlert(t *testing.T) {
	expected := &Line{
		ID:        "line-Green",
		Color:     "00843D",
		LongName:  "Green Line",
		ShortName: "",
		SortOrder: 10032,
		TextColor: "FFFFFF",
		Routes:    []Route(nil),
	}
	server := httptest.NewServer(handlerForServer(t, fmt.Sprintf("%s/%s", linesAPIPath, "line-Green")))
	defer server.Close()
	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()
	actual, _, err := mbtaClient.Lines.GetLine("line-Green", &GetLineRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}

func TestGetAllAlerts(t *testing.T) {
	updatedAt1, _ := parseISO8601Time("2019-06-08T16:35:24-04:00")
	createdAt1, _ := parseISO8601Time("2019-06-08T16:35:24-04:00")
	activePeriodStart1, _ := parseISO8601Time("2019-06-08T16:35:24-04:00")

	url2, _ := url.Parse("https://www.mbta.com/bus-schedule-changes")
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
			URL:           url2,
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
	}
	server := httptest.NewServer(handlerForServer(t, alertsAPIPath))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Alerts.GetAllAlerts(&GetAllAlertsRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}
