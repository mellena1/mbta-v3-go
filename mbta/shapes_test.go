package mbta

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestGetShape(t *testing.T) {
	expected := &Shape{
		ID:          "660085",
		Priority:    3,
		Name:        "Dudley Station via Allston",
		DirectionID: 1,
		Stops: []*Stop{
			&Stop{ID: "22549"},
			&Stop{ID: "32549"},
		},
		Route: &Route{ID: "66"},
	}
	server := httptest.NewServer(handlerForServer(t, fmt.Sprintf("%s/%s", shapesAPIPath, "660085")))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Shapes.GetShape("660085", GetShapeRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}

func TestGetAllShapes(t *testing.T) {
	expected := []*Shape{
		&Shape{
			ID:          "660085",
			Priority:    3,
			Name:        "Dudley Station via Allston",
			DirectionID: 1,
			Stops: []*Stop{
				&Stop{ID: "22549"},
				&Stop{ID: "32549"},
			},
			Route: &Route{ID: "66"},
		},
		&Shape{
			ID:          "660113-2",
			Priority:    2,
			Name:        "Franklin Park via Dudley",
			DirectionID: 1,
			Stops: []*Stop{
				&Stop{ID: "925"},
				&Stop{ID: "926"},
			},
			Route: &Route{ID: "66"},
		},
	}
	server := httptest.NewServer(handlerForServer(t, shapesAPIPath))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Shapes.GetAllShapes(GetAllShapesRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}
