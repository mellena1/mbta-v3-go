package mbta

import (
	"net/url"
	"strings"
)

// JSONURL wraps a url.URL to be parsed from JSON
type JSONURL struct {
	URL *url.URL
}

// UnmarshalJSON unmarshals a JSON string URL into a url.URL
func (j *JSONURL) UnmarshalJSON(b []byte) error {
	strURL := strings.Trim(string(b), "\"")
	url, err := url.ParseRequestURI(strURL)
	if err == nil {
		j.URL = url
	}
	return err
}
