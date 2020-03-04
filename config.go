package goqradar

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// LogSource is Qradar log source
type LogSource struct {
	SendingIP                        string               `json:"sending_ip"`
	Internal                         bool                 `json:"internal"`
	LegacyBulkGroupName              string               `json:"legacy_bulk_group_name"`
	ProtocolParameters               []ProtocolParameters `json:"protocol_parameters"`
	Description                      string               `json:"description"`
	CoalesceEvents                   bool                 `json:"coalesce_events"`
	Enabled                          bool                 `json:"enabled"`
	GroupIds                         []int                `json:"group_ids"`
	AverageEps                       int                  `json:"average_eps"`
	Credibility                      int                  `json:"credibility"`
	ID                               int                  `json:"id"`
	StoreEventPayload                bool                 `json:"store_event_payload"`
	TargetEventCollectorID           int                  `json:"target_event_collector_id"`
	ProtocolTypeID                   int                  `json:"protocol_type_id"`
	LanguageID                       int                  `json:"language_id"`
	CreationDate                     int                  `json:"creation_date"`
	LogSourceExtensionID             int                  `json:"log_source_extension_id"`
	WincollectExternalDestinationIds []int                `json:"wincollect_external_destination_ids"`
	Name                             string               `json:"name"`
	AutoDiscovered                   bool                 `json:"auto_discovered"`
	ModifiedDate                     int                  `json:"modified_date"`
	TypeID                           int                  `json:"type_id"`
	LastEventTime                    int                  `json:"last_event_time"`
	RequiresDeploy                   bool                 `json:"requires_deploy"`
	Gateway                          bool                 `json:"gateway"`
	WincollectInternalDestinationID  int                  `json:"wincollect_internal_destination_id"`
	Status                           Status               `json:"status"`
}

// ProtocolParameters is a LogSources parameters
type ProtocolParameters struct {
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Value string `json:"value"`
}

// Messages is a LogSources messages
type Messages struct {
	Severity  string `json:"severity"`
	Text      string `json:"text"`
	Timestamp int    `json:"timestamp"`
}

// Status is a LogSources status
type Status struct {
	LastUpdated int        `json:"last_updated"`
	Messages    []Messages `json:"messages"`
	Status      string     `json:"status"`
}

// LogSourcesPaginatedResponse is the paginated response.
type LogSourcesPaginatedResponse struct {
	Total      int          `json:"total"`
	Min        int          `json:"min"`
	Max        int          `json:"max"`
	LogSources []*LogSource `json:"offenses"`
}

// LogSourcesGroup is a Qradar LogSourcesGroups
type LogSourcesGroup struct {
	Assignable       bool   `json:"assignable"`
	ChildGroupIds    []int  `json:"child_group_ids"`
	Description      string `json:"description"`
	ID               int    `json:"id"`
	ModificationDate int    `json:"modification_date"`
	Name             string `json:"name"`
	Owner            string `json:"owner"`
	ParentID         int    `json:"parent_id"`
}

// LogSourcesGroupsPaginatedResponse is the paginated response.
type LogSourcesGroupsPaginatedResponse struct {
	Total            int                `json:"total"`
	Min              int                `json:"min"`
	Max              int                `json:"max"`
	LogSourcesGroups []*LogSourcesGroup `json:"offenses"`
}

// LogSourcesType is Qradar LogSourcesTypes
type LogSourcesType struct {
	Custom               bool            `json:"custom"`
	DefaultProtocolID    int             `json:"default_protocol_id"`
	ID                   int             `json:"id"`
	Internal             bool            `json:"internal"`
	LogSourceExtensionID int             `json:"log_source_extension_id"`
	Name                 string          `json:"name"`
	ProtocolTypes        []ProtocolTypes `json:"protocol_types"`
	SupportedLanguageIds []int           `json:"supported_language_ids"`
	Version              string          `json:"version"`
}

// ProtocolTypes is a LogSourcesTypes parameter
type ProtocolTypes struct {
	Documented bool `json:"documented"`
	ProtocolID int  `json:"protocol_id"`
}

// LogSourcesTypesPaginatedResponse is the paginated response.
type LogSourcesTypesPaginatedResponse struct {
	Total           int               `json:"total"`
	Min             int               `json:"min"`
	Max             int               `json:"max"`
	LogSourcesTypes []*LogSourcesType `json:"offenses"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListLogSources retrieves a list of log sources.
func (endpoint *Endpoint) ListLogSources(ctx context.Context, fields string, filter string, sort string, min, max int) (*LogSourcesPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(http.MethodGet, "/config/event_sources/log_source_management/log_sources", options...)
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
	response := &LogSourcesPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.LogSources)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// ListLogSourcesGroups retrieves the list of log source groups.
func (endpoint *Endpoint) ListLogSourcesGroups(ctx context.Context, fields string, filter string, min, max int) (*LogSourcesGroupsPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(http.MethodGet, "/config/event_sources/log_source_management/log_source_groups", options...)
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
	response := &LogSourcesGroupsPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.LogSourcesGroups)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// ListLogSourceTypes retrieves a list of log source types.
func (endpoint *Endpoint) ListLogSourceTypes(ctx context.Context, fields string, filter string, min, max int) (*LogSourcesTypesPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(http.MethodGet, "/config/event_sources/log_source_management/log_source_types", options...)
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
	response := &LogSourcesTypesPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.LogSourcesTypes)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
