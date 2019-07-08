package mbta

import (
	"net/http/httptest"
	"testing"

	"golang.org/x/xerrors"
)

func Test_GetAllPredictions(t *testing.T) {
	pTime, _ := parseISO8601Time("2019-06-05T15:49:43-04:00")
	parsedDepartureTime1 := timeToTimeISO8601(pTime)
	pTime, _ = parseISO8601Time("2019-06-05T15:52:20-04:00")
	parsedDepartureTime2 := timeToTimeISO8601(pTime)

	expected := []*Prediction{
		&Prediction{
			ID:                   "prediction-39990839-20:30-NewtonHighlandsRiverside-70196-50",
			ArrivalTime:          nil,
			DepartureTime:        &parsedDepartureTime1,
			DirectionID:          0,
			ScheduleRelationship: nil,
			Status:               nil,
			StopSequence:         50,
			Route:                &Route{ID: "Green-B"},
			Stop:                 &Stop{ID: "70196"},
			Trip:                 &Trip{ID: "39990839-20:30-NewtonHighlandsRiverside"},
			Vehicle:              nil,
			Schedule:             nil,
		},
		&Prediction{
			ID:                   "prediction-39990840-20:30-NewtonHighlandsRiverside-70196-50",
			ArrivalTime:          nil,
			DepartureTime:        &parsedDepartureTime2,
			DirectionID:          0,
			ScheduleRelationship: nil,
			Status:               nil,
			StopSequence:         50,
			Route:                &Route{ID: "Green-B"},
			Stop:                 &Stop{ID: "70196"},
			Trip:                 &Trip{ID: "39990840-20:30-NewtonHighlandsRiverside"},
			Vehicle:              nil,
			Schedule:             nil,
		},
	}
	opts := &GetAllPredictionsRequestConfig{FilterRouteIDs: []string{"Green-B"}}
	fullPath, _ := addOptions(predictionsAPIPath, opts)
	server := httptest.NewServer(handlerForServer(t, fullPath))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Predictions.GetAllPredictions(opts)
	ok(t, err)
	equals(t, expected, actual)
}

func Test_GetAllPredictionsFail(t *testing.T) {
	server := httptest.NewServer(handlerForServer(t, predictionsAPIPath))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	_, _, err := mbtaClient.Predictions.GetAllPredictions(&GetAllPredictionsRequestConfig{})
	equals(t, true, xerrors.Is(err, ErrInvalidConfig))
}
