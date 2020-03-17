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

// LogSource is Qradar log source.
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

// ProtocolParameters is a LogSources parameters.
type ProtocolParameters struct {
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Value string `json:"value"`
}

// Messages is a LogSources messages.
type Messages struct {
	Severity  string `json:"severity"`
	Text      string `json:"text"`
	Timestamp int    `json:"timestamp"`
}

// Status is a LogSources status.
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

// LogSourcesGroup is a Qradar LogSourcesGroups.
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

// LogSourcesType is Qradar LogSourcesTypes.
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

// ProtocolTypes is a LogSourcesTypes parameter.
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

// Host is a Qradar host.
type Host struct {
	AppMemory            int       `json:"app_memory"`
	Appliance            Appliance `json:"appliance"`
	AverageEps           int       `json:"average_eps"`
	AverageFpm           int       `json:"average_fpm"`
	Components           []string  `json:"components"`
	CompressionEnabled   bool      `json:"compression_enabled"`
	Cpus                 int       `json:"cpus"`
	EncryptionEnabled    bool      `json:"encryption_enabled"`
	EpsAllocation        int       `json:"eps_allocation"`
	EpsRateHardwareLimit int       `json:"eps_rate_hardware_limit"`
	FpmAllocation        int       `json:"fpm_allocation"`
	FpmRateHardwareLimit int       `json:"fpm_rate_hardware_limit"`
	Hostname             string    `json:"hostname"`
	ID                   int       `json:"id"`
	LicenseSerialNumber  string    `json:"license_serial_number"`
	PeakEps              int       `json:"peak_eps"`
	PeakFpm              int       `json:"peak_fpm"`
	PrimaryServerID      int       `json:"primary_server_id"`
	PrivateIP            string    `json:"private_ip"`
	PublicIP             string    `json:"public_ip"`
	SecondaryServerID    int       `json:"secondary_server_id"`
	Status               string    `json:"status"`
	TotalMemory          int       `json:"total_memory"`
	Version              string    `json:"version"`
}

// Appliance is Qradar appliance.
type Appliance struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// HostsPaginatedResponse is the paginated response.
type HostsPaginatedResponse struct {
	Total int     `json:"total"`
	Min   int     `json:"min"`
	Max   int     `json:"max"`
	Hosts []*Host `json:"offenses"`
}

// Tunnel is a host tunnel.
type Tunnel struct {
	LocalPort     int    `json:"local_port"`
	Name          string `json:"name"`
	RemoteHostID  int    `json:"remote_host_id"`
	RemotePort    int    `json:"remote_port"`
	ReverseSource bool   `json:"reverse_source"`
}

// TunnelsPaginatedResponse is the paginated response.
type TunnelsPaginatedResponse struct {
	Total   int       `json:"total"`
	Min     int       `json:"min"`
	Max     int       `json:"max"`
	Tunnels []*Tunnel `json:"offenses"`
}

// LicensePool is pool of QRadar license.
type LicensePool struct {
	Eps Eps `json:"eps"`
	Fpm Fpm `json:"fpm"`
}

// Eps is an eps.
type Eps struct {
	Allocated     int  `json:"allocated"`
	Overallocated bool `json:"overallocated"`
	Total         int  `json:"total"`
}

// Fpm is a fpm.
type Fpm struct {
	Allocated     int  `json:"allocated"`
	Overallocated bool `json:"overallocated"`
	Total         int  `json:"total"`
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/config/event_sources/log_source_management/log_sources", options...)
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/config/event_sources/log_source_management/log_source_groups", options...)
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/config/event_sources/log_source_management/log_source_types", options...)
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

// ListHosts retrieves a list of hosts.
func (endpoint *Endpoint) ListHosts(ctx context.Context, fields string, filter string, min, max int) (*HostsPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/config/deployment/hosts", options...)
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
	response := &HostsPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.Hosts)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// GetHost retrieves a host
func (endpoint *Endpoint) GetHost(ctx context.Context, id int, fields string) (*Host, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/config/deployment/hosts/"+strconv.Itoa(id), options...)
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
	var response *Host

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// UpdateHost updates a host.
func (endpoint *Endpoint) UpdateHost(ctx context.Context, fields string, data map[string]string, id int) (*Host, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/config/deployment/hosts/"
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
	var host *Host

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &host)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return host, nil
}

// ListTunnels retrieves a list of tunnel of a host.
func (endpoint *Endpoint) ListTunnels(ctx context.Context, fields string, filter string, id int, min, max int) (*TunnelsPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/config/deployment/hosts/"+strconv.Itoa(id)+"/tunnels", options...)
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
	response := &TunnelsPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.Tunnels)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// GetLicensePool retrieves the deployed license pool information.
func (endpoint *Endpoint) GetLicensePool(ctx context.Context, fields string) (*LicensePool, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/config/deployment/license_pool", options...)
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
	var response *LicensePool

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}
