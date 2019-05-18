package mbta

import (
	"net/http"
	"net/url"
	"reflect"

	jsonapi "github.com/michele/go.jsonapi"
)

const (
	defaultBaseURL   = "https://api-v3.mbta.com"
	defaultUserAgent = "mbta-v3-go"
)

// ClientConfig the options for creating a Client
type ClientConfig struct {
	BaseURL   string
	APIKey    string
	UserAgent string
}

type service struct {
	client *Client
}

// Client the client for the MBTA API
type Client struct {
	client *http.Client

	APIKey string

	BaseURL   *url.URL
	UserAgent string

	common   service // Reuse a single struct instead of allocating one for each service on the heap. (same as github.com/google/go-github)
	Stops    *StopService
	Trips    *TripService
	Vehicles *VehicleService
}

// NewClient creates a new Client using the given config options
func NewClient(config ClientConfig) *Client {
	c := &Client{
		client:    http.DefaultClient,
		APIKey:    config.APIKey,
		UserAgent: config.UserAgent,
	}

	if config.BaseURL == "" {
		c.BaseURL, _ = url.Parse(defaultBaseURL)
	} else {
		parsedURL, err := url.Parse(config.BaseURL)
		if err != nil {
			panic(err)
		}
		c.BaseURL = parsedURL
	}

	if config.UserAgent == "" {
		c.UserAgent = defaultUserAgent
	}

	c.common.client = c
	c.Stops = (*StopService)(&c.common)
	c.Trips = (*TripService)(&c.common)
	c.Vehicles = (*VehicleService)(&c.common)

	return c
}

func (c *Client) newGETRequest(path string) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.APIKey != "" {
		req.Header.Set("x-api-key", c.APIKey)
	}
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) doSinglePayload(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = getSpecialError(resp, err); err != nil {
		return nil, err
	}

	err = jsonapi.UnmarshalPayload(resp.Body, v)
	return resp, err
}

func (c *Client) doManyPayload(req *http.Request, v interface{}) ([]interface{}, *http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	if err = getSpecialError(resp, err); err != nil {
		return nil, nil, err
	}

	vals, err := jsonapi.UnmarshalManyPayload(resp.Body, reflect.TypeOf(v))
	return vals, resp, err
}
