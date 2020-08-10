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
type Access interface {
	ListAccessAttempts(context.Context, string, string, string, int, int) (*LoginAttemptPaginatedResponse, error)
}

// Analytics endpoint.
type Analytics interface {
	ListRules(context.Context, string, string, int, int) (*RulesPaginatedResponse, error)
}

// AssetModel endpoint.
type AssetModel interface {
	ListAssets(context.Context, string, string, string, int, int) (*AssetsPaginatedResponse, error)
	UpdateAsset(context.Context, string, map[string]map[string]string) (string, error)
	ListAssetProperties(context.Context, string, string, int, int) (*AssetPropertiePaginatedResponse, error)
	ListAssetsSavedSearchGroups(context.Context, string, string, int, int) (*AssetSavedSearchGroupPaginatedResponse, error)
	GetAssetSavedSearchGroups(context.Context, int, string) (*AssetSavedSearchGroups, error)
	UpdateAssetSavedSeachGroup(context.Context, int, string, map[string]map[string]string) (*AssetSavedSearchGroups, error)
	DeleteAssetSavedSearchGroups(context.Context, int) error
	ListSavedSearches(context.Context, string, string, int, int) (*SavedSearchesPaginatedResponse, error)
	GetAssetSavedSearch(context.Context, int, string) (*SavedSearche, error)
	UpdateAssetSavedSearch(context.Context, int, map[string]map[string]string) (*SavedSearche, error)
	DeleteAssetSavedSearch(context.Context, int) error
	ListAssetSavedSearches(context.Context, string, string, string, int, int) (*AssetBasedOnSavedSearchPaginatedResponse, error)
}

// Auth endpoint.
type Auth interface {
	Logout(context.Context, string) (bool, error)
}

// BackupAndRestore endpoint.
type BackupAndRestore interface {
	ListBackups(context.Context, string, string, string, int, int) (*BackupsPaginatedResponse, error)
	CreateBackup(context.Context, string, map[string]string) (*Backup, error)
	GetBackup(context.Context, int, string) (*Backup, error)
	UpdateBackup(context.Context, int, map[string]string) (*Backup, error)
	DeleteBackup(context.Context, int) (*Backup, error)
	ListRestore(context.Context, string, string, string, int, int) (*RestoresPaginatedResponse, error)
	CreateRestore(context.Context, map[string]string) (*Restore, error)
	GetRestore(context.Context, int, string) (*Restore, error)
	UpdateRestore(context.Context, int, map[string]string) (*Restore, error)
	DeleteRestore(context.Context, int) error
}

// BandwithManager endpoint.
type BandwithManager interface {
	ListConfigurations(context.Context, string, string, string, int, int) (*ConfigurationsPaginatedResponse, error)
	CreateConfiguration(context.Context, map[string]string) error
	GetConfiguration(context.Context, int, string) (*Configuration, error)
	UpdateConfiguration(context.Context, int, map[string]string, string) (*Configuration, error)
	DeleteConfiguration(context.Context, int) error
	ListEgressFilters(context.Context, string, string, string, int, int) (*EgressFiltersPaginatedResponse, error)
	CreateEgressFilter(context.Context, map[string]string) (*EgressFilter, error)
	GetEgressFilter(context.Context, int, string) (*EgressFilter, error)
	UpdateEgressFilter(context.Context, int, map[string]string, string) (*EgressFilter, error)
	DeleteEgressFilter(context.Context, int) error
}

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
type Help interface {
	ListEndpointDocumentationObjects(context.Context, string, string, int, int) (*EndpointDocumentationObjectsPaginatedResponse, error)
	GetEndpointDocumentationObject(context.Context, int, string) (*EndpointDocumentationObject, error)
	ListResourceDocumentationObjects(context.Context, string, string, int, int) (*ResourceDocumentationObjectsPaginatedResponse, error)
	GetResourceDocumentationObject(context.Context, int, string) (*ResourceDocumentationObject, error)
	ListVersionDocumentationObjects(context.Context, string, string, int, int) (*VersionDocumentationObjectsPaginatedResponse, error)
	GetVersionDocumentationObject(context.Context, int, string) (*VersionDocumentationObject, error)
}

// SIEM endpoint.
type SIEM interface {
	ListOffenses(context.Context, string, string, string, int, int) (*OffensePaginatedResponse, error)
	GetOffense(context.Context, int, string) (*Offense, error)
	//UpdateOffense(context.Context, int, string, string, string, string) error
	//ListOffenseNotes(context.Context, string) ([]*Offense, int, error)
	//CreateOffenseNote(context.Context, string) ([]*Note, int, error)
	ListOffenseTypes(context.Context, string, string, string, int, int) (*OffenseTypesPaginatedResponse, error)
	GetOffenseType(context.Context, string, string) (*OffenseType, error)
}

// ReferenceData endpoint.
type ReferenceData interface {
	UpdateBulkLoadRM(context.Context, string, map[string]string, string) (*BulkMap, error)
	DeleteReferenceMap(context.Context, string, string, bool) error
	ListSets(context.Context, string, string, int, int) (*ListSetsPaginatedResponse, error)
	UpdateBulkLoadRS(context.Context, string, []string, string) (*Set, error)
	DeleteReferenceSet(context.Context, string, string, bool) error
	UpdateBulkLoadRT(context.Context, string, map[string]map[string]string, string) (*BulkTable, error)
	DeleteReferenceTable(context.Context, string, string, bool) error
	UpdateBulkLoadRMM(context.Context, string, map[string]map[string]string, string) (*BulkMapOfMap, error)
	DeleteReferenceMapOfMap(context.Context, string, string, bool) error
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
	GetSearchesResults(context.Context, string, int, int) (*SearchesResult, error)
	PostSearches(context.Context, string, int) (*Searches, error)
}
