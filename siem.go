package goqradar

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

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

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListOffenses returns the offenses with given fields and filters.
func (endpoint *Endpoint) ListOffenses(ctx context.Context, fields, filter, sort string, min, max int) ([]*Offense, int, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, 0, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/api/siem/offenses"
	parameters := url.Values{}

	// Set fields
	if fields != "" {
		parameters.Add("fields", fields)
	}

	// Set filter
	if filter != "" {
		parameters.Add("filter", filter)
	}

	// Set sort
	if sort != "" {
		parameters.Add("sort", sort)
	}

	// Encode parameters
	reqURL.RawQuery = parameters.Encode()

	// Create the request
	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return nil, 0, fmt.Errorf("error while creating the request : %s", err)
	}
	req = req.WithContext(ctx)

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Range", fmt.Sprintf("items=%d-%d", min, max))

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("error while doing the request : %s", err)
	}
	defer resp.Body.Close()

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("error while reading the request : %s", err)
	}

	// Process the Content-Range
	_, _, total, err := parseContentRange(resp.Header.Get("Content-Range"))
	if err != nil {
		return nil, 0, fmt.Errorf("error while parsing the content-range: %s", err)
	}

	// Prepare the response
	var offenses []*Offense

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &offenses)
	if err != nil {
		return nil, 0, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return offenses, total, nil
}

// GetOffense returns the offense with given ID.
func (endpoint *Endpoint) GetOffense(ctx context.Context, id, fields string) (*Offense, error) {
	return nil, nil
}

// UpdateOffense with given ID.
func (endpoint *Endpoint) UpdateOffense(ctx context.Context, id string) ([]*Offense, int, error) {
	return nil, 0, nil
}

// ListOffenseNotes ...
func (endpoint *Endpoint) ListOffenseNotes(ctx context.Context, id string) ([]*Offense, int, error) {
	return nil, 0, nil
}
