package qradar

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	defaultVersion = "9.0"
)

// A Client manages communication with the QRadar API.
type Client struct {
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL
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
func NewClient(httpClient *http.Client, b, t string) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	// Parse the URL
	baseURL, err := url.Parse(b)
	if err != nil {
		return nil, fmt.Errorf("error while parsing the base URL : %s", err)
	}

	// Create the client
	c := &Client{
		client:  httpClient,
		BaseURL: baseURL,
		Token:   t,
		Version: defaultVersion,
	}

	// Add the services
	c.SIEM = &SIEMService{client: c}

	return c, nil
}
