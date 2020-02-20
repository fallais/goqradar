package goqradar

import (
	"net/http"
)

const (
	defaultVersion = "12.0"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Client is a client for QRadar REST API.
type Client struct {
	client *http.Client

	// BaseURL is the base URL for API requests.
	BaseURL string

	// Token is the security token.
	Token string

	// Version is the API version.
	Version string

	// Endpoints
	Access             Access
	Analytics          Analytics
	Ariel              Ariel
	AssetModel         AssetModel
	Auth               Auth
	BackupAndRestore   BackupAndRestore
	BandwithManager    BandwithManager
	Config             Config
	DataClassification DataClassification
	Forensics          Forensics
	GUIAppFramework    GUIAppFramework
	Health             Health
	HealthData         HealthData
	Help               Help
	ReferenceData      ReferenceData
	SIEM               SIEM
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewClient returns a new QRadar API client.
func NewClient(httpClient *http.Client, baseURL, token string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// Create the client
	c := &Client{
		client:  httpClient,
		BaseURL: baseURL,
		Token:   token,
		Version: defaultVersion,
	}

	// Add the endpoints
	c.Access = &Endpoint{client: c}
	//c.Ariel = &Endpoint{client: c}
	c.AssetModel = &Endpoint{client: c}
	c.SIEM = &Endpoint{client: c}
	c.ReferenceData = &Endpoint{client: c}
	c.Ariel = &Endpoint{client: c}

	return c
}
