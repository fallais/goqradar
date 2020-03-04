package goqradar

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Arule is a Qradar analytics rules
type Arule struct {
	AverageCapacity      int    `json:"average_capacity"`
	BaseCapacity         int    `json:"base_capacity"`
	BaseHostID           int    `json:"base_host_id"`
	CapacityTimestamp    int    `json:"capacity_timestamp"`
	CreationDate         int    `json:"creation_date"`
	Enabled              bool   `json:"enabled"`
	ID                   int    `json:"id"`
	Identifier           string `json:"identifier"`
	LinkedRuleIdentifier string `json:"linked_rule_identifier"`
	ModificationDate     int    `json:"modification_date"`
	Name                 string `json:"name"`
	Origin               string `json:"origin"`
	Owner                string `json:"owner"`
	Type                 string `json:"type"`
}

// RulesPaginatedResponse is the paginated response.
type RulesPaginatedResponse struct {
	Total int      `json:"total"`
	Min   int      `json:"min"`
	Max   int      `json:"max"`
	Rules []*Arule `json:"offenses"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListRules retrieves a list of rules
func (endpoint *Endpoint) ListRules(ctx context.Context, fields string, filter string, min, max int) (*RulesPaginatedResponse, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}
	if filter != "" {
		options = append(options, WithParam("filter", filter))
	}
	options = append(options, WithHeader("Range", fmt.Sprintf("items=%d-%d", min, max)))

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/analytics/rules", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Process the Content-Range
	min, max, total, err := parseContentRange(resp.Header.Get("Content-Range"))
	if err != nil {
		return nil, fmt.Errorf("error while parsing the content-range: %s", err)
	}

	// Prepare the response
	response := &RulesPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.Rules)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
