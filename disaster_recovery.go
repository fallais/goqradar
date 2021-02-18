package goqradar

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

//------------------------------------------------------------------------------
// Structures
//------------------------------------------------------------------------------

// ArielCopyProfile is an ariel copy profile
type ArielCopyProfile struct {
	BandwidthLimit                        int    `json:"bandwidth_limit"`
	DestinationHostIP                     string `json:"destination_host_ip"`
	DestinationPort                       int    `json:"destination_port"`
	Enabled                               bool   `json:"enabled"`
	EndDate                               int    `json:"end_date"`
	ExcludeEventRetentionBucketIds        []int  `json:"exclude_event_retention_bucket_ids"`
	ExcludeFlowRetentionBucketIds         []int  `json:"exclude_flow_retention_bucket_ids"`
	Frequency                             int    `json:"frequency"`
	HostID                                int    `json:"host_id"`
	ID                                    int    `json:"id"`
	LastErrorArielContentTransferred      int    `json:"last_error_ariel_content_transferred"`
	LastErrorArielType                    string `json:"last_error_ariel_type"`
	LastErrorDate                         int    `json:"last_error_date"`
	LastErrorMaxThresholdSurpassedDate    int    `json:"last_error_max_threshold_surpassed_date"`
	LastSuccessfulArielContentTransferred int    `json:"last_successful_ariel_content_transferred"`
	LastSuccessfulTransferArielType       string `json:"last_successful_transfer_ariel_type"`
	LastSuccessfulTransferDate            int    `json:"last_successful_transfer_date"`
	StartDate                             int    `json:"start_date"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListArielCopyProfiles returns a list of the ariel copy profile with given fields, filters and sort.
func (endpoint *Endpoint) ListArielCopyProfiles(ctx context.Context, fields, filter string) ([]*ArielCopyProfile, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}
	if filter != "" {
		options = append(options, WithParam("filter", filter))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/disaster_recovery/ariel_copy_profiles", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response []*ArielCopyProfile

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// CreateArielCopyProfille creates a ariel copy profile.
func (endpoint *Endpoint) CreateArielCopyProfille(ctx context.Context, data map[string]interface{}, fields string) error {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/disaster_recovery/ariel_copy_profiles"

	// Create the data
	d, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Error while marshalling the values : %s", err)
	}

	// Create the request
	req, err := http.NewRequest("POST", reqURL.String(), bytes.NewBuffer(d))
	if err != nil {
		return fmt.Errorf("Error while creating the request : %s", err)
	}

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/json")
	if fields != "" {
		req.Header.Set("fields", fields)
	}

	// Do the request
	_, err = endpoint.client.client.Do(req)
	if err != nil {
		return fmt.Errorf("Error while doing the request : %s", err)
	}

	return nil
}

// GetArielCopyProfile retrieves a Ariel copy profile.
func (endpoint *Endpoint) GetArielCopyProfile(ctx context.Context, id int, fields string) (*ArielCopyProfile, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/disaster_recovery/ariel_copy_profiles/"+strconv.Itoa(id), options...)
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
	var response *ArielCopyProfile

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// UpdateArielCopyProfile by id
func (endpoint *Endpoint) UpdateArielCopyProfile(ctx context.Context, id int, data map[string]interface{}, fields string) (*ArielCopyProfile, error) {

	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/disaster_recovery/ariel_copy_profiles/"
	reqURL.Path += strconv.Itoa(id)

	// Create the data
	d, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Error while marshalling the values : %s", err)
	}

	// Create the request
	req, err := http.NewRequest("POST", reqURL.String(), bytes.NewBuffer(d))
	if err != nil {
		return nil, fmt.Errorf("Error while creating the request : %s", err)
	}

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/json")
	if fields != "" {
		req.Header.Set("fields", fields)
	}

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while doing the request : %s", err)
	}
	defer resp.Body.Close()

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *ArielCopyProfile

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// DeleteArielCopyProfile by ID
func (endpoint *Endpoint) DeleteArielCopyProfile(ctx context.Context, id int) error {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/disaster_recovery/ariel_copy_profiles/"
	reqURL.Path += strconv.Itoa(id)

	// Create the request
	req, err := http.NewRequest("DELETE", reqURL.String(), nil)
	if err != nil {
		return fmt.Errorf("Error while creating the request : %s", err)
	}

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/json")

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return fmt.Errorf("Error while doing the request : %s", err)
	}

	// Check the status code
	if resp.StatusCode != 204 {
		return fmt.Errorf("Status code is %d : Error while reading the body", resp.StatusCode)
	}

	return nil

}
