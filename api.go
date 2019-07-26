package qradar

import (
	"fmt"
	"net/http"
	"net/url"
)

const BaseURI = "/api/"

// RequestOptions ...
type RequestOptions struct {
	Params  *url.Values
	Data    map[string]interface{}
	Headers http.Header
}

// ListOffenses returns the offenses.
func (c *qradar) ListOffenses() ([]*Offense, error) {
	resp, err := c.call(http.MethodGet, "siem/offenses")

	return nil, nil
}

// GetOffense returns the offense with given id.
func (c *qradar) GetOffense(id string, options RequestOptions) (*http.Response, error) {
	resp, err := c.call(http.MethodGet, fmt.Sprintf("siem/offenses/%s", id), options)

	return nil, nil
}

// UpdateOffense ...
func (c *qradar) UpdateOffense(id string, options RequestOptions) (*http.Response, error) {
	resp, err := c.call(http.MethodPost, fmt.Sprintf("siem/offenses/%s", id), options)

	return nil, nil
}
