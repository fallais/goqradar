package goqradar

import (
	"net/http"
)

const (
	defaultVersion = "9.0"
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

	// Services
	SIEM          *SIEMService
	ReferenceData *ReferenceDataService
	//Analytics
	//Ariel
	//AssetModel
	//Auth
	//Config
	//DataClassification
	//Forensics
	//GUIAppFramework
	//HealthData
	//Help
	//QNI
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
	c.SIEM = &SIEMService{client: c}

	return c
}
