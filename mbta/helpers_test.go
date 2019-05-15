package mbta

import (
	"net/url"
	"testing"
)

func Test_addCommaSeparatedListToQuery(t *testing.T) {
	tests := []struct {
		key      string
		vals     []string
		expected string
	}{
		{"a", []string{"x", "y", "z"}, "x,y,z"},
		{"b", []string{}, ""},
	}

	query := make(url.Values)
	for _, test := range tests {
		addCommaSeparatedListToQuery(query, test.key, test.vals)
		equals(t, test.expected, query.Get(test.key))
		query.Del(test.key)
	}
}

func Test_addToQuery(t *testing.T) {
	tests := []struct {
		key      string
		val      string
		expected string
	}{
		{"a", "x", "x"},
		{"b", "", ""},
	}

	query := make(url.Values)
	for _, test := range tests {
		addToQuery(query, test.key, test.val)
		equals(t, test.expected, query.Get(test.key))
		query.Del(test.key)
	}
}
