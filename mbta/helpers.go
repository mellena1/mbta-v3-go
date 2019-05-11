package mbta

import (
	"net/url"
	"strings"
)

func addCommaSeparatedListToQuery(query url.Values, key string, l []string) {
	if len(l) > 0 {
		query.Add(key, strings.Join(l, ","))
	}
}

func addToQuery(query url.Values, key string, val string) {
	if val != "" {
		query.Add(key, val)
	}
}
