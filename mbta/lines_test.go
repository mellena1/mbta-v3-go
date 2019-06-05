package mbta

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestGetLine(t *testing.T) {
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

func TestGetAllLines(t *testing.T) {
	expected := []*Line{
		&Line{
			ID:        "line-Green",
			Color:     "00843D",
			LongName:  "Green Line",
			ShortName: "",
			SortOrder: 10032,
			TextColor: "FFFFFF",
			Routes:    []Route(nil),
		},
		&Line{
			ID:        "line-Mattapan",
			Color:     "DA291C",
			LongName:  "Mattapan Trolley",
			ShortName: "",
			SortOrder: 10011,
			TextColor: "FFFFFF",
			Routes:    []Route(nil),
		},
	}
	server := httptest.NewServer(handlerForServer(t, linesAPIPath))
	defer server.Close()
	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()
	actual, _, err := mbtaClient.Lines.GetAllLines(&GetAllLinesRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}
