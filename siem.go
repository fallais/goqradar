package qradar

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
}

// SIEMService is the service of SIEM endpoint.
type SIEMService service
