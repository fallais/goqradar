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

// Recovery is recovery
type Recovery struct {
	AssignedTo              string   `json:"assigned_to"`
	Bpf                     string   `json:"bpf"`
	CaseID                  int      `json:"case_id"`
	CollectionNameSuffix    string   `json:"collection_name_suffix"`
	ID                      int      `json:"id"`
	RecoveryTaskIds         []int    `json:"recovery_task_ids"`
	RecoveryWindowEndTime   int      `json:"recovery_window_end_time"`
	RecoveryWindowStartTime int      `json:"recovery_window_start_time"`
	SessionIds              []string `json:"session_ids"`
	Tags                    []string `json:"tags"`
}

// RecoveriesPaginatedResponse is the paginated response.
type RecoveriesPaginatedResponse struct {
	Total      int         `json:"total"`
	Min        int         `json:"min"`
	Max        int         `json:"max"`
	Recoveries []*Recovery `json:"offense_types"`
}

// RecoveryTask is a recovery task
type RecoveryTask struct {
	Assignee                string   `json:"assignee"`
	Bpf                     string   `json:"bpf"`
	CaptureDeviceIP         string   `json:"capture_device_ip"`
	CaseID                  int      `json:"case_id"`
	CollectionName          string   `json:"collection_name"`
	ID                      int      `json:"id"`
	ManagedHostHostname     string   `json:"managed_host_hostname"`
	RecoveryID              int      `json:"recovery_id"`
	RecoveryWindowEndTime   int      `json:"recovery_window_end_time"`
	RecoveryWindowStartTime int      `json:"recovery_window_start_time"`
	Status                  string   `json:"status"`
	Tags                    []string `json:"tags"`
	TaskEndTime             int      `json:"task_end_time"`
	TaskStartTime           int      `json:"task_start_time"`
}

// RecoveryTasksPaginatedResponse is the paginated response.
type RecoveryTasksPaginatedResponse struct {
	Total         int             `json:"total"`
	Min           int             `json:"min"`
	Max           int             `json:"max"`
	RecoveryTasks []*RecoveryTask `json:"offense_types"`
}

// CaseCreateTask is a case create task
type CaseCreateTask struct {
	AssignedTo []string `json:"assigned_to"`
	CaseID     int      `json:"case_id"`
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	State      string   `json:"state"`
}

// Case is a case
type Case struct {
	AssignedTo []string `json:"assigned_to"`
	ID         int      `json:"id"`
	Name       string   `json:"name"`
}

// CasesPaginatedResponse is the paginated response.
type CasesPaginatedResponse struct {
	Total int     `json:"total"`
	Min   int     `json:"min"`
	Max   int     `json:"max"`
	Cases []*Case `json:"offense_types"`
}

// CreateCase is the struct for case who will be created
type CreateCase struct {
	AssignedTo []string `json:"assigned_to"`
	CaseID     int      `json:"case_id"`
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	State      string   `json:"state"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListRecoveries returns the recoveries with given fields and filters.
func (endpoint *Endpoint) ListRecoveries(ctx context.Context, fields, filter string, min, max int) (*RecoveriesPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/forensics/capture/recoveries", options...)
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
	response := &RecoveriesPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.Recoveries)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// CreateRecovery creates a recovery
func (endpoint *Endpoint) CreateRecovery(ctx context.Context, data map[string]string, fields string) (*Recovery, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/forensics/capture/recoveries"

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

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *Recovery

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// GetRecovery retrieves a recovery .
func (endpoint *Endpoint) GetRecovery(ctx context.Context, id int, fields string) (*Recovery, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/forensics/capture/recoveries/"+strconv.Itoa(id), options...)
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
	var response *Recovery

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// ListRecoveryTasks returns the recoveries with given fields and filters.
func (endpoint *Endpoint) ListRecoveryTasks(ctx context.Context, fields, filter string, min, max int) (*RecoveryTasksPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/forensics/capture/recovery_tasks", options...)
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
	response := &RecoveryTasksPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.RecoveryTasks)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// GetRecoveryTask retrieves a recovery task.
func (endpoint *Endpoint) GetRecoveryTask(ctx context.Context, id int, fields string) (*RecoveryTask, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/forensics/capture/recovery_tasks/"+strconv.Itoa(id), options...)
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
	var response *RecoveryTask

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// GetCaseCreatetask retrieves a case create task.
func (endpoint *Endpoint) GetCaseCreatetask(ctx context.Context, id int, fields string) (*CaseCreateTask, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/forensics/case_management/case_create_tasks/"+strconv.Itoa(id), options...)
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
	var response *CaseCreateTask

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// ListCases returns the cases with given fields and filters.
func (endpoint *Endpoint) ListCases(ctx context.Context, fields, filter string, min, max int) (*CasesPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/forensics/case_management/cases", options...)
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
	response := &CasesPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.Cases)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// CreateCase creates a case
func (endpoint *Endpoint) CreateCase(ctx context.Context, data map[string]string, fields string) (*CreateCase, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/forensics/case_management/cases"

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

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *CreateCase

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// GetCase retrieves a case.
func (endpoint *Endpoint) GetCase(ctx context.Context, id int, fields string) (*Case, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/forensics/case_management/cases/"+strconv.Itoa(id), options...)
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
	var response *Case

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}
