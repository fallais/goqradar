package goqradar

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

//------------------------------------------------------------------------------
// Structures
//------------------------------------------------------------------------------

// LoginAttempt is a login attempt
type LoginAttempt struct {
	AttemptResult string `json:"attempt_result"`
	AttemptTime   int    `json:"attempt_time"`
	RemoteIP      string `json:"remote_ip"`
	UserID        int    `json:"user_id"`
	AttemptMethod string `json:"attempt_method"`
}

// LoginAttemptPaginatedResponse is the paginated response.
type LoginAttemptPaginatedResponse struct {
	Total         int             `json:"total"`
	Min           int             `json:"min"`
	Max           int             `json:"max"`
	LoginAttempts []*LoginAttempt `json:"offense_types"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListAccessAttempts returns the list of the login attempts with given fields, filters and sort.
func (endpoint *Endpoint) ListAccessAttempts(ctx context.Context, fields, filter, sort string, min, max int) (*LoginAttemptPaginatedResponse, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}
	if filter != "" {
		options = append(options, WithParam("filter", filter))
	}
	if sort != "" {
		options = append(options, WithParam("sort", sort))
	}
	options = append(options, WithHeader("Range", fmt.Sprintf("items=%d-%d", min, max)))

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/access/login_attempts", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Process the Content-Range
	min, max, total, err := parseContentRange(resp.Header.Get("Content-Range"))
	if err != nil {
		return nil, fmt.Errorf("error while parsing the content-range [%s]: %s", resp.Header.Get("Content-Range"), err)
	}

	// Prepare the response
	response := &LoginAttemptPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.LoginAttempts)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
