package goqradar

import (
	"net/http"
)

const (
	defaultVersion = "12.0"
)

// Client is a client for QRadar REST API.
type Client struct {
	client *http.Client

	// Base URL for API requests.
	BaseURL string

	// Token
	Token string

	// Version
	Version string

	// Endpoints
	Access             Access
	SIEM               SIEM
	ReferenceData      ReferenceData
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
}

// NewClient returns a new QRadar API client.
func NewClient(httpClient *http.Client, baseURL, token string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	// Create the client
	c := &Client{
		client:  httpClient,
		BaseURL: baseURL,
		Token:   token,
		Version: defaultVersion,
	}

	// Add the services
	c.SIEM = &Service{
		client: c,
	}

	return c
}
