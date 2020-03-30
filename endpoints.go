package goqradar

import (
	"context"
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
type Config interface {
	ListLogSources(context.Context, string, string, string, int, int) (*LogSourcesPaginatedResponse, error)
	ListLogSourcesGroups(context.Context, string, string, int, int) (*LogSourcesGroupsPaginatedResponse, error)
	ListLogSourceTypes(context.Context, string, string, int, int) (*LogSourcesTypesPaginatedResponse, error)
	ListHosts(context.Context, string, string, int, int) (*HostsPaginatedResponse, error)
	GetHost(context.Context, int, string) (*Host, error)
	UpdateHost(context.Context, string, map[string]string, int) (*Host, error)
	ListTunnels(context.Context, string, string, int, int, int) (*TunnelsPaginatedResponse, error)
	GetLicensePool(context.Context, string) (*LicensePool, error)
}

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
	ListOffenses(context.Context, string, string, string, int, int) ([]*Offense, int, error)
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
	GetSavedSearch(context.Context, int, string) (*SavedSearch, error)
	ListSavedSearch(context.Context, string, string, int, int) (*SavedSearchPaginatedResponse, error)
	GetSavedSearchDependentTask(context.Context, int, string) (*SavedSearchDependentTask, error)
	GetSearchesID(context.Context, string, string) (*Searches, error)
	ListSearches(context.Context, string, string, int, int) (*SearchesPaginatedResponse, error)
	GetDatabase(context.Context, string, string, string, int, int) (*Database, error)
	ListDatabase(context.Context, string, int, int) (*DatabasePaginatedResponse, error)
	PostSearches(context.Context, string, int) (*Searches, error)
	GetSearchesResults(context.Context, string, int, int) (*SearchesResult, error)
}
