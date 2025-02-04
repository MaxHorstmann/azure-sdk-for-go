//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlflexibleservers

// AdministratorsClientCreateResponse contains the response from method AdministratorsClient.BeginCreate.
type AdministratorsClientCreateResponse struct {
	// Represents an Active Directory administrator.
	ActiveDirectoryAdministrator
}

// AdministratorsClientDeleteResponse contains the response from method AdministratorsClient.BeginDelete.
type AdministratorsClientDeleteResponse struct {
	// placeholder for future response values
}

// AdministratorsClientGetResponse contains the response from method AdministratorsClient.Get.
type AdministratorsClientGetResponse struct {
	// Represents an Active Directory administrator.
	ActiveDirectoryAdministrator
}

// AdministratorsClientListByServerResponse contains the response from method AdministratorsClient.NewListByServerPager.
type AdministratorsClientListByServerResponse struct {
	// A list of active directory administrators.
	AdministratorListResult
}

// BackupsClientGetResponse contains the response from method BackupsClient.Get.
type BackupsClientGetResponse struct {
	// Server backup properties
	ServerBackup
}

// BackupsClientListByServerResponse contains the response from method BackupsClient.NewListByServerPager.
type BackupsClientListByServerResponse struct {
	// A list of server backups.
	ServerBackupListResult
}

// CheckNameAvailabilityClientExecuteResponse contains the response from method CheckNameAvailabilityClient.Execute.
type CheckNameAvailabilityClientExecuteResponse struct {
	// Represents a resource name availability.
	NameAvailability
}

// CheckNameAvailabilityWithLocationClientExecuteResponse contains the response from method CheckNameAvailabilityWithLocationClient.Execute.
type CheckNameAvailabilityWithLocationClientExecuteResponse struct {
	// Represents a resource name availability.
	NameAvailability
}

// ConfigurationsClientGetResponse contains the response from method ConfigurationsClient.Get.
type ConfigurationsClientGetResponse struct {
	// Represents a Configuration.
	Configuration
}

// ConfigurationsClientListByServerResponse contains the response from method ConfigurationsClient.NewListByServerPager.
type ConfigurationsClientListByServerResponse struct {
	// A list of server configurations.
	ConfigurationListResult
}

// ConfigurationsClientPutResponse contains the response from method ConfigurationsClient.BeginPut.
type ConfigurationsClientPutResponse struct {
	// Represents a Configuration.
	Configuration
}

// ConfigurationsClientUpdateResponse contains the response from method ConfigurationsClient.BeginUpdate.
type ConfigurationsClientUpdateResponse struct {
	// Represents a Configuration.
	Configuration
}

// DatabasesClientCreateResponse contains the response from method DatabasesClient.BeginCreate.
type DatabasesClientCreateResponse struct {
	// Represents a Database.
	Database
}

// DatabasesClientDeleteResponse contains the response from method DatabasesClient.BeginDelete.
type DatabasesClientDeleteResponse struct {
	// placeholder for future response values
}

// DatabasesClientGetResponse contains the response from method DatabasesClient.Get.
type DatabasesClientGetResponse struct {
	// Represents a Database.
	Database
}

// DatabasesClientListByServerResponse contains the response from method DatabasesClient.NewListByServerPager.
type DatabasesClientListByServerResponse struct {
	// A List of databases.
	DatabaseListResult
}

// FirewallRulesClientCreateOrUpdateResponse contains the response from method FirewallRulesClient.BeginCreateOrUpdate.
type FirewallRulesClientCreateOrUpdateResponse struct {
	// Represents a server firewall rule.
	FirewallRule
}

// FirewallRulesClientDeleteResponse contains the response from method FirewallRulesClient.BeginDelete.
type FirewallRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// FirewallRulesClientGetResponse contains the response from method FirewallRulesClient.Get.
type FirewallRulesClientGetResponse struct {
	// Represents a server firewall rule.
	FirewallRule
}

// FirewallRulesClientListByServerResponse contains the response from method FirewallRulesClient.NewListByServerPager.
type FirewallRulesClientListByServerResponse struct {
	// A list of firewall rules.
	FirewallRuleListResult
}

// GetPrivateDNSZoneSuffixClientExecuteResponse contains the response from method GetPrivateDNSZoneSuffixClient.Execute.
type GetPrivateDNSZoneSuffixClientExecuteResponse struct {
	// Represents a resource name availability.
	Value *string
}

// LocationBasedCapabilitiesClientExecuteResponse contains the response from method LocationBasedCapabilitiesClient.NewExecutePager.
type LocationBasedCapabilitiesClientExecuteResponse struct {
	// location capability
	CapabilitiesListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	// A list of resource provider operations.
	OperationListResult
}

// ReplicasClientListByServerResponse contains the response from method ReplicasClient.NewListByServerPager.
type ReplicasClientListByServerResponse struct {
	// A list of servers.
	ServerListResult
}

// ServersClientCreateResponse contains the response from method ServersClient.BeginCreate.
type ServersClientCreateResponse struct {
	// Represents a server.
	Server
}

// ServersClientDeleteResponse contains the response from method ServersClient.BeginDelete.
type ServersClientDeleteResponse struct {
	// placeholder for future response values
}

// ServersClientGetResponse contains the response from method ServersClient.Get.
type ServersClientGetResponse struct {
	// Represents a server.
	Server
}

// ServersClientListByResourceGroupResponse contains the response from method ServersClient.NewListByResourceGroupPager.
type ServersClientListByResourceGroupResponse struct {
	// A list of servers.
	ServerListResult
}

// ServersClientListResponse contains the response from method ServersClient.NewListPager.
type ServersClientListResponse struct {
	// A list of servers.
	ServerListResult
}

// ServersClientRestartResponse contains the response from method ServersClient.BeginRestart.
type ServersClientRestartResponse struct {
	// placeholder for future response values
}

// ServersClientStartResponse contains the response from method ServersClient.BeginStart.
type ServersClientStartResponse struct {
	// placeholder for future response values
}

// ServersClientStopResponse contains the response from method ServersClient.BeginStop.
type ServersClientStopResponse struct {
	// placeholder for future response values
}

// ServersClientUpdateResponse contains the response from method ServersClient.BeginUpdate.
type ServersClientUpdateResponse struct {
	// Represents a server.
	Server
}

// VirtualNetworkSubnetUsageClientExecuteResponse contains the response from method VirtualNetworkSubnetUsageClient.Execute.
type VirtualNetworkSubnetUsageClientExecuteResponse struct {
	// Virtual network subnet usage data.
	VirtualNetworkSubnetUsageResult
}
