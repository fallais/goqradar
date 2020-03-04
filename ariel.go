package goqradar

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// SavedSearch is a QRadar SavedSearch
type SavedSearch struct {
	Aql           string `json:"aql"`
	CreationDate  int    `json:"creation_date"`
	Database      string `json:"database"`
	Description   string `json:"description"`
	ID            int    `json:"id"`
	IsAggregate   bool   `json:"is_aggregate"`
	IsDashboard   bool   `json:"is_dashboard"`
	IsDefault     bool   `json:"is_default"`
	IsQuickSearch bool   `json:"is_quick_search"`
	IsShared      bool   `json:"is_shared"`
	ModifiedDate  int    `json:"modified_date"`
	Name          string `json:"name"`
	Owner         string `json:"owner"`
	UID           string `json:"uid"`
}

// SavedSearchDependentTask is a QRadar SavedSearchDependentTask
type SavedSearchDependentTask struct {
	CancelledBy        string           `json:"cancelled_by"`
	Completed          int              `json:"completed"`
	Created            int              `json:"created"`
	CreatedBy          string           `json:"created_by"`
	ID                 int              `json:"id"`
	Maximum            int              `json:"maximum"`
	Message            string           `json:"message"`
	Modified           int              `json:"modified"`
	Name               string           `json:"name"`
	NumberOfDependents int              `json:"number_of_dependents"`
	Progress           int              `json:"progress"`
	Started            int              `json:"started"`
	Status             string           `json:"status"`
	TaskComponents     []TaskComponents `json:"task_components"`
}

// TaskComponents is a QRadar TaskComponents
type TaskComponents struct {
	Completed          int    `json:"completed"`
	Created            int    `json:"created"`
	Maximum            int    `json:"maximum"`
	Message            string `json:"message"`
	Modified           int    `json:"modified"`
	NumberOfDependents int    `json:"number_of_dependents"`
	Progress           int    `json:"progress"`
	Started            int    `json:"started"`
	Status             string `json:"status"`
	TaskSubType        string `json:"task_sub_type"`
}

// Searches is a QRadar searches status
type Searches struct {
	CursorID                 string          `json:"cursor_id"`
	CompressedDataFileCount  int             `json:"compressed_data_file_count"`
	CompressedDataTotalSize  int             `json:"compressed_data_total_size"`
	DataFileCount            int             `json:"data_file_count"`
	DataTotalSize            int             `json:"data_total_size"`
	IndexFileCount           int             `json:"index_file_count"`
	IndexTotalSize           int             `json:"index_total_size"`
	ProcessedRecordCount     int             `json:"processed_record_count"`
	ErrorMessages            []ErrorMessages `json:"error_messages"`
	DesiredRetentionTimeMsec int             `json:"desired_retention_time_msec"`
	Progress                 int             `json:"progress"`
	ProgressDetails          []int           `json:"progress_details"`
	QueryExecutionTime       int             `json:"query_execution_time"`
	QueryString              string          `json:"query_string"`
	RecordCount              int             `json:"record_count"`
	SaveResults              bool            `json:"save_results"`
	Status                   string          `json:"status"`
	Snapshot                 Snapshot        `json:"snapshot"`
	SubsearchIds             []string        `json:"subsearch_ids"`
	SearchID                 string          `json:"search_id"`
}

// ErrorMessages is a QRadar ErrorMessages
type ErrorMessages struct {
	Code     string   `json:"code"`
	Contexts []string `json:"contexts"`
	Message  string   `json:"message"`
	Severity string   `json:"severity"`
}

// Events is a QRadar Events
type Events struct {
	Sourceip   string `json:"sourceip"`
	Starttime  int64  `json:"starttime"`
	Qid        int    `json:"qid"`
	Sourceport int    `json:"sourceport"`
}

// Snapshot is a QRadar TaskCompSnapshotonents
type Snapshot struct {
	Events []Events `json:"events"`
}

// Database is a QRadar database
type Database struct {
	Total   int       `json:"total"`
	Min     int       `json:"min"`
	Max     int       `json:"max"`
	Columns []Columns `json:"columns"`
}

// Columns is a QRadar columns
type Columns struct {
	ArgumentType    string `json:"argument_type"`
	Indexable       bool   `json:"indexable"`
	Name            string `json:"name"`
	Nullable        bool   `json:"nullable"`
	ObjectValueType string `json:"object_value_type"`
	ProviderName    string `json:"provider_name"`
}

// SavedSearchPaginatedResponse is the paginated response.
type SavedSearchPaginatedResponse struct {
	Total       int            `json:"total"`
	Min         int            `json:"min"`
	Max         int            `json:"max"`
	SavedSearch []*SavedSearch `json:"offenses"`
}

// SearchesPaginatedResponse is the paginated response.
type SearchesPaginatedResponse struct {
	Total    int         `json:"total"`
	Min      int         `json:"min"`
	Max      int         `json:"max"`
	Searches []*Searches `json:"offenses"`
}

// DatabasePaginatedResponse is the paginated response.
type DatabasePaginatedResponse struct {
	Total     int      `json:"total"`
	Min       int      `json:"min"`
	Max       int      `json:"max"`
	Databases []string `json:"offenses"`
}

// SearchesResults is the result of an AQL
type SearchesResults struct {
	Events []EventsResults `json:"events"`
}

// EventsResults is an AQL event
type EventsResults struct {
	SourceIP      string `json:"sourceIP"`
	DestinationIP string `json:"destinationIP"`
	Qid           int    `json:"qid"`
}

// SearchesResultsPaginatedResponse is the paginated response.
type SearchesResultsPaginatedResponse struct {
	Total           int                `json:"total"`
	Min             int                `json:"min"`
	Max             int                `json:"max"`
	SearchesResults []*SearchesResults `json:"offenses"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// GetSavedSearch retrieves an Ariel saved search
func (endpoint *Endpoint) GetSavedSearch(ctx context.Context, id int, fields string) (*SavedSearch, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/ariel/saved_searches/"+strconv.Itoa(id), options...)
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
	var response *SavedSearch

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// ListSavedSearch retrieves the list of Ariel saved search
func (endpoint *Endpoint) ListSavedSearch(ctx context.Context, fields string, filter string, min, max int) (*SavedSearchPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/ariel/saved_searches", options...)
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
	response := &SavedSearchPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.SavedSearch)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// GetSavedSearchDependentTask retrieves the dependent Ariel saved search task status
func (endpoint *Endpoint) GetSavedSearchDependentTask(ctx context.Context, taskID int, fields string) (*SavedSearchDependentTask, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/ariel/saved_search_dependent_tasks/"+strconv.Itoa(taskID), options...)
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
	var response *SavedSearchDependentTask

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// GetSearchesID retrieves status information for a search, based on the search ID parameter
func (endpoint *Endpoint) GetSearchesID(ctx context.Context, searchID string, prefer string) (*Searches, error) {
	// Options
	options := []Option{}
	if prefer != "" {
		options = append(options, WithParam("prefer", prefer))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/ariel/searches/"+searchID, options...)
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
	var response *Searches

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// ListSearches retrieves the list of Ariel searches
func (endpoint *Endpoint) ListSearches(ctx context.Context, fields string, filter string, min, max int) (*SearchesPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/ariel/searches", options...)
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
	response := &SearchesPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.Searches)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// GetDatabase retrieve the columns that are defined for the specified Ariel database
func (endpoint *Endpoint) GetDatabase(ctx context.Context, databaseName string, fields string, filter string, min, max int) (*Database, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/ariel/databases/"+databaseName, options...)
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
	response := &Database{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.Columns)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// ListDatabase retrieves the list of Ariel saved search
func (endpoint *Endpoint) ListDatabase(ctx context.Context, filter string, min, max int) (*DatabasePaginatedResponse, error) {
	// Options
	options := []Option{}
	if filter != "" {
		options = append(options, WithParam("filter", filter))
	}
	options = append(options, WithHeader("Range", fmt.Sprintf("items=%d-%d", min, max)))

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/ariel/databases", options...)
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
	response := &DatabasePaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.Databases)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// PostSearches retrieve the results of the Ariel search that is identified by the search ID
func (endpoint *Endpoint) PostSearches(ctx context.Context, searchID string, saveResult bool, status string) (*Searches, error) {
	// Options
	options := []Option{}
	options = append(options, WithParam("saveResult", strconv.FormatBool(saveResult)))

	if status != "" {
		options = append(options, WithParam("status", status))
	}

	// Do the request
	resp, err := endpoint.client.do(http.MethodPost, "/ariel/searches/"+searchID, options...)
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
	var response *Searches

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil

}

// GetSearchesResults retrieve the the results of the Ariel search that is identified by the search ID
func (endpoint *Endpoint) GetSearchesResults(ctx context.Context, searchID string, min, max int) (*SearchesResultsPaginatedResponse, error) {
	// Options
	options := []Option{}
	options = append(options, WithHeader("Range", fmt.Sprintf("items=%d-%d", min, max)))

	// Do the request
	resp, err := endpoint.client.do(http.MethodGet, "/ariel/searches/"+searchID+"/results", options...)
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
	response := &SearchesResultsPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.SearchesResults)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
