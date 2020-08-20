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
	DisasterRecovery   DisasterRecovery
	DynamicSearch      DynamicSearch
	Forensics          Forensics
	GUIAppFramework    GUIAppFramework
	Health             Health
	HealthData         HealthData
	Help               Help
	Qni                Qni
	Qrm                Qrm
	Qvm                Qvm
	ReferenceData      ReferenceData
	Scanner            Scanner
	Services           Services
	SIEM               SIEM
	StagedConfig       StagedConfig
	System             System
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
	c.Analytics = &Endpoint{client: c}
	c.Ariel = &Endpoint{client: c}
	c.Auth = &Endpoint{client: c}
	c.AssetModel = &Endpoint{client: c}
	c.BackupAndRestore = &Endpoint{client: c}
	c.BandwithManager = &Endpoint{client: c}
	c.Config = &Endpoint{client: c}
	c.DataClassification = &Endpoint{client: c}
	c.DisasterRecovery = &Endpoint{client: c}
	c.DynamicSearch = &Endpoint{client: c}
	c.Forensics = &Endpoint{client: c}
	c.GUIAppFramework = &Endpoint{client: c}
	c.HealthData = &Endpoint{client: c}
	c.Help = &Endpoint{client: c}
	c.Qni = &Endpoint{client: c}
	c.Qrm = &Endpoint{client: c}
	c.Qvm = &Endpoint{client: c}
	c.ReferenceData = &Endpoint{client: c}
	c.Scanner = &Endpoint{client: c}
	c.Services = &Endpoint{client: c}
	c.SIEM = &Endpoint{client: c}
	c.StagedConfig = &Endpoint{client: c}
	c.System = &Endpoint{client: c}

	return c
}
