package goqradar

import (
	"context"
	"net/http"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Endpoint is an API endpoint.
type Endpoint struct {
	client *Client
}

//------------------------------------------------------------------------------
// Interfaces
//------------------------------------------------------------------------------

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
	ListOffenses(context.Context, string, string, string, int, int) (*OffensePaginatedResponse, error)
	GetOffense(context.Context, string, string) (*Offense, error)
	UpdateOffense(context.Context, string) ([]*Offense, int, error)
	ListOffenseNotes(context.Context, string) ([]*Offense, int, error)
}

// ReferenceData endpoint.
type ReferenceData interface {
	UpdateBulkLoadRM(context.Context, string, map[string]string, string) (*BulkMap, error)
	DeleteReferenceMap(context.Context, string, string, bool) error
	ListSets(context.Context, string, string, int, int) ([]*Set, error)
	UpdateBulkLoadRS(context.Context, string, []string, string) (*Set, error)
	DeleteReferenceSet(context.Context, string, string, bool) error
	UpdateBulkLoadRT(context.Context, string, map[string]map[string]string, string) (*BulkTable, error)
	DeleteReferenceTable(context.Context, string, string, bool) error
}

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
