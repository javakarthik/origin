// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package recoveryservices

import original "github.com/Azure/azure-sdk-for-go/services/recoveryservices/mgmt/2016-06-01/recoveryservices"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type AuthType = original.AuthType

const (
	AAD                  AuthType = original.AAD
	AccessControlService AuthType = original.AccessControlService
	ACS                  AuthType = original.ACS
	AzureActiveDirectory AuthType = original.AzureActiveDirectory
	Invalid              AuthType = original.Invalid
)

type AuthTypeBasicResourceCertificateDetails = original.AuthTypeBasicResourceCertificateDetails

const (
	AuthTypeAccessControlService       AuthTypeBasicResourceCertificateDetails = original.AuthTypeAccessControlService
	AuthTypeAzureActiveDirectory       AuthTypeBasicResourceCertificateDetails = original.AuthTypeAzureActiveDirectory
	AuthTypeResourceCertificateDetails AuthTypeBasicResourceCertificateDetails = original.AuthTypeResourceCertificateDetails
)

type SkuName = original.SkuName

const (
	RS0      SkuName = original.RS0
	Standard SkuName = original.Standard
)

type TriggerType = original.TriggerType

const (
	ForcedUpgrade TriggerType = original.ForcedUpgrade
	UserTriggered TriggerType = original.UserTriggered
)

type UsagesUnit = original.UsagesUnit

const (
	Bytes          UsagesUnit = original.Bytes
	BytesPerSecond UsagesUnit = original.BytesPerSecond
	Count          UsagesUnit = original.Count
	CountPerSecond UsagesUnit = original.CountPerSecond
	Percent        UsagesUnit = original.Percent
	Seconds        UsagesUnit = original.Seconds
)

type VaultUpgradeState = original.VaultUpgradeState

const (
	Failed     VaultUpgradeState = original.Failed
	InProgress VaultUpgradeState = original.InProgress
	Unknown    VaultUpgradeState = original.Unknown
	Upgraded   VaultUpgradeState = original.Upgraded
)

type CertificateRequest = original.CertificateRequest
type ClientDiscoveryDisplay = original.ClientDiscoveryDisplay
type ClientDiscoveryForLogSpecification = original.ClientDiscoveryForLogSpecification
type ClientDiscoveryForProperties = original.ClientDiscoveryForProperties
type ClientDiscoveryForServiceSpecification = original.ClientDiscoveryForServiceSpecification
type ClientDiscoveryResponse = original.ClientDiscoveryResponse
type ClientDiscoveryResponseIterator = original.ClientDiscoveryResponseIterator
type ClientDiscoveryResponsePage = original.ClientDiscoveryResponsePage
type ClientDiscoveryValueForSingleAPI = original.ClientDiscoveryValueForSingleAPI
type JobsSummary = original.JobsSummary
type MonitoringSummary = original.MonitoringSummary
type NameInfo = original.NameInfo
type PatchTrackedResource = original.PatchTrackedResource
type PatchVault = original.PatchVault
type RawCertificateData = original.RawCertificateData
type ReplicationUsage = original.ReplicationUsage
type ReplicationUsageList = original.ReplicationUsageList
type Resource = original.Resource
type ResourceCertificateAndAadDetails = original.ResourceCertificateAndAadDetails
type ResourceCertificateAndAcsDetails = original.ResourceCertificateAndAcsDetails
type BasicResourceCertificateDetails = original.BasicResourceCertificateDetails
type ResourceCertificateDetails = original.ResourceCertificateDetails
type Sku = original.Sku
type TrackedResource = original.TrackedResource
type UpgradeDetails = original.UpgradeDetails
type Vault = original.Vault
type VaultCertificateResponse = original.VaultCertificateResponse
type VaultExtendedInfo = original.VaultExtendedInfo
type VaultExtendedInfoResource = original.VaultExtendedInfoResource
type VaultList = original.VaultList
type VaultListIterator = original.VaultListIterator
type VaultListPage = original.VaultListPage
type VaultProperties = original.VaultProperties
type VaultUsage = original.VaultUsage
type VaultUsageList = original.VaultUsageList
type OperationsClient = original.OperationsClient
type RegisteredIdentitiesClient = original.RegisteredIdentitiesClient
type ReplicationUsagesClient = original.ReplicationUsagesClient
type UsagesClient = original.UsagesClient
type VaultCertificatesClient = original.VaultCertificatesClient
type VaultExtendedInfoClient = original.VaultExtendedInfoClient
type VaultsClient = original.VaultsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAuthTypeValues() []AuthType {
	return original.PossibleAuthTypeValues()
}
func PossibleAuthTypeBasicResourceCertificateDetailsValues() []AuthTypeBasicResourceCertificateDetails {
	return original.PossibleAuthTypeBasicResourceCertificateDetailsValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleTriggerTypeValues() []TriggerType {
	return original.PossibleTriggerTypeValues()
}
func PossibleUsagesUnitValues() []UsagesUnit {
	return original.PossibleUsagesUnitValues()
}
func PossibleVaultUpgradeStateValues() []VaultUpgradeState {
	return original.PossibleVaultUpgradeStateValues()
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRegisteredIdentitiesClient(subscriptionID string) RegisteredIdentitiesClient {
	return original.NewRegisteredIdentitiesClient(subscriptionID)
}
func NewRegisteredIdentitiesClientWithBaseURI(baseURI string, subscriptionID string) RegisteredIdentitiesClient {
	return original.NewRegisteredIdentitiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewReplicationUsagesClient(subscriptionID string) ReplicationUsagesClient {
	return original.NewReplicationUsagesClient(subscriptionID)
}
func NewReplicationUsagesClientWithBaseURI(baseURI string, subscriptionID string) ReplicationUsagesClient {
	return original.NewReplicationUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVaultCertificatesClient(subscriptionID string) VaultCertificatesClient {
	return original.NewVaultCertificatesClient(subscriptionID)
}
func NewVaultCertificatesClientWithBaseURI(baseURI string, subscriptionID string) VaultCertificatesClient {
	return original.NewVaultCertificatesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVaultExtendedInfoClient(subscriptionID string) VaultExtendedInfoClient {
	return original.NewVaultExtendedInfoClient(subscriptionID)
}
func NewVaultExtendedInfoClientWithBaseURI(baseURI string, subscriptionID string) VaultExtendedInfoClient {
	return original.NewVaultExtendedInfoClientWithBaseURI(baseURI, subscriptionID)
}
func NewVaultsClient(subscriptionID string) VaultsClient {
	return original.NewVaultsClient(subscriptionID)
}
func NewVaultsClientWithBaseURI(baseURI string, subscriptionID string) VaultsClient {
	return original.NewVaultsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
