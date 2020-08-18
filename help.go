package goqradar

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//------------------------------------------------------------------------------
// Structures
//------------------------------------------------------------------------------

// EndpointDocumentationObject is an endpoint documentation objects that are currently in the system.
type EndpointDocumentationObject struct {
	CallerHasAccess bool   `json:"caller_has_access"`
	Deprecated      bool   `json:"deprecated"`
	Description     string `json:"description"`
	ErrorResponses  []struct {
		Description             string `json:"description"`
		ResponseCode            int    `json:"response_code"`
		ResponseCodeDescription string `json:"response_code_description"`
		UniqueCode              int    `json:"unique_code"`
	} `json:"error_responses"`
	HTTPMethod          string `json:"http_method"`
	ID                  int    `json:"id"`
	LastModifiedVersion string `json:"last_modified_version"`
	Parameters          []struct {
		DefaultValue string `json:"default_value"`
		Description  string `json:"description"`
		MimeTypes    []struct {
			DataType string `json:"data_type"`
			MimeType string `json:"mime_type"`
			Sample   string `json:"sample"`
		} `json:"mime_types"`
		ParameterName string `json:"parameter_name"`
		Type          string `json:"type"`
	} `json:"parameters"`
	Path                string `json:"path"`
	ResourceID          int    `json:"resource_id"`
	ResponseDescription string `json:"response_description"`
	ResponseMimeTypes   []struct {
		MimeType  string `json:"mime_type"`
		Sample    string `json:"sample"`
		MediaType string `json:"media_type"`
	} `json:"response_mime_types"`
	SuccessResponses []struct {
		Description             string `json:"description"`
		ResponseCode            int    `json:"response_code"`
		ResponseCodeDescription string `json:"response_code_description"`
	} `json:"success_responses"`
	Summary string `json:"summary"`
	Version string `json:"version"`
}

// EndpointDocumentationObjectsPaginatedResponse is the paginated response.
type EndpointDocumentationObjectsPaginatedResponse struct {
	Total                        int                            `json:"total"`
	Min                          int                            `json:"min"`
	Max                          int                            `json:"max"`
	EndpointDocumentationObjects []*EndpointDocumentationObject `json:"offense_types"`
}

// ResourceDocumentationObject is an endpoint resource documentation objects that are currently in the system.
type ResourceDocumentationObject struct {
	ChildResourceIds []int  `json:"child_resource_ids"`
	EndpointIds      []int  `json:"endpoint_ids"`
	ID               int    `json:"id"`
	ParentResourceID int    `json:"parent_resource_id"`
	Path             string `json:"path"`
	Resource         string `json:"resource"`
	Version          string `json:"version"`
}

// ResourceDocumentationObjectsPaginatedResponse is the paginated response.
type ResourceDocumentationObjectsPaginatedResponse struct {
	Total                        int                            `json:"total"`
	Min                          int                            `json:"min"`
	Max                          int                            `json:"max"`
	ResourceDocumentationObjects []*ResourceDocumentationObject `json:"offense_types"`
}

// VersionDocumentationObject is a version documentation objects that are currently in the system.
type VersionDocumentationObject struct {
	Deprecated      bool   `json:"deprecated"`
	ID              int    `json:"id"`
	Removed         bool   `json:"removed"`
	RootResourceIds []int  `json:"root_resource_ids"`
	Version         string `json:"version"`
}

// VersionDocumentationObjectsPaginatedResponse is the paginated response.
type VersionDocumentationObjectsPaginatedResponse struct {
	Total                       int                           `json:"total"`
	Min                         int                           `json:"min"`
	Max                         int                           `json:"max"`
	VersionDocumentationObjects []*VersionDocumentationObject `json:"offense_types"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListEndpointDocumentationObjects returns the object with given fields, filters and sort.
func (endpoint *Endpoint) ListEndpointDocumentationObjects(ctx context.Context, fields, filter string, min, max int) (*EndpointDocumentationObjectsPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/help/endpoints", options...)
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
	response := &EndpointDocumentationObjectsPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.EndpointDocumentationObjects)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// GetEndpointDocumentationObject retrieves an endpoint object by his id.
func (endpoint *Endpoint) GetEndpointDocumentationObject(ctx context.Context, id int, fields string) (*EndpointDocumentationObject, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/help/endpoints/"+strconv.Itoa(id), options...)
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
	var response *EndpointDocumentationObject

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// ListResourceDocumentationObjects returns the object with given fields, filters and sort.
func (endpoint *Endpoint) ListResourceDocumentationObjects(ctx context.Context, fields, filter string, min, max int) (*ResourceDocumentationObjectsPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/help/resources", options...)
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
	response := &ResourceDocumentationObjectsPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.ResourceDocumentationObjects)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// GetResourceDocumentationObject retrieves an resource object by his id.
func (endpoint *Endpoint) GetResourceDocumentationObject(ctx context.Context, id int, fields string) (*ResourceDocumentationObject, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/help/resources/"+strconv.Itoa(id), options...)
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
	var response *ResourceDocumentationObject

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// ListVersionDocumentationObjects returns the object with given fields, filters and sort.
func (endpoint *Endpoint) ListVersionDocumentationObjects(ctx context.Context, fields, filter string, min, max int) (*VersionDocumentationObjectsPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/help/versions", options...)
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
	response := &VersionDocumentationObjectsPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.VersionDocumentationObjects)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// GetVersionDocumentationObject retrieves an version object by his id.
func (endpoint *Endpoint) GetVersionDocumentationObject(ctx context.Context, id int, fields string) (*VersionDocumentationObject, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/help/versions/"+strconv.Itoa(id), options...)
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
	var response *VersionDocumentationObject

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}
