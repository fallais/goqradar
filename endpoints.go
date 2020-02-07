package goqradar

import (
	"context"
	"net/http"
)

// Access endpoint.
type Access interface{}

// Analytics endpoint.
type Analytics interface{}

// AssetModel endpoint.
type AssetModel interface{}

// Auth endpoint.
type Auth interface{}

// BackupAndRestore endpoint.
type BackupAndRestore interface{}

// BandwithManager endpoint.
type BandwithManager interface{}

// Config endpoint.
type Config interface{}

// DataClassification endpoint.
type DataClassification interface{}

// Forensics endpoint.
type Forensics interface{}

// GUIAppFramework endpoint.
type GUIAppFramework interface{}

// Health endpoint.
type Health interface{}

// HealthData endpoint.
type HealthData interface{}

// Help endpoint.
type Help interface{}

// SIEM endpoint.
type SIEM interface {
	ListOffenses(context.Context, string, string, int, int) ([]*Offense, int, error)
}

// ReferenceData endpoint.
type ReferenceData interface{}

// Ariel endpoint.
type Ariel interface {
	ListDatabases() (*http.Response, error)
	GetDatabase(string) (*http.Response, error)
	ListSearches() (*http.Response, error)
	CreateSearch(string) (*http.Response, error)
	GetSearch(string) (*http.Response, error)
	GetSearchResults(string) (*http.Response, error)
	UpdateSearch(string) (*http.Response, error)
	DeleteSearch(string) (*http.Response, error)
}
