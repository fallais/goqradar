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

// Asset is a QRadar asset
type Asset struct {
	VulnerabilityCount int          `json:"vulnerability_count"`
	Interfaces         []Interfaces `json:"interfaces"`
	RiskScoreSum       float64      `json:"risk_score_sum"`
	Hostnames          []Hostnames  `json:"hostnames"`
	ID                 int          `json:"id"`
	DomainID           int          `json:"domain_id"`
	Properties         []Properties `json:"properties"`
	Users              []Users      `json:"users"`
	Products           []Products   `json:"products"`
}

// IPAddresses is a list of QRadar Ip addresse
type IPAddresses struct {
	LastSeenProfiler  int    `json:"last_seen_profiler"`
	Created           int    `json:"created"`
	FirstSeenScanner  int    `json:"first_seen_scanner"`
	LastSeenScanner   int    `json:"last_seen_scanner"`
	NetworkID         int    `json:"network_id"`
	ID                int    `json:"id"`
	Type              string `json:"type"`
	FirstSeenProfiler int    `json:"first_seen_profiler"`
	Value             string `json:"value"`
}

// Interfaces is a list of QRadar interface
type Interfaces struct {
	MacAddress        string        `json:"mac_address"`
	LastSeenProfiler  int           `json:"last_seen_profiler"`
	Created           int           `json:"created"`
	FirstSeenScanner  int           `json:"first_seen_scanner"`
	LastSeenScanner   int           `json:"last_seen_scanner"`
	IPAddresses       []IPAddresses `json:"ip_addresses"`
	ID                int           `json:"id"`
	FirstSeenProfiler int           `json:"first_seen_profiler"`
}

// Hostnames is a list of QRadar hostname
type Hostnames struct {
	LastSeenProfiler  int    `json:"last_seen_profiler"`
	Created           int    `json:"created"`
	Name              string `json:"name"`
	FirstSeenScanner  int    `json:"first_seen_scanner"`
	LastSeenScanner   int    `json:"last_seen_scanner"`
	ID                int    `json:"id"`
	Type              string `json:"type"`
	FirstSeenProfiler int    `json:"first_seen_profiler"`
}

// Properties is a list of QRadar propertie
type Properties struct {
	LastReported   int    `json:"last_reported"`
	Name           string `json:"name"`
	TypeID         int    `json:"type_id"`
	ID             int    `json:"id"`
	LastReportedBy string `json:"last_reported_by"`
	Value          string `json:"value"`
}

// Users is a list of QRadar user
type Users struct {
	LastSeenProfiler  int    `json:"last_seen_profiler"`
	FirstSeenScanner  int    `json:"first_seen_scanner"`
	LastSeenScanner   int    `json:"last_seen_scanner"`
	ID                int    `json:"id"`
	FirstSeenProfiler int    `json:"first_seen_profiler"`
	Username          string `json:"username"`
}

// Products is a list of QRadar product
type Products struct {
	LastScannedFor    int `json:"last_scanned_for"`
	LastSeenProfiler  int `json:"last_seen_profiler"`
	ProductVariantID  int `json:"product_variant_id"`
	FirstSeenScanner  int `json:"first_seen_scanner"`
	LastSeenScanner   int `json:"last_seen_scanner"`
	ID                int `json:"id"`
	FirstSeenProfiler int `json:"first_seen_profiler"`
}

// AssetsPaginatedResponse is the paginated response.
type AssetsPaginatedResponse struct {
	Total  int      `json:"total"`
	Min    int      `json:"min"`
	Max    int      `json:"max"`
	Assets []*Asset `json:"offense_types"`
}

// AssetPropertie is an asset propertie
type AssetPropertie struct {
	Custom   bool   `json:"custom"`
	DataType string `json:"data_type"`
	Display  bool   `json:"display"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	State    int    `json:"state"`
}

// AssetPropertiePaginatedResponse is the paginated response.
type AssetPropertiePaginatedResponse struct {
	Total           int               `json:"total"`
	Min             int               `json:"min"`
	Max             int               `json:"max"`
	AssetProperties []*AssetPropertie `json:"offense_types"`
}

// AssetSavedSearchGroups is an asset saved search groups
type AssetSavedSearchGroups struct {
	ChildGroups  []int    `json:"child_groups"`
	ChildItems   []string `json:"child_items"`
	Description  string   `json:"description"`
	ID           int      `json:"id"`
	Level        int      `json:"level"`
	ModifiedTime int      `json:"modified_time"`
	Name         string   `json:"name"`
	Owner        string   `json:"owner"`
	ParentID     int      `json:"parent_id"`
	Type         string   `json:"type"`
}

// AssetSavedSearchGroupPaginatedResponse is the paginated response.
type AssetSavedSearchGroupPaginatedResponse struct {
	Total                   int                       `json:"total"`
	Min                     int                       `json:"min"`
	Max                     int                       `json:"max"`
	AssetsSavedSearchGroups []*AssetSavedSearchGroups `json:"offense_types"`
}

// SavedSearche is a saved searches
type SavedSearche struct {
	Columns     []Column `json:"columns"`
	Description string   `json:"description"`
	Filters     []Filter `json:"filters"`
	ID          int      `json:"id"`
	IsShared    bool     `json:"is_shared"`
	Name        string   `json:"name"`
	Owner       string   `json:"owner"`
}

// Column is a column
type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Filter is a filter
type Filter struct {
	Operator  string `json:"operator"`
	Parameter string `json:"parameter"`
	Value     string `json:"value"`
}

// SavedSearchesPaginatedResponse is the paginated response.
type SavedSearchesPaginatedResponse struct {
	Total         int             `json:"total"`
	Min           int             `json:"min"`
	Max           int             `json:"max"`
	SavedSearches []*SavedSearche `json:"offense_types"`
}

// AssetBasedOnSavedSearch is an asset based on  the result of an asset saved search
type AssetBasedOnSavedSearch struct {
	DomainID   int          `json:"domain_id"`
	ID         int          `json:"id"`
	Interfaces []Interfaces `json:"interfaces"`
	Properties []Properties `json:"properties"`
}

// AssetBasedOnSavedSearchPaginatedResponse is the paginated response.
type AssetBasedOnSavedSearchPaginatedResponse struct {
	Total                    int                        `json:"total"`
	Min                      int                        `json:"min"`
	Max                      int                        `json:"max"`
	AssetsBasedOnSavedSearch []*AssetBasedOnSavedSearch `json:"offense_types"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListAssets returns the assets with given fields, filters and sort.
func (endpoint *Endpoint) ListAssets(ctx context.Context, fields, filter, sort string, min, max int) (*AssetsPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "asset_model/assets", options...)
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
	response := &AssetsPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.Assets)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// UpdateAsset by name
func (endpoint *Endpoint) UpdateAsset(ctx context.Context, name string, data map[string]map[string]string) (string, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return "", fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/asset_model/assets/"
	reqURL.Path += name

	// Create the data
	d, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("Error while marshalling the values : %s", err)
	}

	// Create the request
	req, err := http.NewRequest("POST", reqURL.String(), bytes.NewBuffer(d))
	if err != nil {
		return "", fmt.Errorf("Error while creating the request : %s", err)
	}

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/json")

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error while doing the request : %s", err)
	}
	defer resp.Body.Close()

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response string

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return "", fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// ListAssetProperties returns the assets properties.
func (endpoint *Endpoint) ListAssetProperties(ctx context.Context, fields, filter string, min, max int) (*AssetPropertiePaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/asset_model/properties", options...)
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
	response := &AssetPropertiePaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.AssetProperties)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// ListAssetsSavedSearchGroups returns the assets saved searh groups.
func (endpoint *Endpoint) ListAssetsSavedSearchGroups(ctx context.Context, fields, filter string, min, max int) (*AssetSavedSearchGroupPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/asset_model/saved_search_groups", options...)
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
	response := &AssetSavedSearchGroupPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.AssetsSavedSearchGroups)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// GetAssetSavedSearchGroups retrieves an asset saved search group.
func (endpoint *Endpoint) GetAssetSavedSearchGroups(ctx context.Context, id int, fields string) (*AssetSavedSearchGroups, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/asset_model/saved_search_groups/"+strconv.Itoa(id), options...)
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
	var response *AssetSavedSearchGroups

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// UpdateAssetSavedSeachGroup update the owner of the asset saved search group
func (endpoint *Endpoint) UpdateAssetSavedSeachGroup(ctx context.Context, name int, fields string, data map[string]map[string]string) (*AssetSavedSearchGroups, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/asset_model/saved_search_groups/"
	reqURL.Path += strconv.Itoa(name)

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
	req.Header.Set("fields", fields)

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
	var response *AssetSavedSearchGroups

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// DeleteAssetSavedSearchGroups by groupID
func (endpoint *Endpoint) DeleteAssetSavedSearchGroups(ctx context.Context, name int) error {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/asset_model/saved_search_groups/"
	reqURL.Path += strconv.Itoa(name)

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
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != 204 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Status code is %d : Error while reading the body", resp.StatusCode)
		}

		return fmt.Errorf("Status code is %d : %s", resp.StatusCode, string(body))
	}

	return nil
}

// ListSavedSearches returns a list of saved searches, filters and sort.
func (endpoint *Endpoint) ListSavedSearches(ctx context.Context, fields, filter string, min, max int) (*SavedSearchesPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/asset_model/saved_searches", options...)
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
	response := &SavedSearchesPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.SavedSearches)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// GetAssetSavedSearch retrieves an asset saved search .
func (endpoint *Endpoint) GetAssetSavedSearch(ctx context.Context, id int, fields string) (*SavedSearche, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/asset_model/saved_searches/"+strconv.Itoa(id), options...)
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
	var response *SavedSearche

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// UpdateAssetSavedSearch by name
func (endpoint *Endpoint) UpdateAssetSavedSearch(ctx context.Context, name int, data map[string]map[string]string, fields string) (*SavedSearche, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/asset_model/saved_searches/"
	reqURL.Path += strconv.Itoa(name)

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
	req.Header.Set("fields", fields)

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
	var response *SavedSearche

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// DeleteAssetSavedSearch by groupID
func (endpoint *Endpoint) DeleteAssetSavedSearch(ctx context.Context, name int) error {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/asset_model/saved_searches/"
	reqURL.Path += strconv.Itoa(name)

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
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != 204 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Status code is %d : Error while reading the body", resp.StatusCode)
		}

		return fmt.Errorf("Status code is %d : %s", resp.StatusCode, string(body))
	}

	return nil
}

// ListAssetSavedSearches returns a list of saved searches, filters and sort.
func (endpoint *Endpoint) ListAssetSavedSearches(ctx context.Context, name, fields, filter string, min, max int) (*AssetBasedOnSavedSearchPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/asset_model/saved_searches/"+name+"/results", options...)
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
	response := &AssetBasedOnSavedSearchPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.AssetsBasedOnSavedSearch)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
