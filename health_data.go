package goqradar

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//------------------------------------------------------------------------------
// Structures
//------------------------------------------------------------------------------

// SecurityDataCount is an array of security data object
type SecurityDataCount struct {
	Assets          int `json:"assets"`
	LogSources      int `json:"log_sources"`
	Offenses        int `json:"offenses"`
	Rules           int `json:"rules"`
	Vulnerabilities int `json:"vulnerabilities"`
}

//TopOffense is QRadar offense
type TopOffense struct {
	Count       int    `json:"count"`
	OffenseID   int    `json:"offense_id"`
	OffenseName string `json:"offense_name"`
}

// TopOffensesPaginatedResponse is the paginated response.
type TopOffensesPaginatedResponse struct {
	Total       int           `json:"total"`
	Min         int           `json:"min"`
	Max         int           `json:"max"`
	TopOffenses []*TopOffense `json:"offense_types"`
}

// TopRule is a QRadar top rules
type TopRule struct {
	Count    int    `json:"count"`
	RuleID   int    `json:"rule_id"`
	RuleName string `json:"rule_name"`
}

// TopRulesPaginatedResponse is the paginated response.
type TopRulesPaginatedResponse struct {
	Total    int        `json:"total"`
	Min      int        `json:"min"`
	Max      int        `json:"max"`
	TopRules []*TopRule `json:"offense_types"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// GetSecurityDataCount retrieves the count of security data object .
func (endpoint *Endpoint) GetSecurityDataCount(ctx context.Context, fields string) (*SecurityDataCount, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/health_data/security_data_count", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *SecurityDataCount

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// ListTopOffenses returns the Top Offenses in the system sorted by update count with given fields and filters.
func (endpoint *Endpoint) ListTopOffenses(ctx context.Context, fields, filter string, min, max int) (*TopOffensesPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/health_data/top_offenses", options...)
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
	response := &TopOffensesPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.TopOffenses)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// ListTopRules returns the Top Rules in the system sorted by response count with given fields and filters.
func (endpoint *Endpoint) ListTopRules(ctx context.Context, fields, filter string, min, max int) (*TopRulesPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/health_data/top_rules", options...)
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
	response := &TopRulesPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.TopRules)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
