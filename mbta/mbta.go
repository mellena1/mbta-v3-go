package mbta

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	jsonapi "github.com/michele/go.jsonapi"
)

const (
	defaultBaseURL = "https://api-v3.mbta.com"
)

// Config the options for creating a Client
type Config struct {
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
	Vehicles *VehicleService
}

// NewClient creates a new Client using the given config options
func NewClient(config Config) *Client {
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

	c.common.client = c
	c.Stops = (*StopService)(&c.common)
	c.Vehicles = (*VehicleService)(&c.common)

	return c
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.APIKey != "" {
		req.Header.Set("x-api-key", c.APIKey)
	}
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = jsonapi.UnmarshalPayload(resp.Body, v)
	return resp, err
}
