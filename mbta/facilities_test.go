package mbta

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestGetFacility(t *testing.T) {
	expected := &Facility{
		ID:        "park-NB-0127",
		Latitude:  float64Ptr(42.28123),
		Longitude: float64Ptr(-71.237271),
		Name:      "Needham Center Parking Lot",
		ShortName: "Parking Lot",
		Type:      FacilityParkingArea,
		Properties: []FacilityProperty{
			FacilityProperty{
				Name:  "attended",
				Value: "2",
			},
			FacilityProperty{
				Name:  "capacity",
				Value: "35",
			},
			FacilityProperty{
				Name:  "contact",
				Value: "Town of Needham, Parking Clerk",
			},
			FacilityProperty{
				Name:  "contact-phone",
				Value: "781-455-7500",
			},
		},
		Stop: &Stop{
			ID: "place-NB-0127",
		},
	}
	server := httptest.NewServer(handlerForServer(t, fmt.Sprintf("%s/%s", facilitiesAPIPath, "park-NB-0127")))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Facilities.GetFacility("park-NB-0127", &GetFacilityRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}

func TestGetAllFacilities(t *testing.T) {
	expected := []*Facility{
		&Facility{
			ID:        "986",
			Latitude:  nil,
			Longitude: nil,
			Name:      "Porter Elevator 986 (Commuter Rail platform to lobby)",
			ShortName: "Commuter Rail platform to lobby",
			Type:      FacilityElevator,
			Properties: []FacilityProperty{
				FacilityProperty{
					Name:  "alternate-service-text",
					Value: "See station personnel or use the call box to request assistance.",
				},
				FacilityProperty{
					Name:  "excludes-stop",
					Value: "23151",
				},
			},
			Stop: &Stop{
				ID: "place-portr",
			},
		},
		&Facility{
			ID:        "retailsale-302029",
			Latitude:  float64Ptr(42.3415202),
			Longitude: float64Ptr(-71.0868242),
			Name:      "Symphony Market",
			ShortName: "Symphony Market",
			Type:      FacilityFareVendingRetailer,
			Properties: []FacilityProperty{
				FacilityProperty{
					Name:  "address",
					Value: "291 Huntington Ave, Boston, MA 02115",
				},
				FacilityProperty{
					Name:  "enclosed",
					Value: "1",
				},
			},
			Stop: nil,
		},
	}
	server := httptest.NewServer(handlerForServer(t, facilitiesAPIPath))
	defer server.Close()

	mbtaClient := NewClient(ClientConfig{BaseURL: server.URL})
	mbtaClient.client = server.Client()

	actual, _, err := mbtaClient.Facilities.GetAllFacilities(&GetAllFacilitiesRequestConfig{})
	ok(t, err)
	equals(t, expected, actual)
}
