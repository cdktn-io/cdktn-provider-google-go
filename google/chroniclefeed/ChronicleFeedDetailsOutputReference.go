// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/chroniclefeed/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleFeedDetailsOutputReference interface {
	cdktn.ComplexObject
	AmazonKinesisFirehoseSettings() ChronicleFeedDetailsAmazonKinesisFirehoseSettingsOutputReference
	AmazonKinesisFirehoseSettingsInput() *ChronicleFeedDetailsAmazonKinesisFirehoseSettings
	AmazonS3Settings() ChronicleFeedDetailsAmazonS3SettingsOutputReference
	AmazonS3SettingsInput() *ChronicleFeedDetailsAmazonS3Settings
	AmazonS3V2Settings() ChronicleFeedDetailsAmazonS3V2SettingsOutputReference
	AmazonS3V2SettingsInput() *ChronicleFeedDetailsAmazonS3V2Settings
	AmazonSqsSettings() ChronicleFeedDetailsAmazonSqsSettingsOutputReference
	AmazonSqsSettingsInput() *ChronicleFeedDetailsAmazonSqsSettings
	AmazonSqsV2Settings() ChronicleFeedDetailsAmazonSqsV2SettingsOutputReference
	AmazonSqsV2SettingsInput() *ChronicleFeedDetailsAmazonSqsV2Settings
	AnomaliSettings() ChronicleFeedDetailsAnomaliSettingsOutputReference
	AnomaliSettingsInput() *ChronicleFeedDetailsAnomaliSettings
	AssetNamespace() *string
	SetAssetNamespace(val *string)
	AssetNamespaceInput() *string
	AwsEc2HostsSettings() ChronicleFeedDetailsAwsEc2HostsSettingsOutputReference
	AwsEc2HostsSettingsInput() *ChronicleFeedDetailsAwsEc2HostsSettings
	AwsEc2InstancesSettings() ChronicleFeedDetailsAwsEc2InstancesSettingsOutputReference
	AwsEc2InstancesSettingsInput() *ChronicleFeedDetailsAwsEc2InstancesSettings
	AwsEc2VpcsSettings() ChronicleFeedDetailsAwsEc2VpcsSettingsOutputReference
	AwsEc2VpcsSettingsInput() *ChronicleFeedDetailsAwsEc2VpcsSettings
	AwsIamSettings() ChronicleFeedDetailsAwsIamSettingsOutputReference
	AwsIamSettingsInput() *ChronicleFeedDetailsAwsIamSettings
	AzureAdAuditSettings() ChronicleFeedDetailsAzureAdAuditSettingsOutputReference
	AzureAdAuditSettingsInput() *ChronicleFeedDetailsAzureAdAuditSettings
	AzureAdContextSettings() ChronicleFeedDetailsAzureAdContextSettingsOutputReference
	AzureAdContextSettingsInput() *ChronicleFeedDetailsAzureAdContextSettings
	AzureAdSettings() ChronicleFeedDetailsAzureAdSettingsOutputReference
	AzureAdSettingsInput() *ChronicleFeedDetailsAzureAdSettings
	AzureBlobStoreSettings() ChronicleFeedDetailsAzureBlobStoreSettingsOutputReference
	AzureBlobStoreSettingsInput() *ChronicleFeedDetailsAzureBlobStoreSettings
	AzureBlobStoreV2Settings() ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference
	AzureBlobStoreV2SettingsInput() *ChronicleFeedDetailsAzureBlobStoreV2Settings
	AzureEventHubSettings() ChronicleFeedDetailsAzureEventHubSettingsOutputReference
	AzureEventHubSettingsInput() *ChronicleFeedDetailsAzureEventHubSettings
	AzureMdmIntuneSettings() ChronicleFeedDetailsAzureMdmIntuneSettingsOutputReference
	AzureMdmIntuneSettingsInput() *ChronicleFeedDetailsAzureMdmIntuneSettings
	CloudPassageSettings() ChronicleFeedDetailsCloudPassageSettingsOutputReference
	CloudPassageSettingsInput() *ChronicleFeedDetailsCloudPassageSettings
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	CortexXdrSettings() ChronicleFeedDetailsCortexXdrSettingsOutputReference
	CortexXdrSettingsInput() *ChronicleFeedDetailsCortexXdrSettings
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CrowdstrikeAlertsSettings() ChronicleFeedDetailsCrowdstrikeAlertsSettingsOutputReference
	CrowdstrikeAlertsSettingsInput() *ChronicleFeedDetailsCrowdstrikeAlertsSettings
	CrowdstrikeDetectsSettings() ChronicleFeedDetailsCrowdstrikeDetectsSettingsOutputReference
	CrowdstrikeDetectsSettingsInput() *ChronicleFeedDetailsCrowdstrikeDetectsSettings
	DummyLogTypeSettings() ChronicleFeedDetailsDummyLogTypeSettingsOutputReference
	DummyLogTypeSettingsInput() *ChronicleFeedDetailsDummyLogTypeSettings
	DuoAuthSettings() ChronicleFeedDetailsDuoAuthSettingsOutputReference
	DuoAuthSettingsInput() *ChronicleFeedDetailsDuoAuthSettings
	DuoUserContextSettings() ChronicleFeedDetailsDuoUserContextSettingsOutputReference
	DuoUserContextSettingsInput() *ChronicleFeedDetailsDuoUserContextSettings
	FeedSourceType() *string
	SetFeedSourceType(val *string)
	FeedSourceTypeInput() *string
	FoxItStixSettings() ChronicleFeedDetailsFoxItStixSettingsOutputReference
	FoxItStixSettingsInput() *ChronicleFeedDetailsFoxItStixSettings
	// Experimental.
	Fqn() *string
	GcsSettings() ChronicleFeedDetailsGcsSettingsOutputReference
	GcsSettingsInput() *ChronicleFeedDetailsGcsSettings
	GcsV2Settings() ChronicleFeedDetailsGcsV2SettingsOutputReference
	GcsV2SettingsInput() *ChronicleFeedDetailsGcsV2Settings
	GoogleCloudIdentityDevicesSettings() ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsOutputReference
	GoogleCloudIdentityDevicesSettingsInput() *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettings
	GoogleCloudIdentityDeviceUsersSettings() ChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettingsOutputReference
	GoogleCloudIdentityDeviceUsersSettingsInput() *ChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettings
	GoogleCloudStorageEventDrivenSettings() ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference
	GoogleCloudStorageEventDrivenSettingsInput() *ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings
	HttpSettings() ChronicleFeedDetailsHttpSettingsOutputReference
	HttpSettingsInput() *ChronicleFeedDetailsHttpSettings
	HttpsPushAmazonKinesisFirehoseSettings() ChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettingsOutputReference
	HttpsPushAmazonKinesisFirehoseSettingsInput() *ChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettings
	HttpsPushGoogleCloudPubsubSettings() ChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettingsOutputReference
	HttpsPushGoogleCloudPubsubSettingsInput() *ChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettings
	HttpsPushWebhookSettings() ChronicleFeedDetailsHttpsPushWebhookSettingsOutputReference
	HttpsPushWebhookSettingsInput() *ChronicleFeedDetailsHttpsPushWebhookSettings
	ImpervaWafSettings() ChronicleFeedDetailsImpervaWafSettingsOutputReference
	ImpervaWafSettingsInput() *ChronicleFeedDetailsImpervaWafSettings
	InternalValue() *ChronicleFeedDetails
	SetInternalValue(val *ChronicleFeedDetails)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LogType() *string
	SetLogType(val *string)
	LogTypeInput() *string
	MandiantIocSettings() ChronicleFeedDetailsMandiantIocSettingsOutputReference
	MandiantIocSettingsInput() *ChronicleFeedDetailsMandiantIocSettings
	MicrosoftGraphAlertSettings() ChronicleFeedDetailsMicrosoftGraphAlertSettingsOutputReference
	MicrosoftGraphAlertSettingsInput() *ChronicleFeedDetailsMicrosoftGraphAlertSettings
	MicrosoftSecurityCenterAlertSettings() ChronicleFeedDetailsMicrosoftSecurityCenterAlertSettingsOutputReference
	MicrosoftSecurityCenterAlertSettingsInput() *ChronicleFeedDetailsMicrosoftSecurityCenterAlertSettings
	MimecastMailSettings() ChronicleFeedDetailsMimecastMailSettingsOutputReference
	MimecastMailSettingsInput() *ChronicleFeedDetailsMimecastMailSettings
	MimecastMailV2Settings() ChronicleFeedDetailsMimecastMailV2SettingsOutputReference
	MimecastMailV2SettingsInput() *ChronicleFeedDetailsMimecastMailV2Settings
	NetskopeAlertSettings() ChronicleFeedDetailsNetskopeAlertSettingsOutputReference
	NetskopeAlertSettingsInput() *ChronicleFeedDetailsNetskopeAlertSettings
	NetskopeAlertV2Settings() ChronicleFeedDetailsNetskopeAlertV2SettingsOutputReference
	NetskopeAlertV2SettingsInput() *ChronicleFeedDetailsNetskopeAlertV2Settings
	Office365Settings() ChronicleFeedDetailsOffice365SettingsOutputReference
	Office365SettingsInput() *ChronicleFeedDetailsOffice365Settings
	OktaSettings() ChronicleFeedDetailsOktaSettingsOutputReference
	OktaSettingsInput() *ChronicleFeedDetailsOktaSettings
	OktaUserContextSettings() ChronicleFeedDetailsOktaUserContextSettingsOutputReference
	OktaUserContextSettingsInput() *ChronicleFeedDetailsOktaUserContextSettings
	PanIocSettings() ChronicleFeedDetailsPanIocSettingsOutputReference
	PanIocSettingsInput() *ChronicleFeedDetailsPanIocSettings
	PanPrismaCloudSettings() ChronicleFeedDetailsPanPrismaCloudSettingsOutputReference
	PanPrismaCloudSettingsInput() *ChronicleFeedDetailsPanPrismaCloudSettings
	ProofpointMailSettings() ChronicleFeedDetailsProofpointMailSettingsOutputReference
	ProofpointMailSettingsInput() *ChronicleFeedDetailsProofpointMailSettings
	ProofpointOnDemandSettings() ChronicleFeedDetailsProofpointOnDemandSettingsOutputReference
	ProofpointOnDemandSettingsInput() *ChronicleFeedDetailsProofpointOnDemandSettings
	PubsubSettings() ChronicleFeedDetailsPubsubSettingsOutputReference
	PubsubSettingsInput() *ChronicleFeedDetailsPubsubSettings
	QualysScanSettings() ChronicleFeedDetailsQualysScanSettingsOutputReference
	QualysScanSettingsInput() *ChronicleFeedDetailsQualysScanSettings
	QualysVmSettings() ChronicleFeedDetailsQualysVmSettingsOutputReference
	QualysVmSettingsInput() *ChronicleFeedDetailsQualysVmSettings
	Rapid7InsightSettings() ChronicleFeedDetailsRapid7InsightSettingsOutputReference
	Rapid7InsightSettingsInput() *ChronicleFeedDetailsRapid7InsightSettings
	RecordedFutureIocSettings() ChronicleFeedDetailsRecordedFutureIocSettingsOutputReference
	RecordedFutureIocSettingsInput() *ChronicleFeedDetailsRecordedFutureIocSettings
	RhIsacIocSettings() ChronicleFeedDetailsRhIsacIocSettingsOutputReference
	RhIsacIocSettingsInput() *ChronicleFeedDetailsRhIsacIocSettings
	SalesforceSettings() ChronicleFeedDetailsSalesforceSettingsOutputReference
	SalesforceSettingsInput() *ChronicleFeedDetailsSalesforceSettings
	SentineloneAlertSettings() ChronicleFeedDetailsSentineloneAlertSettingsOutputReference
	SentineloneAlertSettingsInput() *ChronicleFeedDetailsSentineloneAlertSettings
	ServiceNowCmdbSettings() ChronicleFeedDetailsServiceNowCmdbSettingsOutputReference
	ServiceNowCmdbSettingsInput() *ChronicleFeedDetailsServiceNowCmdbSettings
	SftpSettings() ChronicleFeedDetailsSftpSettingsOutputReference
	SftpSettingsInput() *ChronicleFeedDetailsSftpSettings
	StsMigrationReadiness() *string
	SymantecEventExportSettings() ChronicleFeedDetailsSymantecEventExportSettingsOutputReference
	SymantecEventExportSettingsInput() *ChronicleFeedDetailsSymantecEventExportSettings
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThinkstCanarySettings() ChronicleFeedDetailsThinkstCanarySettingsOutputReference
	ThinkstCanarySettingsInput() *ChronicleFeedDetailsThinkstCanarySettings
	ThreatConnectIocSettings() ChronicleFeedDetailsThreatConnectIocSettingsOutputReference
	ThreatConnectIocSettingsInput() *ChronicleFeedDetailsThreatConnectIocSettings
	ThreatConnectIocV3Settings() ChronicleFeedDetailsThreatConnectIocV3SettingsOutputReference
	ThreatConnectIocV3SettingsInput() *ChronicleFeedDetailsThreatConnectIocV3Settings
	TrellixHxAlertsSettings() ChronicleFeedDetailsTrellixHxAlertsSettingsOutputReference
	TrellixHxAlertsSettingsInput() *ChronicleFeedDetailsTrellixHxAlertsSettings
	TrellixHxBulkAcqsSettings() ChronicleFeedDetailsTrellixHxBulkAcqsSettingsOutputReference
	TrellixHxBulkAcqsSettingsInput() *ChronicleFeedDetailsTrellixHxBulkAcqsSettings
	TrellixHxHostsSettings() ChronicleFeedDetailsTrellixHxHostsSettingsOutputReference
	TrellixHxHostsSettingsInput() *ChronicleFeedDetailsTrellixHxHostsSettings
	WebhookSettings() ChronicleFeedDetailsWebhookSettingsOutputReference
	WebhookSettingsInput() *ChronicleFeedDetailsWebhookSettings
	WorkdaySettings() ChronicleFeedDetailsWorkdaySettingsOutputReference
	WorkdaySettingsInput() *ChronicleFeedDetailsWorkdaySettings
	WorkspaceActivitySettings() ChronicleFeedDetailsWorkspaceActivitySettingsOutputReference
	WorkspaceActivitySettingsInput() *ChronicleFeedDetailsWorkspaceActivitySettings
	WorkspaceAlertsSettings() ChronicleFeedDetailsWorkspaceAlertsSettingsOutputReference
	WorkspaceAlertsSettingsInput() *ChronicleFeedDetailsWorkspaceAlertsSettings
	WorkspaceChromeOsSettings() ChronicleFeedDetailsWorkspaceChromeOsSettingsOutputReference
	WorkspaceChromeOsSettingsInput() *ChronicleFeedDetailsWorkspaceChromeOsSettings
	WorkspaceGroupsSettings() ChronicleFeedDetailsWorkspaceGroupsSettingsOutputReference
	WorkspaceGroupsSettingsInput() *ChronicleFeedDetailsWorkspaceGroupsSettings
	WorkspaceMobileSettings() ChronicleFeedDetailsWorkspaceMobileSettingsOutputReference
	WorkspaceMobileSettingsInput() *ChronicleFeedDetailsWorkspaceMobileSettings
	WorkspacePrivilegesSettings() ChronicleFeedDetailsWorkspacePrivilegesSettingsOutputReference
	WorkspacePrivilegesSettingsInput() *ChronicleFeedDetailsWorkspacePrivilegesSettings
	WorkspaceUsersSettings() ChronicleFeedDetailsWorkspaceUsersSettingsOutputReference
	WorkspaceUsersSettingsInput() *ChronicleFeedDetailsWorkspaceUsersSettings
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutAmazonKinesisFirehoseSettings(value *ChronicleFeedDetailsAmazonKinesisFirehoseSettings)
	PutAmazonS3Settings(value *ChronicleFeedDetailsAmazonS3Settings)
	PutAmazonS3V2Settings(value *ChronicleFeedDetailsAmazonS3V2Settings)
	PutAmazonSqsSettings(value *ChronicleFeedDetailsAmazonSqsSettings)
	PutAmazonSqsV2Settings(value *ChronicleFeedDetailsAmazonSqsV2Settings)
	PutAnomaliSettings(value *ChronicleFeedDetailsAnomaliSettings)
	PutAwsEc2HostsSettings(value *ChronicleFeedDetailsAwsEc2HostsSettings)
	PutAwsEc2InstancesSettings(value *ChronicleFeedDetailsAwsEc2InstancesSettings)
	PutAwsEc2VpcsSettings(value *ChronicleFeedDetailsAwsEc2VpcsSettings)
	PutAwsIamSettings(value *ChronicleFeedDetailsAwsIamSettings)
	PutAzureAdAuditSettings(value *ChronicleFeedDetailsAzureAdAuditSettings)
	PutAzureAdContextSettings(value *ChronicleFeedDetailsAzureAdContextSettings)
	PutAzureAdSettings(value *ChronicleFeedDetailsAzureAdSettings)
	PutAzureBlobStoreSettings(value *ChronicleFeedDetailsAzureBlobStoreSettings)
	PutAzureBlobStoreV2Settings(value *ChronicleFeedDetailsAzureBlobStoreV2Settings)
	PutAzureEventHubSettings(value *ChronicleFeedDetailsAzureEventHubSettings)
	PutAzureMdmIntuneSettings(value *ChronicleFeedDetailsAzureMdmIntuneSettings)
	PutCloudPassageSettings(value *ChronicleFeedDetailsCloudPassageSettings)
	PutCortexXdrSettings(value *ChronicleFeedDetailsCortexXdrSettings)
	PutCrowdstrikeAlertsSettings(value *ChronicleFeedDetailsCrowdstrikeAlertsSettings)
	PutCrowdstrikeDetectsSettings(value *ChronicleFeedDetailsCrowdstrikeDetectsSettings)
	PutDummyLogTypeSettings(value *ChronicleFeedDetailsDummyLogTypeSettings)
	PutDuoAuthSettings(value *ChronicleFeedDetailsDuoAuthSettings)
	PutDuoUserContextSettings(value *ChronicleFeedDetailsDuoUserContextSettings)
	PutFoxItStixSettings(value *ChronicleFeedDetailsFoxItStixSettings)
	PutGcsSettings(value *ChronicleFeedDetailsGcsSettings)
	PutGcsV2Settings(value *ChronicleFeedDetailsGcsV2Settings)
	PutGoogleCloudIdentityDevicesSettings(value *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettings)
	PutGoogleCloudIdentityDeviceUsersSettings(value *ChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettings)
	PutGoogleCloudStorageEventDrivenSettings(value *ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings)
	PutHttpSettings(value *ChronicleFeedDetailsHttpSettings)
	PutHttpsPushAmazonKinesisFirehoseSettings(value *ChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettings)
	PutHttpsPushGoogleCloudPubsubSettings(value *ChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettings)
	PutHttpsPushWebhookSettings(value *ChronicleFeedDetailsHttpsPushWebhookSettings)
	PutImpervaWafSettings(value *ChronicleFeedDetailsImpervaWafSettings)
	PutMandiantIocSettings(value *ChronicleFeedDetailsMandiantIocSettings)
	PutMicrosoftGraphAlertSettings(value *ChronicleFeedDetailsMicrosoftGraphAlertSettings)
	PutMicrosoftSecurityCenterAlertSettings(value *ChronicleFeedDetailsMicrosoftSecurityCenterAlertSettings)
	PutMimecastMailSettings(value *ChronicleFeedDetailsMimecastMailSettings)
	PutMimecastMailV2Settings(value *ChronicleFeedDetailsMimecastMailV2Settings)
	PutNetskopeAlertSettings(value *ChronicleFeedDetailsNetskopeAlertSettings)
	PutNetskopeAlertV2Settings(value *ChronicleFeedDetailsNetskopeAlertV2Settings)
	PutOffice365Settings(value *ChronicleFeedDetailsOffice365Settings)
	PutOktaSettings(value *ChronicleFeedDetailsOktaSettings)
	PutOktaUserContextSettings(value *ChronicleFeedDetailsOktaUserContextSettings)
	PutPanIocSettings(value *ChronicleFeedDetailsPanIocSettings)
	PutPanPrismaCloudSettings(value *ChronicleFeedDetailsPanPrismaCloudSettings)
	PutProofpointMailSettings(value *ChronicleFeedDetailsProofpointMailSettings)
	PutProofpointOnDemandSettings(value *ChronicleFeedDetailsProofpointOnDemandSettings)
	PutPubsubSettings(value *ChronicleFeedDetailsPubsubSettings)
	PutQualysScanSettings(value *ChronicleFeedDetailsQualysScanSettings)
	PutQualysVmSettings(value *ChronicleFeedDetailsQualysVmSettings)
	PutRapid7InsightSettings(value *ChronicleFeedDetailsRapid7InsightSettings)
	PutRecordedFutureIocSettings(value *ChronicleFeedDetailsRecordedFutureIocSettings)
	PutRhIsacIocSettings(value *ChronicleFeedDetailsRhIsacIocSettings)
	PutSalesforceSettings(value *ChronicleFeedDetailsSalesforceSettings)
	PutSentineloneAlertSettings(value *ChronicleFeedDetailsSentineloneAlertSettings)
	PutServiceNowCmdbSettings(value *ChronicleFeedDetailsServiceNowCmdbSettings)
	PutSftpSettings(value *ChronicleFeedDetailsSftpSettings)
	PutSymantecEventExportSettings(value *ChronicleFeedDetailsSymantecEventExportSettings)
	PutThinkstCanarySettings(value *ChronicleFeedDetailsThinkstCanarySettings)
	PutThreatConnectIocSettings(value *ChronicleFeedDetailsThreatConnectIocSettings)
	PutThreatConnectIocV3Settings(value *ChronicleFeedDetailsThreatConnectIocV3Settings)
	PutTrellixHxAlertsSettings(value *ChronicleFeedDetailsTrellixHxAlertsSettings)
	PutTrellixHxBulkAcqsSettings(value *ChronicleFeedDetailsTrellixHxBulkAcqsSettings)
	PutTrellixHxHostsSettings(value *ChronicleFeedDetailsTrellixHxHostsSettings)
	PutWebhookSettings(value *ChronicleFeedDetailsWebhookSettings)
	PutWorkdaySettings(value *ChronicleFeedDetailsWorkdaySettings)
	PutWorkspaceActivitySettings(value *ChronicleFeedDetailsWorkspaceActivitySettings)
	PutWorkspaceAlertsSettings(value *ChronicleFeedDetailsWorkspaceAlertsSettings)
	PutWorkspaceChromeOsSettings(value *ChronicleFeedDetailsWorkspaceChromeOsSettings)
	PutWorkspaceGroupsSettings(value *ChronicleFeedDetailsWorkspaceGroupsSettings)
	PutWorkspaceMobileSettings(value *ChronicleFeedDetailsWorkspaceMobileSettings)
	PutWorkspacePrivilegesSettings(value *ChronicleFeedDetailsWorkspacePrivilegesSettings)
	PutWorkspaceUsersSettings(value *ChronicleFeedDetailsWorkspaceUsersSettings)
	ResetAmazonKinesisFirehoseSettings()
	ResetAmazonS3Settings()
	ResetAmazonS3V2Settings()
	ResetAmazonSqsSettings()
	ResetAmazonSqsV2Settings()
	ResetAnomaliSettings()
	ResetAssetNamespace()
	ResetAwsEc2HostsSettings()
	ResetAwsEc2InstancesSettings()
	ResetAwsEc2VpcsSettings()
	ResetAwsIamSettings()
	ResetAzureAdAuditSettings()
	ResetAzureAdContextSettings()
	ResetAzureAdSettings()
	ResetAzureBlobStoreSettings()
	ResetAzureBlobStoreV2Settings()
	ResetAzureEventHubSettings()
	ResetAzureMdmIntuneSettings()
	ResetCloudPassageSettings()
	ResetCortexXdrSettings()
	ResetCrowdstrikeAlertsSettings()
	ResetCrowdstrikeDetectsSettings()
	ResetDummyLogTypeSettings()
	ResetDuoAuthSettings()
	ResetDuoUserContextSettings()
	ResetFeedSourceType()
	ResetFoxItStixSettings()
	ResetGcsSettings()
	ResetGcsV2Settings()
	ResetGoogleCloudIdentityDevicesSettings()
	ResetGoogleCloudIdentityDeviceUsersSettings()
	ResetGoogleCloudStorageEventDrivenSettings()
	ResetHttpSettings()
	ResetHttpsPushAmazonKinesisFirehoseSettings()
	ResetHttpsPushGoogleCloudPubsubSettings()
	ResetHttpsPushWebhookSettings()
	ResetImpervaWafSettings()
	ResetLabels()
	ResetMandiantIocSettings()
	ResetMicrosoftGraphAlertSettings()
	ResetMicrosoftSecurityCenterAlertSettings()
	ResetMimecastMailSettings()
	ResetMimecastMailV2Settings()
	ResetNetskopeAlertSettings()
	ResetNetskopeAlertV2Settings()
	ResetOffice365Settings()
	ResetOktaSettings()
	ResetOktaUserContextSettings()
	ResetPanIocSettings()
	ResetPanPrismaCloudSettings()
	ResetProofpointMailSettings()
	ResetProofpointOnDemandSettings()
	ResetPubsubSettings()
	ResetQualysScanSettings()
	ResetQualysVmSettings()
	ResetRapid7InsightSettings()
	ResetRecordedFutureIocSettings()
	ResetRhIsacIocSettings()
	ResetSalesforceSettings()
	ResetSentineloneAlertSettings()
	ResetServiceNowCmdbSettings()
	ResetSftpSettings()
	ResetSymantecEventExportSettings()
	ResetThinkstCanarySettings()
	ResetThreatConnectIocSettings()
	ResetThreatConnectIocV3Settings()
	ResetTrellixHxAlertsSettings()
	ResetTrellixHxBulkAcqsSettings()
	ResetTrellixHxHostsSettings()
	ResetWebhookSettings()
	ResetWorkdaySettings()
	ResetWorkspaceActivitySettings()
	ResetWorkspaceAlertsSettings()
	ResetWorkspaceChromeOsSettings()
	ResetWorkspaceGroupsSettings()
	ResetWorkspaceMobileSettings()
	ResetWorkspacePrivilegesSettings()
	ResetWorkspaceUsersSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleFeedDetailsOutputReference
type jsiiProxy_ChronicleFeedDetailsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AmazonKinesisFirehoseSettings() ChronicleFeedDetailsAmazonKinesisFirehoseSettingsOutputReference {
	var returns ChronicleFeedDetailsAmazonKinesisFirehoseSettingsOutputReference
	_jsii_.Get(
		j,
		"amazonKinesisFirehoseSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AmazonKinesisFirehoseSettingsInput() *ChronicleFeedDetailsAmazonKinesisFirehoseSettings {
	var returns *ChronicleFeedDetailsAmazonKinesisFirehoseSettings
	_jsii_.Get(
		j,
		"amazonKinesisFirehoseSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AmazonS3Settings() ChronicleFeedDetailsAmazonS3SettingsOutputReference {
	var returns ChronicleFeedDetailsAmazonS3SettingsOutputReference
	_jsii_.Get(
		j,
		"amazonS3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AmazonS3SettingsInput() *ChronicleFeedDetailsAmazonS3Settings {
	var returns *ChronicleFeedDetailsAmazonS3Settings
	_jsii_.Get(
		j,
		"amazonS3SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AmazonS3V2Settings() ChronicleFeedDetailsAmazonS3V2SettingsOutputReference {
	var returns ChronicleFeedDetailsAmazonS3V2SettingsOutputReference
	_jsii_.Get(
		j,
		"amazonS3V2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AmazonS3V2SettingsInput() *ChronicleFeedDetailsAmazonS3V2Settings {
	var returns *ChronicleFeedDetailsAmazonS3V2Settings
	_jsii_.Get(
		j,
		"amazonS3V2SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AmazonSqsSettings() ChronicleFeedDetailsAmazonSqsSettingsOutputReference {
	var returns ChronicleFeedDetailsAmazonSqsSettingsOutputReference
	_jsii_.Get(
		j,
		"amazonSqsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AmazonSqsSettingsInput() *ChronicleFeedDetailsAmazonSqsSettings {
	var returns *ChronicleFeedDetailsAmazonSqsSettings
	_jsii_.Get(
		j,
		"amazonSqsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AmazonSqsV2Settings() ChronicleFeedDetailsAmazonSqsV2SettingsOutputReference {
	var returns ChronicleFeedDetailsAmazonSqsV2SettingsOutputReference
	_jsii_.Get(
		j,
		"amazonSqsV2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AmazonSqsV2SettingsInput() *ChronicleFeedDetailsAmazonSqsV2Settings {
	var returns *ChronicleFeedDetailsAmazonSqsV2Settings
	_jsii_.Get(
		j,
		"amazonSqsV2SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AnomaliSettings() ChronicleFeedDetailsAnomaliSettingsOutputReference {
	var returns ChronicleFeedDetailsAnomaliSettingsOutputReference
	_jsii_.Get(
		j,
		"anomaliSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AnomaliSettingsInput() *ChronicleFeedDetailsAnomaliSettings {
	var returns *ChronicleFeedDetailsAnomaliSettings
	_jsii_.Get(
		j,
		"anomaliSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AssetNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AssetNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AwsEc2HostsSettings() ChronicleFeedDetailsAwsEc2HostsSettingsOutputReference {
	var returns ChronicleFeedDetailsAwsEc2HostsSettingsOutputReference
	_jsii_.Get(
		j,
		"awsEc2HostsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AwsEc2HostsSettingsInput() *ChronicleFeedDetailsAwsEc2HostsSettings {
	var returns *ChronicleFeedDetailsAwsEc2HostsSettings
	_jsii_.Get(
		j,
		"awsEc2HostsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AwsEc2InstancesSettings() ChronicleFeedDetailsAwsEc2InstancesSettingsOutputReference {
	var returns ChronicleFeedDetailsAwsEc2InstancesSettingsOutputReference
	_jsii_.Get(
		j,
		"awsEc2InstancesSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AwsEc2InstancesSettingsInput() *ChronicleFeedDetailsAwsEc2InstancesSettings {
	var returns *ChronicleFeedDetailsAwsEc2InstancesSettings
	_jsii_.Get(
		j,
		"awsEc2InstancesSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AwsEc2VpcsSettings() ChronicleFeedDetailsAwsEc2VpcsSettingsOutputReference {
	var returns ChronicleFeedDetailsAwsEc2VpcsSettingsOutputReference
	_jsii_.Get(
		j,
		"awsEc2VpcsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AwsEc2VpcsSettingsInput() *ChronicleFeedDetailsAwsEc2VpcsSettings {
	var returns *ChronicleFeedDetailsAwsEc2VpcsSettings
	_jsii_.Get(
		j,
		"awsEc2VpcsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AwsIamSettings() ChronicleFeedDetailsAwsIamSettingsOutputReference {
	var returns ChronicleFeedDetailsAwsIamSettingsOutputReference
	_jsii_.Get(
		j,
		"awsIamSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AwsIamSettingsInput() *ChronicleFeedDetailsAwsIamSettings {
	var returns *ChronicleFeedDetailsAwsIamSettings
	_jsii_.Get(
		j,
		"awsIamSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AzureAdAuditSettings() ChronicleFeedDetailsAzureAdAuditSettingsOutputReference {
	var returns ChronicleFeedDetailsAzureAdAuditSettingsOutputReference
	_jsii_.Get(
		j,
		"azureAdAuditSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AzureAdAuditSettingsInput() *ChronicleFeedDetailsAzureAdAuditSettings {
	var returns *ChronicleFeedDetailsAzureAdAuditSettings
	_jsii_.Get(
		j,
		"azureAdAuditSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AzureAdContextSettings() ChronicleFeedDetailsAzureAdContextSettingsOutputReference {
	var returns ChronicleFeedDetailsAzureAdContextSettingsOutputReference
	_jsii_.Get(
		j,
		"azureAdContextSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AzureAdContextSettingsInput() *ChronicleFeedDetailsAzureAdContextSettings {
	var returns *ChronicleFeedDetailsAzureAdContextSettings
	_jsii_.Get(
		j,
		"azureAdContextSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AzureAdSettings() ChronicleFeedDetailsAzureAdSettingsOutputReference {
	var returns ChronicleFeedDetailsAzureAdSettingsOutputReference
	_jsii_.Get(
		j,
		"azureAdSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AzureAdSettingsInput() *ChronicleFeedDetailsAzureAdSettings {
	var returns *ChronicleFeedDetailsAzureAdSettings
	_jsii_.Get(
		j,
		"azureAdSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AzureBlobStoreSettings() ChronicleFeedDetailsAzureBlobStoreSettingsOutputReference {
	var returns ChronicleFeedDetailsAzureBlobStoreSettingsOutputReference
	_jsii_.Get(
		j,
		"azureBlobStoreSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AzureBlobStoreSettingsInput() *ChronicleFeedDetailsAzureBlobStoreSettings {
	var returns *ChronicleFeedDetailsAzureBlobStoreSettings
	_jsii_.Get(
		j,
		"azureBlobStoreSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AzureBlobStoreV2Settings() ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference {
	var returns ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference
	_jsii_.Get(
		j,
		"azureBlobStoreV2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AzureBlobStoreV2SettingsInput() *ChronicleFeedDetailsAzureBlobStoreV2Settings {
	var returns *ChronicleFeedDetailsAzureBlobStoreV2Settings
	_jsii_.Get(
		j,
		"azureBlobStoreV2SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AzureEventHubSettings() ChronicleFeedDetailsAzureEventHubSettingsOutputReference {
	var returns ChronicleFeedDetailsAzureEventHubSettingsOutputReference
	_jsii_.Get(
		j,
		"azureEventHubSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AzureEventHubSettingsInput() *ChronicleFeedDetailsAzureEventHubSettings {
	var returns *ChronicleFeedDetailsAzureEventHubSettings
	_jsii_.Get(
		j,
		"azureEventHubSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AzureMdmIntuneSettings() ChronicleFeedDetailsAzureMdmIntuneSettingsOutputReference {
	var returns ChronicleFeedDetailsAzureMdmIntuneSettingsOutputReference
	_jsii_.Get(
		j,
		"azureMdmIntuneSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) AzureMdmIntuneSettingsInput() *ChronicleFeedDetailsAzureMdmIntuneSettings {
	var returns *ChronicleFeedDetailsAzureMdmIntuneSettings
	_jsii_.Get(
		j,
		"azureMdmIntuneSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) CloudPassageSettings() ChronicleFeedDetailsCloudPassageSettingsOutputReference {
	var returns ChronicleFeedDetailsCloudPassageSettingsOutputReference
	_jsii_.Get(
		j,
		"cloudPassageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) CloudPassageSettingsInput() *ChronicleFeedDetailsCloudPassageSettings {
	var returns *ChronicleFeedDetailsCloudPassageSettings
	_jsii_.Get(
		j,
		"cloudPassageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) CortexXdrSettings() ChronicleFeedDetailsCortexXdrSettingsOutputReference {
	var returns ChronicleFeedDetailsCortexXdrSettingsOutputReference
	_jsii_.Get(
		j,
		"cortexXdrSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) CortexXdrSettingsInput() *ChronicleFeedDetailsCortexXdrSettings {
	var returns *ChronicleFeedDetailsCortexXdrSettings
	_jsii_.Get(
		j,
		"cortexXdrSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) CrowdstrikeAlertsSettings() ChronicleFeedDetailsCrowdstrikeAlertsSettingsOutputReference {
	var returns ChronicleFeedDetailsCrowdstrikeAlertsSettingsOutputReference
	_jsii_.Get(
		j,
		"crowdstrikeAlertsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) CrowdstrikeAlertsSettingsInput() *ChronicleFeedDetailsCrowdstrikeAlertsSettings {
	var returns *ChronicleFeedDetailsCrowdstrikeAlertsSettings
	_jsii_.Get(
		j,
		"crowdstrikeAlertsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) CrowdstrikeDetectsSettings() ChronicleFeedDetailsCrowdstrikeDetectsSettingsOutputReference {
	var returns ChronicleFeedDetailsCrowdstrikeDetectsSettingsOutputReference
	_jsii_.Get(
		j,
		"crowdstrikeDetectsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) CrowdstrikeDetectsSettingsInput() *ChronicleFeedDetailsCrowdstrikeDetectsSettings {
	var returns *ChronicleFeedDetailsCrowdstrikeDetectsSettings
	_jsii_.Get(
		j,
		"crowdstrikeDetectsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) DummyLogTypeSettings() ChronicleFeedDetailsDummyLogTypeSettingsOutputReference {
	var returns ChronicleFeedDetailsDummyLogTypeSettingsOutputReference
	_jsii_.Get(
		j,
		"dummyLogTypeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) DummyLogTypeSettingsInput() *ChronicleFeedDetailsDummyLogTypeSettings {
	var returns *ChronicleFeedDetailsDummyLogTypeSettings
	_jsii_.Get(
		j,
		"dummyLogTypeSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) DuoAuthSettings() ChronicleFeedDetailsDuoAuthSettingsOutputReference {
	var returns ChronicleFeedDetailsDuoAuthSettingsOutputReference
	_jsii_.Get(
		j,
		"duoAuthSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) DuoAuthSettingsInput() *ChronicleFeedDetailsDuoAuthSettings {
	var returns *ChronicleFeedDetailsDuoAuthSettings
	_jsii_.Get(
		j,
		"duoAuthSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) DuoUserContextSettings() ChronicleFeedDetailsDuoUserContextSettingsOutputReference {
	var returns ChronicleFeedDetailsDuoUserContextSettingsOutputReference
	_jsii_.Get(
		j,
		"duoUserContextSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) DuoUserContextSettingsInput() *ChronicleFeedDetailsDuoUserContextSettings {
	var returns *ChronicleFeedDetailsDuoUserContextSettings
	_jsii_.Get(
		j,
		"duoUserContextSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) FeedSourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"feedSourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) FeedSourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"feedSourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) FoxItStixSettings() ChronicleFeedDetailsFoxItStixSettingsOutputReference {
	var returns ChronicleFeedDetailsFoxItStixSettingsOutputReference
	_jsii_.Get(
		j,
		"foxItStixSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) FoxItStixSettingsInput() *ChronicleFeedDetailsFoxItStixSettings {
	var returns *ChronicleFeedDetailsFoxItStixSettings
	_jsii_.Get(
		j,
		"foxItStixSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) GcsSettings() ChronicleFeedDetailsGcsSettingsOutputReference {
	var returns ChronicleFeedDetailsGcsSettingsOutputReference
	_jsii_.Get(
		j,
		"gcsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) GcsSettingsInput() *ChronicleFeedDetailsGcsSettings {
	var returns *ChronicleFeedDetailsGcsSettings
	_jsii_.Get(
		j,
		"gcsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) GcsV2Settings() ChronicleFeedDetailsGcsV2SettingsOutputReference {
	var returns ChronicleFeedDetailsGcsV2SettingsOutputReference
	_jsii_.Get(
		j,
		"gcsV2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) GcsV2SettingsInput() *ChronicleFeedDetailsGcsV2Settings {
	var returns *ChronicleFeedDetailsGcsV2Settings
	_jsii_.Get(
		j,
		"gcsV2SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) GoogleCloudIdentityDevicesSettings() ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsOutputReference {
	var returns ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsOutputReference
	_jsii_.Get(
		j,
		"googleCloudIdentityDevicesSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) GoogleCloudIdentityDevicesSettingsInput() *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettings {
	var returns *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettings
	_jsii_.Get(
		j,
		"googleCloudIdentityDevicesSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) GoogleCloudIdentityDeviceUsersSettings() ChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettingsOutputReference {
	var returns ChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettingsOutputReference
	_jsii_.Get(
		j,
		"googleCloudIdentityDeviceUsersSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) GoogleCloudIdentityDeviceUsersSettingsInput() *ChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettings {
	var returns *ChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettings
	_jsii_.Get(
		j,
		"googleCloudIdentityDeviceUsersSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) GoogleCloudStorageEventDrivenSettings() ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference {
	var returns ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference
	_jsii_.Get(
		j,
		"googleCloudStorageEventDrivenSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) GoogleCloudStorageEventDrivenSettingsInput() *ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings {
	var returns *ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings
	_jsii_.Get(
		j,
		"googleCloudStorageEventDrivenSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) HttpSettings() ChronicleFeedDetailsHttpSettingsOutputReference {
	var returns ChronicleFeedDetailsHttpSettingsOutputReference
	_jsii_.Get(
		j,
		"httpSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) HttpSettingsInput() *ChronicleFeedDetailsHttpSettings {
	var returns *ChronicleFeedDetailsHttpSettings
	_jsii_.Get(
		j,
		"httpSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) HttpsPushAmazonKinesisFirehoseSettings() ChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettingsOutputReference {
	var returns ChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettingsOutputReference
	_jsii_.Get(
		j,
		"httpsPushAmazonKinesisFirehoseSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) HttpsPushAmazonKinesisFirehoseSettingsInput() *ChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettings {
	var returns *ChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettings
	_jsii_.Get(
		j,
		"httpsPushAmazonKinesisFirehoseSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) HttpsPushGoogleCloudPubsubSettings() ChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettingsOutputReference {
	var returns ChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettingsOutputReference
	_jsii_.Get(
		j,
		"httpsPushGoogleCloudPubsubSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) HttpsPushGoogleCloudPubsubSettingsInput() *ChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettings {
	var returns *ChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettings
	_jsii_.Get(
		j,
		"httpsPushGoogleCloudPubsubSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) HttpsPushWebhookSettings() ChronicleFeedDetailsHttpsPushWebhookSettingsOutputReference {
	var returns ChronicleFeedDetailsHttpsPushWebhookSettingsOutputReference
	_jsii_.Get(
		j,
		"httpsPushWebhookSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) HttpsPushWebhookSettingsInput() *ChronicleFeedDetailsHttpsPushWebhookSettings {
	var returns *ChronicleFeedDetailsHttpsPushWebhookSettings
	_jsii_.Get(
		j,
		"httpsPushWebhookSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) ImpervaWafSettings() ChronicleFeedDetailsImpervaWafSettingsOutputReference {
	var returns ChronicleFeedDetailsImpervaWafSettingsOutputReference
	_jsii_.Get(
		j,
		"impervaWafSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) ImpervaWafSettingsInput() *ChronicleFeedDetailsImpervaWafSettings {
	var returns *ChronicleFeedDetailsImpervaWafSettings
	_jsii_.Get(
		j,
		"impervaWafSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) InternalValue() *ChronicleFeedDetails {
	var returns *ChronicleFeedDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) LogTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) MandiantIocSettings() ChronicleFeedDetailsMandiantIocSettingsOutputReference {
	var returns ChronicleFeedDetailsMandiantIocSettingsOutputReference
	_jsii_.Get(
		j,
		"mandiantIocSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) MandiantIocSettingsInput() *ChronicleFeedDetailsMandiantIocSettings {
	var returns *ChronicleFeedDetailsMandiantIocSettings
	_jsii_.Get(
		j,
		"mandiantIocSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) MicrosoftGraphAlertSettings() ChronicleFeedDetailsMicrosoftGraphAlertSettingsOutputReference {
	var returns ChronicleFeedDetailsMicrosoftGraphAlertSettingsOutputReference
	_jsii_.Get(
		j,
		"microsoftGraphAlertSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) MicrosoftGraphAlertSettingsInput() *ChronicleFeedDetailsMicrosoftGraphAlertSettings {
	var returns *ChronicleFeedDetailsMicrosoftGraphAlertSettings
	_jsii_.Get(
		j,
		"microsoftGraphAlertSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) MicrosoftSecurityCenterAlertSettings() ChronicleFeedDetailsMicrosoftSecurityCenterAlertSettingsOutputReference {
	var returns ChronicleFeedDetailsMicrosoftSecurityCenterAlertSettingsOutputReference
	_jsii_.Get(
		j,
		"microsoftSecurityCenterAlertSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) MicrosoftSecurityCenterAlertSettingsInput() *ChronicleFeedDetailsMicrosoftSecurityCenterAlertSettings {
	var returns *ChronicleFeedDetailsMicrosoftSecurityCenterAlertSettings
	_jsii_.Get(
		j,
		"microsoftSecurityCenterAlertSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) MimecastMailSettings() ChronicleFeedDetailsMimecastMailSettingsOutputReference {
	var returns ChronicleFeedDetailsMimecastMailSettingsOutputReference
	_jsii_.Get(
		j,
		"mimecastMailSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) MimecastMailSettingsInput() *ChronicleFeedDetailsMimecastMailSettings {
	var returns *ChronicleFeedDetailsMimecastMailSettings
	_jsii_.Get(
		j,
		"mimecastMailSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) MimecastMailV2Settings() ChronicleFeedDetailsMimecastMailV2SettingsOutputReference {
	var returns ChronicleFeedDetailsMimecastMailV2SettingsOutputReference
	_jsii_.Get(
		j,
		"mimecastMailV2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) MimecastMailV2SettingsInput() *ChronicleFeedDetailsMimecastMailV2Settings {
	var returns *ChronicleFeedDetailsMimecastMailV2Settings
	_jsii_.Get(
		j,
		"mimecastMailV2SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) NetskopeAlertSettings() ChronicleFeedDetailsNetskopeAlertSettingsOutputReference {
	var returns ChronicleFeedDetailsNetskopeAlertSettingsOutputReference
	_jsii_.Get(
		j,
		"netskopeAlertSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) NetskopeAlertSettingsInput() *ChronicleFeedDetailsNetskopeAlertSettings {
	var returns *ChronicleFeedDetailsNetskopeAlertSettings
	_jsii_.Get(
		j,
		"netskopeAlertSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) NetskopeAlertV2Settings() ChronicleFeedDetailsNetskopeAlertV2SettingsOutputReference {
	var returns ChronicleFeedDetailsNetskopeAlertV2SettingsOutputReference
	_jsii_.Get(
		j,
		"netskopeAlertV2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) NetskopeAlertV2SettingsInput() *ChronicleFeedDetailsNetskopeAlertV2Settings {
	var returns *ChronicleFeedDetailsNetskopeAlertV2Settings
	_jsii_.Get(
		j,
		"netskopeAlertV2SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) Office365Settings() ChronicleFeedDetailsOffice365SettingsOutputReference {
	var returns ChronicleFeedDetailsOffice365SettingsOutputReference
	_jsii_.Get(
		j,
		"office365Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) Office365SettingsInput() *ChronicleFeedDetailsOffice365Settings {
	var returns *ChronicleFeedDetailsOffice365Settings
	_jsii_.Get(
		j,
		"office365SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) OktaSettings() ChronicleFeedDetailsOktaSettingsOutputReference {
	var returns ChronicleFeedDetailsOktaSettingsOutputReference
	_jsii_.Get(
		j,
		"oktaSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) OktaSettingsInput() *ChronicleFeedDetailsOktaSettings {
	var returns *ChronicleFeedDetailsOktaSettings
	_jsii_.Get(
		j,
		"oktaSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) OktaUserContextSettings() ChronicleFeedDetailsOktaUserContextSettingsOutputReference {
	var returns ChronicleFeedDetailsOktaUserContextSettingsOutputReference
	_jsii_.Get(
		j,
		"oktaUserContextSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) OktaUserContextSettingsInput() *ChronicleFeedDetailsOktaUserContextSettings {
	var returns *ChronicleFeedDetailsOktaUserContextSettings
	_jsii_.Get(
		j,
		"oktaUserContextSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) PanIocSettings() ChronicleFeedDetailsPanIocSettingsOutputReference {
	var returns ChronicleFeedDetailsPanIocSettingsOutputReference
	_jsii_.Get(
		j,
		"panIocSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) PanIocSettingsInput() *ChronicleFeedDetailsPanIocSettings {
	var returns *ChronicleFeedDetailsPanIocSettings
	_jsii_.Get(
		j,
		"panIocSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) PanPrismaCloudSettings() ChronicleFeedDetailsPanPrismaCloudSettingsOutputReference {
	var returns ChronicleFeedDetailsPanPrismaCloudSettingsOutputReference
	_jsii_.Get(
		j,
		"panPrismaCloudSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) PanPrismaCloudSettingsInput() *ChronicleFeedDetailsPanPrismaCloudSettings {
	var returns *ChronicleFeedDetailsPanPrismaCloudSettings
	_jsii_.Get(
		j,
		"panPrismaCloudSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) ProofpointMailSettings() ChronicleFeedDetailsProofpointMailSettingsOutputReference {
	var returns ChronicleFeedDetailsProofpointMailSettingsOutputReference
	_jsii_.Get(
		j,
		"proofpointMailSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) ProofpointMailSettingsInput() *ChronicleFeedDetailsProofpointMailSettings {
	var returns *ChronicleFeedDetailsProofpointMailSettings
	_jsii_.Get(
		j,
		"proofpointMailSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) ProofpointOnDemandSettings() ChronicleFeedDetailsProofpointOnDemandSettingsOutputReference {
	var returns ChronicleFeedDetailsProofpointOnDemandSettingsOutputReference
	_jsii_.Get(
		j,
		"proofpointOnDemandSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) ProofpointOnDemandSettingsInput() *ChronicleFeedDetailsProofpointOnDemandSettings {
	var returns *ChronicleFeedDetailsProofpointOnDemandSettings
	_jsii_.Get(
		j,
		"proofpointOnDemandSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) PubsubSettings() ChronicleFeedDetailsPubsubSettingsOutputReference {
	var returns ChronicleFeedDetailsPubsubSettingsOutputReference
	_jsii_.Get(
		j,
		"pubsubSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) PubsubSettingsInput() *ChronicleFeedDetailsPubsubSettings {
	var returns *ChronicleFeedDetailsPubsubSettings
	_jsii_.Get(
		j,
		"pubsubSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) QualysScanSettings() ChronicleFeedDetailsQualysScanSettingsOutputReference {
	var returns ChronicleFeedDetailsQualysScanSettingsOutputReference
	_jsii_.Get(
		j,
		"qualysScanSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) QualysScanSettingsInput() *ChronicleFeedDetailsQualysScanSettings {
	var returns *ChronicleFeedDetailsQualysScanSettings
	_jsii_.Get(
		j,
		"qualysScanSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) QualysVmSettings() ChronicleFeedDetailsQualysVmSettingsOutputReference {
	var returns ChronicleFeedDetailsQualysVmSettingsOutputReference
	_jsii_.Get(
		j,
		"qualysVmSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) QualysVmSettingsInput() *ChronicleFeedDetailsQualysVmSettings {
	var returns *ChronicleFeedDetailsQualysVmSettings
	_jsii_.Get(
		j,
		"qualysVmSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) Rapid7InsightSettings() ChronicleFeedDetailsRapid7InsightSettingsOutputReference {
	var returns ChronicleFeedDetailsRapid7InsightSettingsOutputReference
	_jsii_.Get(
		j,
		"rapid7InsightSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) Rapid7InsightSettingsInput() *ChronicleFeedDetailsRapid7InsightSettings {
	var returns *ChronicleFeedDetailsRapid7InsightSettings
	_jsii_.Get(
		j,
		"rapid7InsightSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) RecordedFutureIocSettings() ChronicleFeedDetailsRecordedFutureIocSettingsOutputReference {
	var returns ChronicleFeedDetailsRecordedFutureIocSettingsOutputReference
	_jsii_.Get(
		j,
		"recordedFutureIocSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) RecordedFutureIocSettingsInput() *ChronicleFeedDetailsRecordedFutureIocSettings {
	var returns *ChronicleFeedDetailsRecordedFutureIocSettings
	_jsii_.Get(
		j,
		"recordedFutureIocSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) RhIsacIocSettings() ChronicleFeedDetailsRhIsacIocSettingsOutputReference {
	var returns ChronicleFeedDetailsRhIsacIocSettingsOutputReference
	_jsii_.Get(
		j,
		"rhIsacIocSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) RhIsacIocSettingsInput() *ChronicleFeedDetailsRhIsacIocSettings {
	var returns *ChronicleFeedDetailsRhIsacIocSettings
	_jsii_.Get(
		j,
		"rhIsacIocSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) SalesforceSettings() ChronicleFeedDetailsSalesforceSettingsOutputReference {
	var returns ChronicleFeedDetailsSalesforceSettingsOutputReference
	_jsii_.Get(
		j,
		"salesforceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) SalesforceSettingsInput() *ChronicleFeedDetailsSalesforceSettings {
	var returns *ChronicleFeedDetailsSalesforceSettings
	_jsii_.Get(
		j,
		"salesforceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) SentineloneAlertSettings() ChronicleFeedDetailsSentineloneAlertSettingsOutputReference {
	var returns ChronicleFeedDetailsSentineloneAlertSettingsOutputReference
	_jsii_.Get(
		j,
		"sentineloneAlertSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) SentineloneAlertSettingsInput() *ChronicleFeedDetailsSentineloneAlertSettings {
	var returns *ChronicleFeedDetailsSentineloneAlertSettings
	_jsii_.Get(
		j,
		"sentineloneAlertSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) ServiceNowCmdbSettings() ChronicleFeedDetailsServiceNowCmdbSettingsOutputReference {
	var returns ChronicleFeedDetailsServiceNowCmdbSettingsOutputReference
	_jsii_.Get(
		j,
		"serviceNowCmdbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) ServiceNowCmdbSettingsInput() *ChronicleFeedDetailsServiceNowCmdbSettings {
	var returns *ChronicleFeedDetailsServiceNowCmdbSettings
	_jsii_.Get(
		j,
		"serviceNowCmdbSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) SftpSettings() ChronicleFeedDetailsSftpSettingsOutputReference {
	var returns ChronicleFeedDetailsSftpSettingsOutputReference
	_jsii_.Get(
		j,
		"sftpSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) SftpSettingsInput() *ChronicleFeedDetailsSftpSettings {
	var returns *ChronicleFeedDetailsSftpSettings
	_jsii_.Get(
		j,
		"sftpSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) StsMigrationReadiness() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsMigrationReadiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) SymantecEventExportSettings() ChronicleFeedDetailsSymantecEventExportSettingsOutputReference {
	var returns ChronicleFeedDetailsSymantecEventExportSettingsOutputReference
	_jsii_.Get(
		j,
		"symantecEventExportSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) SymantecEventExportSettingsInput() *ChronicleFeedDetailsSymantecEventExportSettings {
	var returns *ChronicleFeedDetailsSymantecEventExportSettings
	_jsii_.Get(
		j,
		"symantecEventExportSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) ThinkstCanarySettings() ChronicleFeedDetailsThinkstCanarySettingsOutputReference {
	var returns ChronicleFeedDetailsThinkstCanarySettingsOutputReference
	_jsii_.Get(
		j,
		"thinkstCanarySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) ThinkstCanarySettingsInput() *ChronicleFeedDetailsThinkstCanarySettings {
	var returns *ChronicleFeedDetailsThinkstCanarySettings
	_jsii_.Get(
		j,
		"thinkstCanarySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) ThreatConnectIocSettings() ChronicleFeedDetailsThreatConnectIocSettingsOutputReference {
	var returns ChronicleFeedDetailsThreatConnectIocSettingsOutputReference
	_jsii_.Get(
		j,
		"threatConnectIocSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) ThreatConnectIocSettingsInput() *ChronicleFeedDetailsThreatConnectIocSettings {
	var returns *ChronicleFeedDetailsThreatConnectIocSettings
	_jsii_.Get(
		j,
		"threatConnectIocSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) ThreatConnectIocV3Settings() ChronicleFeedDetailsThreatConnectIocV3SettingsOutputReference {
	var returns ChronicleFeedDetailsThreatConnectIocV3SettingsOutputReference
	_jsii_.Get(
		j,
		"threatConnectIocV3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) ThreatConnectIocV3SettingsInput() *ChronicleFeedDetailsThreatConnectIocV3Settings {
	var returns *ChronicleFeedDetailsThreatConnectIocV3Settings
	_jsii_.Get(
		j,
		"threatConnectIocV3SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) TrellixHxAlertsSettings() ChronicleFeedDetailsTrellixHxAlertsSettingsOutputReference {
	var returns ChronicleFeedDetailsTrellixHxAlertsSettingsOutputReference
	_jsii_.Get(
		j,
		"trellixHxAlertsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) TrellixHxAlertsSettingsInput() *ChronicleFeedDetailsTrellixHxAlertsSettings {
	var returns *ChronicleFeedDetailsTrellixHxAlertsSettings
	_jsii_.Get(
		j,
		"trellixHxAlertsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) TrellixHxBulkAcqsSettings() ChronicleFeedDetailsTrellixHxBulkAcqsSettingsOutputReference {
	var returns ChronicleFeedDetailsTrellixHxBulkAcqsSettingsOutputReference
	_jsii_.Get(
		j,
		"trellixHxBulkAcqsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) TrellixHxBulkAcqsSettingsInput() *ChronicleFeedDetailsTrellixHxBulkAcqsSettings {
	var returns *ChronicleFeedDetailsTrellixHxBulkAcqsSettings
	_jsii_.Get(
		j,
		"trellixHxBulkAcqsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) TrellixHxHostsSettings() ChronicleFeedDetailsTrellixHxHostsSettingsOutputReference {
	var returns ChronicleFeedDetailsTrellixHxHostsSettingsOutputReference
	_jsii_.Get(
		j,
		"trellixHxHostsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) TrellixHxHostsSettingsInput() *ChronicleFeedDetailsTrellixHxHostsSettings {
	var returns *ChronicleFeedDetailsTrellixHxHostsSettings
	_jsii_.Get(
		j,
		"trellixHxHostsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WebhookSettings() ChronicleFeedDetailsWebhookSettingsOutputReference {
	var returns ChronicleFeedDetailsWebhookSettingsOutputReference
	_jsii_.Get(
		j,
		"webhookSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WebhookSettingsInput() *ChronicleFeedDetailsWebhookSettings {
	var returns *ChronicleFeedDetailsWebhookSettings
	_jsii_.Get(
		j,
		"webhookSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WorkdaySettings() ChronicleFeedDetailsWorkdaySettingsOutputReference {
	var returns ChronicleFeedDetailsWorkdaySettingsOutputReference
	_jsii_.Get(
		j,
		"workdaySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WorkdaySettingsInput() *ChronicleFeedDetailsWorkdaySettings {
	var returns *ChronicleFeedDetailsWorkdaySettings
	_jsii_.Get(
		j,
		"workdaySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WorkspaceActivitySettings() ChronicleFeedDetailsWorkspaceActivitySettingsOutputReference {
	var returns ChronicleFeedDetailsWorkspaceActivitySettingsOutputReference
	_jsii_.Get(
		j,
		"workspaceActivitySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WorkspaceActivitySettingsInput() *ChronicleFeedDetailsWorkspaceActivitySettings {
	var returns *ChronicleFeedDetailsWorkspaceActivitySettings
	_jsii_.Get(
		j,
		"workspaceActivitySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WorkspaceAlertsSettings() ChronicleFeedDetailsWorkspaceAlertsSettingsOutputReference {
	var returns ChronicleFeedDetailsWorkspaceAlertsSettingsOutputReference
	_jsii_.Get(
		j,
		"workspaceAlertsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WorkspaceAlertsSettingsInput() *ChronicleFeedDetailsWorkspaceAlertsSettings {
	var returns *ChronicleFeedDetailsWorkspaceAlertsSettings
	_jsii_.Get(
		j,
		"workspaceAlertsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WorkspaceChromeOsSettings() ChronicleFeedDetailsWorkspaceChromeOsSettingsOutputReference {
	var returns ChronicleFeedDetailsWorkspaceChromeOsSettingsOutputReference
	_jsii_.Get(
		j,
		"workspaceChromeOsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WorkspaceChromeOsSettingsInput() *ChronicleFeedDetailsWorkspaceChromeOsSettings {
	var returns *ChronicleFeedDetailsWorkspaceChromeOsSettings
	_jsii_.Get(
		j,
		"workspaceChromeOsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WorkspaceGroupsSettings() ChronicleFeedDetailsWorkspaceGroupsSettingsOutputReference {
	var returns ChronicleFeedDetailsWorkspaceGroupsSettingsOutputReference
	_jsii_.Get(
		j,
		"workspaceGroupsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WorkspaceGroupsSettingsInput() *ChronicleFeedDetailsWorkspaceGroupsSettings {
	var returns *ChronicleFeedDetailsWorkspaceGroupsSettings
	_jsii_.Get(
		j,
		"workspaceGroupsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WorkspaceMobileSettings() ChronicleFeedDetailsWorkspaceMobileSettingsOutputReference {
	var returns ChronicleFeedDetailsWorkspaceMobileSettingsOutputReference
	_jsii_.Get(
		j,
		"workspaceMobileSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WorkspaceMobileSettingsInput() *ChronicleFeedDetailsWorkspaceMobileSettings {
	var returns *ChronicleFeedDetailsWorkspaceMobileSettings
	_jsii_.Get(
		j,
		"workspaceMobileSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WorkspacePrivilegesSettings() ChronicleFeedDetailsWorkspacePrivilegesSettingsOutputReference {
	var returns ChronicleFeedDetailsWorkspacePrivilegesSettingsOutputReference
	_jsii_.Get(
		j,
		"workspacePrivilegesSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WorkspacePrivilegesSettingsInput() *ChronicleFeedDetailsWorkspacePrivilegesSettings {
	var returns *ChronicleFeedDetailsWorkspacePrivilegesSettings
	_jsii_.Get(
		j,
		"workspacePrivilegesSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WorkspaceUsersSettings() ChronicleFeedDetailsWorkspaceUsersSettingsOutputReference {
	var returns ChronicleFeedDetailsWorkspaceUsersSettingsOutputReference
	_jsii_.Get(
		j,
		"workspaceUsersSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference) WorkspaceUsersSettingsInput() *ChronicleFeedDetailsWorkspaceUsersSettings {
	var returns *ChronicleFeedDetailsWorkspaceUsersSettings
	_jsii_.Get(
		j,
		"workspaceUsersSettingsInput",
		&returns,
	)
	return returns
}


func NewChronicleFeedDetailsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleFeedDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleFeedDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleFeedDetailsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleFeedDetailsOutputReference_Override(c ChronicleFeedDetailsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference)SetAssetNamespace(val *string) {
	if err := j.validateSetAssetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetNamespace",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference)SetFeedSourceType(val *string) {
	if err := j.validateSetFeedSourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"feedSourceType",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference)SetInternalValue(val *ChronicleFeedDetails) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference)SetLogType(val *string) {
	if err := j.validateSetLogTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logType",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutAmazonKinesisFirehoseSettings(value *ChronicleFeedDetailsAmazonKinesisFirehoseSettings) {
	if err := c.validatePutAmazonKinesisFirehoseSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAmazonKinesisFirehoseSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutAmazonS3Settings(value *ChronicleFeedDetailsAmazonS3Settings) {
	if err := c.validatePutAmazonS3SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAmazonS3Settings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutAmazonS3V2Settings(value *ChronicleFeedDetailsAmazonS3V2Settings) {
	if err := c.validatePutAmazonS3V2SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAmazonS3V2Settings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutAmazonSqsSettings(value *ChronicleFeedDetailsAmazonSqsSettings) {
	if err := c.validatePutAmazonSqsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAmazonSqsSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutAmazonSqsV2Settings(value *ChronicleFeedDetailsAmazonSqsV2Settings) {
	if err := c.validatePutAmazonSqsV2SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAmazonSqsV2Settings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutAnomaliSettings(value *ChronicleFeedDetailsAnomaliSettings) {
	if err := c.validatePutAnomaliSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAnomaliSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutAwsEc2HostsSettings(value *ChronicleFeedDetailsAwsEc2HostsSettings) {
	if err := c.validatePutAwsEc2HostsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsEc2HostsSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutAwsEc2InstancesSettings(value *ChronicleFeedDetailsAwsEc2InstancesSettings) {
	if err := c.validatePutAwsEc2InstancesSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsEc2InstancesSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutAwsEc2VpcsSettings(value *ChronicleFeedDetailsAwsEc2VpcsSettings) {
	if err := c.validatePutAwsEc2VpcsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsEc2VpcsSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutAwsIamSettings(value *ChronicleFeedDetailsAwsIamSettings) {
	if err := c.validatePutAwsIamSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsIamSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutAzureAdAuditSettings(value *ChronicleFeedDetailsAzureAdAuditSettings) {
	if err := c.validatePutAzureAdAuditSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAzureAdAuditSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutAzureAdContextSettings(value *ChronicleFeedDetailsAzureAdContextSettings) {
	if err := c.validatePutAzureAdContextSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAzureAdContextSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutAzureAdSettings(value *ChronicleFeedDetailsAzureAdSettings) {
	if err := c.validatePutAzureAdSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAzureAdSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutAzureBlobStoreSettings(value *ChronicleFeedDetailsAzureBlobStoreSettings) {
	if err := c.validatePutAzureBlobStoreSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAzureBlobStoreSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutAzureBlobStoreV2Settings(value *ChronicleFeedDetailsAzureBlobStoreV2Settings) {
	if err := c.validatePutAzureBlobStoreV2SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAzureBlobStoreV2Settings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutAzureEventHubSettings(value *ChronicleFeedDetailsAzureEventHubSettings) {
	if err := c.validatePutAzureEventHubSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAzureEventHubSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutAzureMdmIntuneSettings(value *ChronicleFeedDetailsAzureMdmIntuneSettings) {
	if err := c.validatePutAzureMdmIntuneSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAzureMdmIntuneSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutCloudPassageSettings(value *ChronicleFeedDetailsCloudPassageSettings) {
	if err := c.validatePutCloudPassageSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCloudPassageSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutCortexXdrSettings(value *ChronicleFeedDetailsCortexXdrSettings) {
	if err := c.validatePutCortexXdrSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCortexXdrSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutCrowdstrikeAlertsSettings(value *ChronicleFeedDetailsCrowdstrikeAlertsSettings) {
	if err := c.validatePutCrowdstrikeAlertsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCrowdstrikeAlertsSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutCrowdstrikeDetectsSettings(value *ChronicleFeedDetailsCrowdstrikeDetectsSettings) {
	if err := c.validatePutCrowdstrikeDetectsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCrowdstrikeDetectsSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutDummyLogTypeSettings(value *ChronicleFeedDetailsDummyLogTypeSettings) {
	if err := c.validatePutDummyLogTypeSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDummyLogTypeSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutDuoAuthSettings(value *ChronicleFeedDetailsDuoAuthSettings) {
	if err := c.validatePutDuoAuthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDuoAuthSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutDuoUserContextSettings(value *ChronicleFeedDetailsDuoUserContextSettings) {
	if err := c.validatePutDuoUserContextSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDuoUserContextSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutFoxItStixSettings(value *ChronicleFeedDetailsFoxItStixSettings) {
	if err := c.validatePutFoxItStixSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFoxItStixSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutGcsSettings(value *ChronicleFeedDetailsGcsSettings) {
	if err := c.validatePutGcsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGcsSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutGcsV2Settings(value *ChronicleFeedDetailsGcsV2Settings) {
	if err := c.validatePutGcsV2SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGcsV2Settings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutGoogleCloudIdentityDevicesSettings(value *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettings) {
	if err := c.validatePutGoogleCloudIdentityDevicesSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGoogleCloudIdentityDevicesSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutGoogleCloudIdentityDeviceUsersSettings(value *ChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettings) {
	if err := c.validatePutGoogleCloudIdentityDeviceUsersSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGoogleCloudIdentityDeviceUsersSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutGoogleCloudStorageEventDrivenSettings(value *ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings) {
	if err := c.validatePutGoogleCloudStorageEventDrivenSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGoogleCloudStorageEventDrivenSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutHttpSettings(value *ChronicleFeedDetailsHttpSettings) {
	if err := c.validatePutHttpSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutHttpsPushAmazonKinesisFirehoseSettings(value *ChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettings) {
	if err := c.validatePutHttpsPushAmazonKinesisFirehoseSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpsPushAmazonKinesisFirehoseSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutHttpsPushGoogleCloudPubsubSettings(value *ChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettings) {
	if err := c.validatePutHttpsPushGoogleCloudPubsubSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpsPushGoogleCloudPubsubSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutHttpsPushWebhookSettings(value *ChronicleFeedDetailsHttpsPushWebhookSettings) {
	if err := c.validatePutHttpsPushWebhookSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpsPushWebhookSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutImpervaWafSettings(value *ChronicleFeedDetailsImpervaWafSettings) {
	if err := c.validatePutImpervaWafSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putImpervaWafSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutMandiantIocSettings(value *ChronicleFeedDetailsMandiantIocSettings) {
	if err := c.validatePutMandiantIocSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMandiantIocSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutMicrosoftGraphAlertSettings(value *ChronicleFeedDetailsMicrosoftGraphAlertSettings) {
	if err := c.validatePutMicrosoftGraphAlertSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMicrosoftGraphAlertSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutMicrosoftSecurityCenterAlertSettings(value *ChronicleFeedDetailsMicrosoftSecurityCenterAlertSettings) {
	if err := c.validatePutMicrosoftSecurityCenterAlertSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMicrosoftSecurityCenterAlertSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutMimecastMailSettings(value *ChronicleFeedDetailsMimecastMailSettings) {
	if err := c.validatePutMimecastMailSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMimecastMailSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutMimecastMailV2Settings(value *ChronicleFeedDetailsMimecastMailV2Settings) {
	if err := c.validatePutMimecastMailV2SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMimecastMailV2Settings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutNetskopeAlertSettings(value *ChronicleFeedDetailsNetskopeAlertSettings) {
	if err := c.validatePutNetskopeAlertSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNetskopeAlertSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutNetskopeAlertV2Settings(value *ChronicleFeedDetailsNetskopeAlertV2Settings) {
	if err := c.validatePutNetskopeAlertV2SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNetskopeAlertV2Settings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutOffice365Settings(value *ChronicleFeedDetailsOffice365Settings) {
	if err := c.validatePutOffice365SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOffice365Settings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutOktaSettings(value *ChronicleFeedDetailsOktaSettings) {
	if err := c.validatePutOktaSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOktaSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutOktaUserContextSettings(value *ChronicleFeedDetailsOktaUserContextSettings) {
	if err := c.validatePutOktaUserContextSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOktaUserContextSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutPanIocSettings(value *ChronicleFeedDetailsPanIocSettings) {
	if err := c.validatePutPanIocSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPanIocSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutPanPrismaCloudSettings(value *ChronicleFeedDetailsPanPrismaCloudSettings) {
	if err := c.validatePutPanPrismaCloudSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPanPrismaCloudSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutProofpointMailSettings(value *ChronicleFeedDetailsProofpointMailSettings) {
	if err := c.validatePutProofpointMailSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putProofpointMailSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutProofpointOnDemandSettings(value *ChronicleFeedDetailsProofpointOnDemandSettings) {
	if err := c.validatePutProofpointOnDemandSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putProofpointOnDemandSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutPubsubSettings(value *ChronicleFeedDetailsPubsubSettings) {
	if err := c.validatePutPubsubSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPubsubSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutQualysScanSettings(value *ChronicleFeedDetailsQualysScanSettings) {
	if err := c.validatePutQualysScanSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQualysScanSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutQualysVmSettings(value *ChronicleFeedDetailsQualysVmSettings) {
	if err := c.validatePutQualysVmSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQualysVmSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutRapid7InsightSettings(value *ChronicleFeedDetailsRapid7InsightSettings) {
	if err := c.validatePutRapid7InsightSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRapid7InsightSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutRecordedFutureIocSettings(value *ChronicleFeedDetailsRecordedFutureIocSettings) {
	if err := c.validatePutRecordedFutureIocSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRecordedFutureIocSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutRhIsacIocSettings(value *ChronicleFeedDetailsRhIsacIocSettings) {
	if err := c.validatePutRhIsacIocSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRhIsacIocSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutSalesforceSettings(value *ChronicleFeedDetailsSalesforceSettings) {
	if err := c.validatePutSalesforceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSalesforceSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutSentineloneAlertSettings(value *ChronicleFeedDetailsSentineloneAlertSettings) {
	if err := c.validatePutSentineloneAlertSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSentineloneAlertSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutServiceNowCmdbSettings(value *ChronicleFeedDetailsServiceNowCmdbSettings) {
	if err := c.validatePutServiceNowCmdbSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServiceNowCmdbSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutSftpSettings(value *ChronicleFeedDetailsSftpSettings) {
	if err := c.validatePutSftpSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSftpSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutSymantecEventExportSettings(value *ChronicleFeedDetailsSymantecEventExportSettings) {
	if err := c.validatePutSymantecEventExportSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSymantecEventExportSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutThinkstCanarySettings(value *ChronicleFeedDetailsThinkstCanarySettings) {
	if err := c.validatePutThinkstCanarySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putThinkstCanarySettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutThreatConnectIocSettings(value *ChronicleFeedDetailsThreatConnectIocSettings) {
	if err := c.validatePutThreatConnectIocSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putThreatConnectIocSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutThreatConnectIocV3Settings(value *ChronicleFeedDetailsThreatConnectIocV3Settings) {
	if err := c.validatePutThreatConnectIocV3SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putThreatConnectIocV3Settings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutTrellixHxAlertsSettings(value *ChronicleFeedDetailsTrellixHxAlertsSettings) {
	if err := c.validatePutTrellixHxAlertsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTrellixHxAlertsSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutTrellixHxBulkAcqsSettings(value *ChronicleFeedDetailsTrellixHxBulkAcqsSettings) {
	if err := c.validatePutTrellixHxBulkAcqsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTrellixHxBulkAcqsSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutTrellixHxHostsSettings(value *ChronicleFeedDetailsTrellixHxHostsSettings) {
	if err := c.validatePutTrellixHxHostsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTrellixHxHostsSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutWebhookSettings(value *ChronicleFeedDetailsWebhookSettings) {
	if err := c.validatePutWebhookSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWebhookSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutWorkdaySettings(value *ChronicleFeedDetailsWorkdaySettings) {
	if err := c.validatePutWorkdaySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWorkdaySettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutWorkspaceActivitySettings(value *ChronicleFeedDetailsWorkspaceActivitySettings) {
	if err := c.validatePutWorkspaceActivitySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWorkspaceActivitySettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutWorkspaceAlertsSettings(value *ChronicleFeedDetailsWorkspaceAlertsSettings) {
	if err := c.validatePutWorkspaceAlertsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWorkspaceAlertsSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutWorkspaceChromeOsSettings(value *ChronicleFeedDetailsWorkspaceChromeOsSettings) {
	if err := c.validatePutWorkspaceChromeOsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWorkspaceChromeOsSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutWorkspaceGroupsSettings(value *ChronicleFeedDetailsWorkspaceGroupsSettings) {
	if err := c.validatePutWorkspaceGroupsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWorkspaceGroupsSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutWorkspaceMobileSettings(value *ChronicleFeedDetailsWorkspaceMobileSettings) {
	if err := c.validatePutWorkspaceMobileSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWorkspaceMobileSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutWorkspacePrivilegesSettings(value *ChronicleFeedDetailsWorkspacePrivilegesSettings) {
	if err := c.validatePutWorkspacePrivilegesSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWorkspacePrivilegesSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) PutWorkspaceUsersSettings(value *ChronicleFeedDetailsWorkspaceUsersSettings) {
	if err := c.validatePutWorkspaceUsersSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWorkspaceUsersSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAmazonKinesisFirehoseSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetAmazonKinesisFirehoseSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAmazonS3Settings() {
	_jsii_.InvokeVoid(
		c,
		"resetAmazonS3Settings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAmazonS3V2Settings() {
	_jsii_.InvokeVoid(
		c,
		"resetAmazonS3V2Settings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAmazonSqsSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetAmazonSqsSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAmazonSqsV2Settings() {
	_jsii_.InvokeVoid(
		c,
		"resetAmazonSqsV2Settings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAnomaliSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetAnomaliSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAssetNamespace() {
	_jsii_.InvokeVoid(
		c,
		"resetAssetNamespace",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAwsEc2HostsSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsEc2HostsSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAwsEc2InstancesSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsEc2InstancesSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAwsEc2VpcsSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsEc2VpcsSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAwsIamSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsIamSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAzureAdAuditSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureAdAuditSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAzureAdContextSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureAdContextSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAzureAdSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureAdSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAzureBlobStoreSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureBlobStoreSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAzureBlobStoreV2Settings() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureBlobStoreV2Settings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAzureEventHubSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureEventHubSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetAzureMdmIntuneSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureMdmIntuneSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetCloudPassageSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudPassageSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetCortexXdrSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetCortexXdrSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetCrowdstrikeAlertsSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetCrowdstrikeAlertsSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetCrowdstrikeDetectsSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetCrowdstrikeDetectsSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetDummyLogTypeSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetDummyLogTypeSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetDuoAuthSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetDuoAuthSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetDuoUserContextSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetDuoUserContextSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetFeedSourceType() {
	_jsii_.InvokeVoid(
		c,
		"resetFeedSourceType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetFoxItStixSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetFoxItStixSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetGcsSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetGcsSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetGcsV2Settings() {
	_jsii_.InvokeVoid(
		c,
		"resetGcsV2Settings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetGoogleCloudIdentityDevicesSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetGoogleCloudIdentityDevicesSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetGoogleCloudIdentityDeviceUsersSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetGoogleCloudIdentityDeviceUsersSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetGoogleCloudStorageEventDrivenSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetGoogleCloudStorageEventDrivenSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetHttpSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetHttpsPushAmazonKinesisFirehoseSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpsPushAmazonKinesisFirehoseSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetHttpsPushGoogleCloudPubsubSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpsPushGoogleCloudPubsubSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetHttpsPushWebhookSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpsPushWebhookSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetImpervaWafSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetImpervaWafSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetMandiantIocSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetMandiantIocSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetMicrosoftGraphAlertSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetMicrosoftGraphAlertSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetMicrosoftSecurityCenterAlertSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetMicrosoftSecurityCenterAlertSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetMimecastMailSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetMimecastMailSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetMimecastMailV2Settings() {
	_jsii_.InvokeVoid(
		c,
		"resetMimecastMailV2Settings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetNetskopeAlertSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetNetskopeAlertSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetNetskopeAlertV2Settings() {
	_jsii_.InvokeVoid(
		c,
		"resetNetskopeAlertV2Settings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetOffice365Settings() {
	_jsii_.InvokeVoid(
		c,
		"resetOffice365Settings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetOktaSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetOktaSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetOktaUserContextSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetOktaUserContextSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetPanIocSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetPanIocSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetPanPrismaCloudSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetPanPrismaCloudSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetProofpointMailSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetProofpointMailSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetProofpointOnDemandSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetProofpointOnDemandSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetPubsubSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetPubsubSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetQualysScanSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetQualysScanSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetQualysVmSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetQualysVmSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetRapid7InsightSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetRapid7InsightSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetRecordedFutureIocSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetRecordedFutureIocSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetRhIsacIocSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetRhIsacIocSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetSalesforceSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetSalesforceSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetSentineloneAlertSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetSentineloneAlertSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetServiceNowCmdbSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceNowCmdbSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetSftpSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetSftpSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetSymantecEventExportSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetSymantecEventExportSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetThinkstCanarySettings() {
	_jsii_.InvokeVoid(
		c,
		"resetThinkstCanarySettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetThreatConnectIocSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetThreatConnectIocSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetThreatConnectIocV3Settings() {
	_jsii_.InvokeVoid(
		c,
		"resetThreatConnectIocV3Settings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetTrellixHxAlertsSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetTrellixHxAlertsSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetTrellixHxBulkAcqsSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetTrellixHxBulkAcqsSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetTrellixHxHostsSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetTrellixHxHostsSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetWebhookSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetWebhookSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetWorkdaySettings() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkdaySettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetWorkspaceActivitySettings() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkspaceActivitySettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetWorkspaceAlertsSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkspaceAlertsSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetWorkspaceChromeOsSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkspaceChromeOsSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetWorkspaceGroupsSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkspaceGroupsSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetWorkspaceMobileSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkspaceMobileSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetWorkspacePrivilegesSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkspacePrivilegesSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ResetWorkspaceUsersSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkspaceUsersSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

