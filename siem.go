package goqradar

import (
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

// OffenseType  is a QRadar offense type.
type OffenseType struct {
	Custom       bool   `json:"custom"`
	DatabaseType string `json:"database_type"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PropertyName string `json:"property_name"`
}

// Rule is a QRadar rule.
type Rule struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

// Offense is a QRadar offense.
type Offense struct {
	UsernameCount              int      `json:"username_count"`
	Description                string   `json:"description"`
	Rules                      []Rule   `json:"rules"`
	EventCount                 int      `json:"event_count"`
	FlowCount                  int      `json:"flow_count"`
	AssignedTo                 string   `json:"assigned_to"`
	SecurityCategoryCount      int      `json:"security_category_count"`
	FollowUp                   bool     `json:"follow_up"`
	SourceAddressIds           []int    `json:"source_address_ids"`
	SourceCount                int      `json:"source_count"`
	Inactive                   bool     `json:"inactive"`
	Protected                  bool     `json:"protected"`
	CategoryCount              int      `json:"category_count"`
	SourceNetwork              string   `json:"source_network"`
	DestinationNetworks        []string `json:"destination_networks"`
	ClosingUser                string   `json:"closing_user"`
	CloseTime                  int64    `json:"close_time"`
	RemoteDestinationCount     int      `json:"remote_destination_count"`
	StartTime                  int64    `json:"start_time"`
	LastUpdatedTime            int64    `json:"last_updated_time"`
	Credibility                int      `json:"credibility"`
	Magnitude                  int      `json:"magnitude"`
	ID                         int      `json:"id"`
	Categories                 []string `json:"categories"`
	Severity                   int      `json:"severity"`
	PolicyCategoryCount        int      `json:"policy_category_count"`
	DeviceCount                int      `json:"device_count"`
	ClosingReasonID            int      `json:"closing_reason_id"`
	OffenseType                int      `json:"offense_type"`
	Relevance                  int      `json:"relevance"`
	DomainID                   int      `json:"domain_id"`
	OffenseSource              string   `json:"offense_source"`
	LocalDestinationAddressIds []int    `json:"local_destination_address_ids"`
	LocalDestinationCount      int      `json:"local_destination_count"`
	Status                     string   `json:"status"`
	Notes                      []*Note  `json:"notes,omitempty"`
}

// Note is a QRadar note
type Note struct {
	ID         int    `json:"id"`
	CreateTime int    `json:"create_time"`
	NoteText   string `json:"note_text"`
	Username   string `json:"username"`
}

// OffensePaginatedResponse is the paginated response.
type OffensePaginatedResponse struct {
	Total    int        `json:"total"`
	Min      int        `json:"min"`
	Max      int        `json:"max"`
	Offenses []*Offense `json:"offenses"`
}

// OffenseTypesPaginatedResponse is the paginated response.
type OffenseTypesPaginatedResponse struct {
	Total        int            `json:"total"`
	Min          int            `json:"min"`
	Max          int            `json:"max"`
	OffenseTypes []*OffenseType `json:"offense_types"`
}

// LocalDestinationAddress is a QRadar local destination address
type LocalDestinationAddress struct {
	DomainID           int    `json:"domain_id"`
	EventFlowCount     int    `json:"event_flow_count"`
	FirstEventFlowSeen int    `json:"first_event_flow_seen"`
	ID                 int    `json:"id"`
	LastEventFlowSeen  int    `json:"last_event_flow_seen"`
	LocalDestinationIP string `json:"local_destination_ip"`
	Magnitude          int    `json:"magnitude"`
	Network            string `json:"network"`
	OffenseIds         []int  `json:"offense_ids"`
	SourceAddressIds   []int  `json:"source_address_ids"`
}

// LocalDestinationAddressesPaginatedResponse is the paginated response.
type LocalDestinationAddressesPaginatedResponse struct {
	Total                     int                        `json:"total"`
	Min                       int                        `json:"min"`
	Max                       int                        `json:"max"`
	LocalDestinationAddresses []*LocalDestinationAddress `json:"offense_types"`
}

// SourceAddress is a QRadar local source address
type SourceAddress struct {
	DomainID                   int    `json:"domain_id"`
	EventFlowCount             int    `json:"event_flow_count"`
	FirstEventFlowSeen         int    `json:"first_event_flow_seen"`
	ID                         int    `json:"id"`
	LastEventFlowSeen          int    `json:"last_event_flow_seen"`
	LocalDestinationAddressIds []int  `json:"local_destination_address_ids"`
	Magnitude                  int    `json:"magnitude"`
	Network                    string `json:"network"`
	OffenseIds                 []int  `json:"offense_ids"`
	SourceIP                   string `json:"source_ip"`
}

// SourceAddressesPaginatedResponse is the paginated response.
type SourceAddressesPaginatedResponse struct {
	Total           int              `json:"total"`
	Min             int              `json:"min"`
	Max             int              `json:"max"`
	SourceAddresses []*SourceAddress `json:"offense_types"`
}

// OffenseClosingReason is a QRadar local offense closing Reason
type OffenseClosingReason struct {
	ID         int    `json:"id"`
	IsDeleted  bool   `json:"is_deleted"`
	IsReserved bool   `json:"is_reserved"`
	Text       string `json:"text"`
}

// OffenseClosingReasonsPaginatedResponse is the paginated response.
type OffenseClosingReasonsPaginatedResponse struct {
	Total                 int                     `json:"total"`
	Min                   int                     `json:"min"`
	Max                   int                     `json:"max"`
	OffenseClosingReasons []*OffenseClosingReason `json:"offense_types"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListOffenses returns the offenses with given fields, filters and sort.
func (endpoint *Endpoint) ListOffenses(ctx context.Context, fields, filter, sort string, min, max int) (*OffensePaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/siem/offenses", options...)
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
	response := &OffensePaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.Offenses)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// GetOffense returns the offense by given ID.
func (endpoint *Endpoint) GetOffense(ctx context.Context, id int, fields string) (*Offense, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/siem/offenses/"+strconv.Itoa(id), options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *Offense

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// UpdateOffense with given ID.
func (endpoint *Endpoint) UpdateOffense(ctx context.Context, id int, assignedTo string) error {
	return nil
}

// ListOffenseNotes returns the notes of the given offense.
func (endpoint *Endpoint) ListOffenseNotes(ctx context.Context, id string) ([]*Note, int, error) {
	// Options
	options := []Option{}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/siem/offenses/"+id+"/notes", options...)
	if err != nil {
		return nil, 0, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, 0, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response []*Note

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, 0, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, len(response), nil
}

// CreateOffenseNote ...
func (endpoint *Endpoint) CreateOffenseNote(ctx context.Context, id string) ([]*Note, int, error) {
	return nil, 0, nil
}

//------------------------------------------------------------------------------

// ListOffenseTypes returns the offenses type with given fields, filters and sort.
func (endpoint *Endpoint) ListOffenseTypes(ctx context.Context, fields, filter, sort string, min, max int) (*OffenseTypesPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/siem/offenses_types", options...)
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
	response := &OffenseTypesPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.OffenseTypes)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// GetOffenseType returns the offense type by ID with given fields.
func (endpoint *Endpoint) GetOffenseType(ctx context.Context, id, fields string) (*OffenseType, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/siem/offenses_types/"+id, options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *OffenseType

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// ListLocalDestinationAddress returns the local destination addresses with given fields, filters and sort.
func (endpoint *Endpoint) ListLocalDestinationAddress(ctx context.Context, fields, filter string, min, max int) (*LocalDestinationAddressesPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/siem/local_destination_addresses", options...)
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
	response := &LocalDestinationAddressesPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.LocalDestinationAddresses)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// GetLocalDestinationAddress retrieve an offense local destination address whith given filters.
func (endpoint *Endpoint) GetLocalDestinationAddress(ctx context.Context, id int, fields string) (*LocalDestinationAddress, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/siem/local_destination_addresses/"+strconv.Itoa(id), options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *LocalDestinationAddress

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// ListSourceAddresses returns a list offense source addresses currently in the system with given fields, filters and sort.
func (endpoint *Endpoint) ListSourceAddresses(ctx context.Context, fields, filter string, min, max int) (*SourceAddressesPaginatedResponse, error) {
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/siem/source_addresses", options...)
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
	response := &SourceAddressesPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.SourceAddresses)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// GetSourceAddress retrieve an offense source address.
func (endpoint *Endpoint) GetSourceAddress(ctx context.Context, id int, fields string) (*SourceAddress, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/siem/source_addresses/"+strconv.Itoa(id), options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *SourceAddress

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// ListOffenseClosingReasons returns a list of all offense closing reasons with given fields, filters and sort.
func (endpoint *Endpoint) ListOffenseClosingReasons(ctx context.Context, fields, filter string, includeDeleted, includedReserved bool, min, max int) (*OffenseClosingReasonsPaginatedResponse, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}
	if filter != "" {
		options = append(options, WithParam("filter", filter))
	}
	if includeDeleted != false {
		options = append(options, WithParam("includeDeleted", strconv.FormatBool(includeDeleted)))
	}
	if includedReserved != false {
		options = append(options, WithParam("includedReserved", strconv.FormatBool(includedReserved)))
	}
	options = append(options, WithHeader("Range", fmt.Sprintf("items=%d-%d", min, max)))

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/siem/offense_closing_reasons", options...)
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
	response := &OffenseClosingReasonsPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.OffenseClosingReasons)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// CreateOffenseClosingReason create an offense closing reason.
func (endpoint *Endpoint) CreateOffenseClosingReason(ctx context.Context, reason, fields string) (*OffenseClosingReason, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/siem/offense_closing_reasons"

	// Create the request
	req, err := http.NewRequest("POST", reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error while creating the request : %s", err)
	}

	// Add optional parameters
	q := req.URL.Query()
	q.Add("reason", reason)
	q.Add("filters", fields)
	req.URL.RawQuery = q.Encode()

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/json")

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while doing the request : %s", err)
	}

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *OffenseClosingReason

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// GetOffenseClosingReason retrieve an offense closing reason.
func (endpoint *Endpoint) GetOffenseClosingReason(ctx context.Context, id int, fields string) (*OffenseClosingReason, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/siem/offense_closing_reasons/"+strconv.Itoa(id), options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *OffenseClosingReason

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
