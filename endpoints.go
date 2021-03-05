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
	UpdateAssetSavedSearch(context.Context, int, map[string]map[string]string, string) (*SavedSearche, error)
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
	CreateBackup(context.Context, string, string, map[string]string) (*Backup, error)
	GetBackup(context.Context, int, string) (*Backup, error)
	UpdateBackup(context.Context, int, map[string]string, string) (*Backup, error)
	DeleteBackup(context.Context, int) (*Backup, error)
	ListRestore(context.Context, string, string, string, int, int) (*RestoresPaginatedResponse, error)
	CreateRestore(context.Context, map[string]string, string) (*Restore, error)
	GetRestore(context.Context, int, string) (*Restore, error)
	UpdateRestore(context.Context, int, map[string]string, string) (*Restore, error)
	DeleteRestore(context.Context, int) error
}

// BandwithManager endpoint.
type BandwithManager interface {
	ListConfigurations(context.Context, string, string, string, int, int) (*ConfigurationsPaginatedResponse, error)
	CreateConfiguration(context.Context, map[string]string, string) error
	GetConfiguration(context.Context, int, string) (*Configuration, error)
	UpdateConfiguration(context.Context, int, map[string]string, string) (*Configuration, error)
	DeleteConfiguration(context.Context, int) error
	ListEgressFilters(context.Context, string, string, string, int, int) (*EgressFiltersPaginatedResponse, error)
	CreateEgressFilter(context.Context, map[string]string, string) (*EgressFilter, error)
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
type DataClassification interface {
	ListDSMEventMappings(context.Context, string, string, int, int) (*DSMEventMappingsPaginatedResponse, error)
	CreateDSMEventMapping(context.Context, map[string]string, string) (*DSMEventMapping, error)
	GetDSMEventMapping(context.Context, int, string) (*DSMEventMapping, error)
	UpdateDSMEventMapping(context.Context, int, map[string]string, string) (*DSMEventMapping, error)
	ListHLCategories(context.Context, string, string, string, int, int) (*HLCategoriesPaginatedResponse, error)
	GetHLCategory(context.Context, int, string) (*HLCategory, error)
	ListLLCategories(context.Context, string, string, string, int, int) (*LLCategoriesPaginatedResponse, error)
	GetLLCategory(context.Context, int, string) (*LLCategory, error)
	ListQIDRecords(context.Context, string, string, int, int) (*QIDRecordsPaginatedResponse, error)
	CreateQIDRecord(context.Context, map[string]string, string) (*QIDRecord, error)
	GetQIDRecord(context.Context, int, string) (*QIDRecordBYID, error)
	UpdateQIDRecord(context.Context, int, map[string]string, string) (*QIDRecord, error)
}

// DisasterRecovery endpoint.
type DisasterRecovery interface {
	ListArielCopyProfiles(context.Context, string, string) ([]*ArielCopyProfile, error)
	CreateArielCopyProfille(context.Context, map[string]interface{}, string) error
	GetArielCopyProfile(context.Context, int, string) (*ArielCopyProfile, error)
	UpdateArielCopyProfile(context.Context, int, map[string]interface{}, string) (*ArielCopyProfile, error)
	DeleteArielCopyProfile(context.Context, int) error
}

// DynamicSearch endpoint.
type DynamicSearch interface {
	ListSchemas(context.Context, string, string, int, int) (*SchemasPaginatedResponse, error)
	GetSchemas(context.Context, string, string) (*Schemas, error)
	ListFields(context.Context, string, string, string, int, int) (*FieldsPaginatedResponse, error)
	ListFunctions(context.Context, string, string, string, int, int) (*FunctionsPaginatedResponse, error)
	ListOperators(context.Context, string, string, string, int, int) (*OperatorsPaginatedResponse, error)
	ListDynamicSearches(context.Context, string, string, int, int) (*DynamicSearchesPaginatedResponse, error)
	CreateDynamicSearch(context.Context, map[string]interface{}) (*PostedSearch, error)
	GetDynamicSearch(context.Context, string, string) (*Search, error)
	DeleteDynamicSearch(context.Context, string) error
	GetDynamicSearchResult(context.Context, string) (*SearchResult, error)
}

// Forensics endpoint.
type Forensics interface {
	ListRecoveries(context.Context, string, string, int, int) (*RecoveriesPaginatedResponse, error)
	CreateRecovery(context.Context, map[string]string, string) (*Recovery, error)
	GetRecovery(context.Context, int, string) (*Recovery, error)
	ListRecoveryTasks(context.Context, string, string, int, int) (*RecoveryTasksPaginatedResponse, error)
	GetRecoveryTask(context.Context, int, string) (*RecoveryTask, error)
	GetCaseCreatetask(context.Context, int, string) (*CaseCreateTask, error)
	ListCases(context.Context, string, string, int, int) (*CasesPaginatedResponse, error)
	CreateCase(context.Context, map[string]string, string) (*CreateCase, error)
	GetCase(context.Context, int, string) (*Case, error)
}

// GUIAppFramework endpoint.
type GUIAppFramework interface {
	ListStatusAppInstalls(context.Context, string, string, int, int) (*StatusAppInstallsPaginatedResponse, error)
	CreateAppFramework(context.Context, string, string) (*CreatedAppFramework, error)
	GetCreatedAppFramework(context.Context, int, string) (*CreatedAppFramework, error)
	CancelCreatedAppFramework(context.Context, int, string) (*CreatedAppFramework, error)
	GetAuthRequest(context.Context, int, string) (*AuthRequest, error)
	UpdateAuthRequestResponse(context.Context, int, map[string]string, string) (*AuthRequestResponse, error)
	ListAppDefinitions(context.Context, string, string, int, int) (*AppDefinitionsPaginatedResponse, error)
	CreateAppDefinition(context.Context, string, string) (*AppDefinitionStatus, error)
	GetAppDefinition(context.Context, int, string) (*AppDefinition, error)
	CancelAppDefinition(context.Context, int, string) (*AppDefinitionStatus, error)
	DeleteAppDefinition(context.Context, int) error
	UpdateAppDefinition(context.Context, int, string, string) (*AppDefinitionStatus, error)
	ListUserRoleIds(context.Context, int, string, int, int) (*UserRoleIDsPaginatedResponse, error)
	CreateUserRoleID(context.Context, int, int, string) (*UserRoleID, error)
	DeleteUserRoles(context.Context, int, int, string) (*UserRoleID, error)
	ListInstalledApp(context.Context, int, string, string, int, int) (*InstalledAppsPaginatedResponse, error)
	CreateApplication(context.Context, int, int, string, bool) (*InstalledApp, error)
	GetinstalledApp(context.Context, int, string) (*InstalledApp, error)
	UpdateInstalledApp(context.Context, int, int, string, string, string) (*InstalledApp, error)
	DeleteAppInstance(context.Context, int) error
	UpdateApplication(context.Context, int, string, string) (*CreatedAppFramework, error)
	ListRegisteredServices(context.Context, int, int) (*RegisteredServicesPaginatedResponse, error)
	GetRegisteredServices(context.Context, int) (*RegisteredService, error)
}

// Health endpoint.
type Health interface {
	ListQRadarmetrics(context.Context, string, string, int, int) (*QRadarmetricsPaginatedResponse, error)
	GetQRadarmetric(context.Context, int, string) (*QRadarmetric, error)
	UpdateQRadarmetric(context.Context, int, map[string]string, string) (*QRadarmetric, error)
	UpdateQRadarMetricGC(context.Context, map[string]interface{}, string) (*QRadarMetricGC, error)
	ListSystemMetrics(context.Context, string, string, int, int) (*SystemMetricsPaginatedResponse, error)
	GetSystemMetric(context.Context, int, string) (*SystemMetric, error)
	UpdateSystemMetric(context.Context, int, map[string]interface{}, string) (*SystemMetric, error)
	UpdateSystemMetricGC(context.Context, map[string]interface{}, string) (*QRadarMetricGC, error)
}

// HealthData endpoint.
type HealthData interface {
	GetSecurityDataCount(context.Context, string) (*SecurityDataCount, error)
	ListTopOffenses(context.Context, string, string, int, int) (*TopOffensesPaginatedResponse, error)
	ListTopRules(context.Context, string, string, int, int) (*TopRulesPaginatedResponse, error)
}

// Help endpoint.
type Help interface {
	ListEndpointDocumentationObjects(context.Context, string, string, int, int) (*EndpointDocumentationObjectsPaginatedResponse, error)
	GetEndpointDocumentationObject(context.Context, int, string) (*EndpointDocumentationObject, error)
	ListResourceDocumentationObjects(context.Context, string, string, int, int) (*ResourceDocumentationObjectsPaginatedResponse, error)
	GetResourceDocumentationObject(context.Context, int, string) (*ResourceDocumentationObject, error)
	ListVersionDocumentationObjects(context.Context, string, string, int, int) (*VersionDocumentationObjectsPaginatedResponse, error)
	GetVersionDocumentationObject(context.Context, int, string) (*VersionDocumentationObject, error)
}

// Qni endpoint.
type Qni interface{}

// Qrm endpoint.
type Qrm interface{}

// Qvm endpoint.
type Qvm interface{}

// ReferenceData endpoint.
type ReferenceData interface {
	UpdateBulkLoadRM(context.Context, string, map[string]string, string) (*BulkMap, error)
	DeleteReferenceMap(context.Context, string, string, string, bool) error
	ListSets(context.Context, string, string, int, int) (*ListSetsPaginatedResponse, error)
	UpdateBulkLoadRS(context.Context, string, []string, string) (*Set, error)
	DeleteReferenceSet(context.Context, string, string, string, bool) error
	UpdateBulkLoadRT(context.Context, string, string, map[string]map[string]string) (*BulkTable, error)
	DeleteReferenceTable(context.Context, string, string, string, bool) error
	UpdateBulkLoadRMM(context.Context, string, map[string]map[string]string, string) (*BulkMapOfMap, error)
	DeleteReferenceMapOfMap(context.Context, string, string, string, bool) error
}

// Scanner endpoint.
type Scanner interface{}

// Services endpoint.
type Services interface{}

// SIEM endpoint.
type SIEM interface {
	ListOffenses(context.Context, string, string, string, int, int) (*OffensePaginatedResponse, error)
	GetOffense(context.Context, int, string) (*Offense, error)
	UpdateOffense(context.Context, int, int, string, string, string, bool, bool) (*Offense, error)
	ListOffenseNotes(context.Context, string) ([]*Note, int, error)
	CreateOffenseNote(context.Context, int, string, string) (*Note, error)
	ListOffenseTypes(context.Context, string, string, string, int, int) (*OffenseTypesPaginatedResponse, error)
	GetOffenseType(context.Context, string, string) (*OffenseType, error)
	ListLocalDestinationAddress(context.Context, string, string, int, int) (*LocalDestinationAddressesPaginatedResponse, error)
	GetLocalDestinationAddress(context.Context, int, string) (*LocalDestinationAddress, error)
	ListSourceAddresses(context.Context, string, string, int, int) (*SourceAddressesPaginatedResponse, error)
	GetSourceAddress(context.Context, int, string) (*SourceAddress, error)
	ListOffenseClosingReasons(context.Context, string, string, bool, bool, int, int) (*OffenseClosingReasonsPaginatedResponse, error)
	CreateOffenseClosingReason(context.Context, string, string) (*OffenseClosingReason, error)
	GetOffenseClosingReason(context.Context, int, string) (*OffenseClosingReason, error)
}

// StagedConfig endpoint.
type StagedConfig interface{}

// System endpoint.
type System interface{}
