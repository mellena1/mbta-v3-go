package mbta

import (
	"net/url"
	"testing"
)

func TestJSONURL_Unmarshal(t *testing.T) {
	googleURL, _ := url.Parse("http://www.google.com")
	tests := []struct {
		val      string
		expected *url.URL
		valid    bool
	}{
		{"http://www.google.com", googleURL, true},
		{"\"http://www.google.com\"", googleURL, true},
		{"invalid URL", &url.URL{}, false},
	}

	jURL := JSONURL{}
	for _, test := range tests {
		err := jURL.UnmarshalJSON([]byte(test.val))
		if test.valid {
			ok(t, err)
			equals(t, test.expected, jURL.URL)
		} else {
			assert(t, err != nil, "this should have errored")
		}
	}
}
