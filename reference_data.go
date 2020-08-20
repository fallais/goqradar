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

// Map is a QRadar map
type Map struct {
	CreationTime     int    `json:"creation_time"`
	ElementType      string `json:"element_type"`
	KeyLabel         string `json:"key_label"`
	Name             string `json:"name"`
	NumberOfElements int    `json:"number_of_elements"`
	TimeToLive       string `json:"time_to_live"`
	TimeoutType      string `json:"timeout_type"`
	ValueLabel       string `json:"value_label"`
}

// BulkMap is a QRadar bulkmap
type BulkMap struct {
	CreationTime     int    `json:"creation_time"`
	ElementType      string `json:"element_type"`
	Name             string `json:"name"`
	NumberOfElements int    `json:"number_of_elements"`
	TimeToLive       string `json:"time_to_live"`
	TimeoutType      string `json:"timeout_type"`
}

// Set is a QRadar set
type Set struct {
	CreationTime     int    `json:"creation_time"`
	ElementType      string `json:"element_type"`
	Name             string `json:"name"`
	NumberOfElements int    `json:"number_of_elements"`
	TimeToLive       string `json:"time_to_live"`
	TimeoutType      string `json:"timeout_type"`
}

// Table is a QRadar table
type Table struct {
	CreationTime     int          `json:"creation_time"`
	ElementType      string       `json:"element_type"`
	KeyLabel         string       `json:"key_label"`
	KeyNameTypes     KeyNameTypes `json:"key_name_types"`
	Name             string       `json:"name"`
	NumberOfElements int          `json:"number_of_elements"`
	TimeToLive       string       `json:"time_to_live"`
	TimeoutType      string       `json:"timeout_type"`
}

// KeyNameTypes is a keyname...
type KeyNameTypes struct {
	String string `json:"String"`
}

// BulkTable is a QRadar bulktable
type BulkTable struct {
	CreationTime     int    `json:"creation_time"`
	ElementType      string `json:"element_type"`
	Name             string `json:"name"`
	NumberOfElements int    `json:"number_of_elements"`
	TimeToLive       string `json:"time_to_live"`
	TimeoutType      string `json:"timeout_type"`
}

// ListSetsPaginatedResponse is the paginated response.
type ListSetsPaginatedResponse struct {
	Total    int    `json:"total"`
	Min      int    `json:"min"`
	Max      int    `json:"max"`
	ListSets []*Set `json:"offenses"`
}

// BulkMapOfMap is BulkMapOfMap
type BulkMapOfMap struct {
	CreationTime     int    `json:"creation_time"`
	ElementType      string `json:"element_type"`
	Name             string `json:"name"`
	NumberOfElements int    `json:"number_of_elements"`
	TimeToLive       string `json:"time_to_live"`
	TimeoutType      string `json:"timeout_type"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// GetMap by name
func (endpoint *Endpoint) GetMap(ctx context.Context, fields string, filter string, Range string) (*Set, error) {
	return nil, nil
}

// ListMaps list all the sets
func (endpoint *Endpoint) ListMaps(ctx context.Context, fields string, filter string, Range string) (*[]Set, error) {
	return nil, nil
}

// UpdateBulkLoadRM by name
func (endpoint *Endpoint) UpdateBulkLoadRM(ctx context.Context, name string, data map[string]string, fields string) (*BulkMap, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/api/reference_data/maps/bulk_load/"
	reqURL.Path += name

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
	var mape *BulkMap

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &mape)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return mape, nil
}

// DeleteReferenceMap by name
func (endpoint *Endpoint) DeleteReferenceMap(ctx context.Context, name, fields, namespace string, purgeOnly bool) error {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/api/reference_data/maps/"
	reqURL.Path += name
	parameters := url.Values{}
	parameters.Add("purge_only", strconv.FormatBool(purgeOnly))
	parameters.Add("fields", fields)
	parameters.Add("namespace", namespace)
	reqURL.RawQuery = parameters.Encode()

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
	if resp.StatusCode != 202 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Status code is %d : Error while reading the body", resp.StatusCode)
		}

		return fmt.Errorf("Status code is %d : %s", resp.StatusCode, string(body))
	}

	return nil
}

// GetSet by name
func (endpoint *Endpoint) GetSet(ctx context.Context, fields string, filter string, Range string) (*Set, error) {
	return nil, nil
}

// ListSets list all the sets
func (endpoint *Endpoint) ListSets(ctx context.Context, fields string, filter string, min, max int) (*ListSetsPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/reference_data/sets", options...)
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
	response := &ListSetsPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.ListSets)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// UpdateBulkLoadRS by name
func (endpoint *Endpoint) UpdateBulkLoadRS(ctx context.Context, name string, data []string, fields string) (*Set, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/api/reference_data/sets/bulk_load/"
	reqURL.Path += name

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
	var set *Set

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &set)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return set, nil
}

// DeleteReferenceSet removes a reference set or purges its contents
func (endpoint *Endpoint) DeleteReferenceSet(ctx context.Context, name, fields, namespace string, purgeOnly bool) error {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/api/reference_data/sets/"
	reqURL.Path += name
	parameters := url.Values{}
	parameters.Add("purge_only", strconv.FormatBool(purgeOnly))
	parameters.Add("fields", fields)
	parameters.Add("namespace", namespace)
	reqURL.RawQuery = parameters.Encode()

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
	if resp.StatusCode != 202 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Status code is %d : Error while reading the body", resp.StatusCode)
		}

		return fmt.Errorf("Status code is %d : %s", resp.StatusCode, string(body))
	}

	return nil
}

// GetTable by name
func (endpoint *Endpoint) GetTable(ctx context.Context, fields string, filter string, Range string) (*Table, error) {
	return nil, nil
}

// ListTables list all the sets
func (endpoint *Endpoint) ListTables(ctx context.Context, fields string, filter string, Range string) (*[]Table, error) {
	return nil, nil
}

// UpdateBulkLoadRT by name
func (endpoint *Endpoint) UpdateBulkLoadRT(ctx context.Context, name, fields string, data map[string]map[string]string) (*BulkTable, error) {
	// Prepare the URLg
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/api/reference_data/tables/bulk_load/"
	reqURL.Path += name

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
	var table *BulkTable

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &table)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return table, nil
}

// DeleteReferenceTable by name
func (endpoint *Endpoint) DeleteReferenceTable(ctx context.Context, name, fields, namespace string, purgeOnly bool) error {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/api/reference_data/tables/"
	reqURL.Path += name
	parameters := url.Values{}
	parameters.Add("purge_only", strconv.FormatBool(purgeOnly))
	parameters.Add("fields", fields)
	parameters.Add("namespace", namespace)
	reqURL.RawQuery = parameters.Encode()

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
	if resp.StatusCode != 202 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Status code is %d : Error while reading the body", resp.StatusCode)
		}

		return fmt.Errorf("Status code is %d : %s", resp.StatusCode, string(body))
	}

	return nil
}

// UpdateBulkLoadRMM by name
func (endpoint *Endpoint) UpdateBulkLoadRMM(ctx context.Context, name string, data map[string]map[string]string, fields string) (*BulkMapOfMap, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/reference_data/map_of_sets/bulk_load/"
	reqURL.Path += name

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
	var response *BulkMapOfMap

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// DeleteReferenceMapOfMap by name
func (endpoint *Endpoint) DeleteReferenceMapOfMap(ctx context.Context, name, fields, namespace string, purgeOnly bool) error {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/reference_data/map_of_sets/"
	reqURL.Path += name
	parameters := url.Values{}
	parameters.Add("purge_only", strconv.FormatBool(purgeOnly))
	parameters.Add("fields", fields)
	parameters.Add("namespace", namespace)
	reqURL.RawQuery = parameters.Encode()

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
	if resp.StatusCode != 202 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Status code is %d : Error while reading the body", resp.StatusCode)
		}

		return fmt.Errorf("Status code is %d : %s", resp.StatusCode, string(body))
	}

	return nil
}
