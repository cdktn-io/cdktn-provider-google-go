// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetails struct {
	// LogType. Format: projects/{project}/locations/{location}/instances/{instance}/logTypes/{log_type}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#log_type ChronicleFeed#log_type}
	LogType *string `field:"required" json:"logType" yaml:"logType"`
	// amazon_kinesis_firehose_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#amazon_kinesis_firehose_settings ChronicleFeed#amazon_kinesis_firehose_settings}
	AmazonKinesisFirehoseSettings *ChronicleFeedDetailsAmazonKinesisFirehoseSettings `field:"optional" json:"amazonKinesisFirehoseSettings" yaml:"amazonKinesisFirehoseSettings"`
	// amazon_s3_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#amazon_s3_settings ChronicleFeed#amazon_s3_settings}
	AmazonS3Settings *ChronicleFeedDetailsAmazonS3Settings `field:"optional" json:"amazonS3Settings" yaml:"amazonS3Settings"`
	// amazon_s3_v2_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#amazon_s3_v2_settings ChronicleFeed#amazon_s3_v2_settings}
	AmazonS3V2Settings *ChronicleFeedDetailsAmazonS3V2Settings `field:"optional" json:"amazonS3V2Settings" yaml:"amazonS3V2Settings"`
	// amazon_sqs_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#amazon_sqs_settings ChronicleFeed#amazon_sqs_settings}
	AmazonSqsSettings *ChronicleFeedDetailsAmazonSqsSettings `field:"optional" json:"amazonSqsSettings" yaml:"amazonSqsSettings"`
	// amazon_sqs_v2_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#amazon_sqs_v2_settings ChronicleFeed#amazon_sqs_v2_settings}
	AmazonSqsV2Settings *ChronicleFeedDetailsAmazonSqsV2Settings `field:"optional" json:"amazonSqsV2Settings" yaml:"amazonSqsV2Settings"`
	// anomali_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#anomali_settings ChronicleFeed#anomali_settings}
	AnomaliSettings *ChronicleFeedDetailsAnomaliSettings `field:"optional" json:"anomaliSettings" yaml:"anomaliSettings"`
	// The asset namespace to apply to all logs ingested through this feed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#asset_namespace ChronicleFeed#asset_namespace}
	AssetNamespace *string `field:"optional" json:"assetNamespace" yaml:"assetNamespace"`
	// aws_ec2_hosts_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#aws_ec2_hosts_settings ChronicleFeed#aws_ec2_hosts_settings}
	AwsEc2HostsSettings *ChronicleFeedDetailsAwsEc2HostsSettings `field:"optional" json:"awsEc2HostsSettings" yaml:"awsEc2HostsSettings"`
	// aws_ec2_instances_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#aws_ec2_instances_settings ChronicleFeed#aws_ec2_instances_settings}
	AwsEc2InstancesSettings *ChronicleFeedDetailsAwsEc2InstancesSettings `field:"optional" json:"awsEc2InstancesSettings" yaml:"awsEc2InstancesSettings"`
	// aws_ec2_vpcs_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#aws_ec2_vpcs_settings ChronicleFeed#aws_ec2_vpcs_settings}
	AwsEc2VpcsSettings *ChronicleFeedDetailsAwsEc2VpcsSettings `field:"optional" json:"awsEc2VpcsSettings" yaml:"awsEc2VpcsSettings"`
	// aws_iam_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#aws_iam_settings ChronicleFeed#aws_iam_settings}
	AwsIamSettings *ChronicleFeedDetailsAwsIamSettings `field:"optional" json:"awsIamSettings" yaml:"awsIamSettings"`
	// azure_ad_audit_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#azure_ad_audit_settings ChronicleFeed#azure_ad_audit_settings}
	AzureAdAuditSettings *ChronicleFeedDetailsAzureAdAuditSettings `field:"optional" json:"azureAdAuditSettings" yaml:"azureAdAuditSettings"`
	// azure_ad_context_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#azure_ad_context_settings ChronicleFeed#azure_ad_context_settings}
	AzureAdContextSettings *ChronicleFeedDetailsAzureAdContextSettings `field:"optional" json:"azureAdContextSettings" yaml:"azureAdContextSettings"`
	// azure_ad_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#azure_ad_settings ChronicleFeed#azure_ad_settings}
	AzureAdSettings *ChronicleFeedDetailsAzureAdSettings `field:"optional" json:"azureAdSettings" yaml:"azureAdSettings"`
	// azure_blob_store_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#azure_blob_store_settings ChronicleFeed#azure_blob_store_settings}
	AzureBlobStoreSettings *ChronicleFeedDetailsAzureBlobStoreSettings `field:"optional" json:"azureBlobStoreSettings" yaml:"azureBlobStoreSettings"`
	// azure_blob_store_v2_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#azure_blob_store_v2_settings ChronicleFeed#azure_blob_store_v2_settings}
	AzureBlobStoreV2Settings *ChronicleFeedDetailsAzureBlobStoreV2Settings `field:"optional" json:"azureBlobStoreV2Settings" yaml:"azureBlobStoreV2Settings"`
	// azure_event_hub_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#azure_event_hub_settings ChronicleFeed#azure_event_hub_settings}
	AzureEventHubSettings *ChronicleFeedDetailsAzureEventHubSettings `field:"optional" json:"azureEventHubSettings" yaml:"azureEventHubSettings"`
	// azure_mdm_intune_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#azure_mdm_intune_settings ChronicleFeed#azure_mdm_intune_settings}
	AzureMdmIntuneSettings *ChronicleFeedDetailsAzureMdmIntuneSettings `field:"optional" json:"azureMdmIntuneSettings" yaml:"azureMdmIntuneSettings"`
	// cloud_passage_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#cloud_passage_settings ChronicleFeed#cloud_passage_settings}
	CloudPassageSettings *ChronicleFeedDetailsCloudPassageSettings `field:"optional" json:"cloudPassageSettings" yaml:"cloudPassageSettings"`
	// cortex_xdr_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#cortex_xdr_settings ChronicleFeed#cortex_xdr_settings}
	CortexXdrSettings *ChronicleFeedDetailsCortexXdrSettings `field:"optional" json:"cortexXdrSettings" yaml:"cortexXdrSettings"`
	// crowdstrike_alerts_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#crowdstrike_alerts_settings ChronicleFeed#crowdstrike_alerts_settings}
	CrowdstrikeAlertsSettings *ChronicleFeedDetailsCrowdstrikeAlertsSettings `field:"optional" json:"crowdstrikeAlertsSettings" yaml:"crowdstrikeAlertsSettings"`
	// crowdstrike_detects_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#crowdstrike_detects_settings ChronicleFeed#crowdstrike_detects_settings}
	CrowdstrikeDetectsSettings *ChronicleFeedDetailsCrowdstrikeDetectsSettings `field:"optional" json:"crowdstrikeDetectsSettings" yaml:"crowdstrikeDetectsSettings"`
	// dummy_log_type_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#dummy_log_type_settings ChronicleFeed#dummy_log_type_settings}
	DummyLogTypeSettings *ChronicleFeedDetailsDummyLogTypeSettings `field:"optional" json:"dummyLogTypeSettings" yaml:"dummyLogTypeSettings"`
	// duo_auth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#duo_auth_settings ChronicleFeed#duo_auth_settings}
	DuoAuthSettings *ChronicleFeedDetailsDuoAuthSettings `field:"optional" json:"duoAuthSettings" yaml:"duoAuthSettings"`
	// duo_user_context_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#duo_user_context_settings ChronicleFeed#duo_user_context_settings}
	DuoUserContextSettings *ChronicleFeedDetailsDuoUserContextSettings `field:"optional" json:"duoUserContextSettings" yaml:"duoUserContextSettings"`
	// Source Type of the feed.
	//
	// Possible values:
	// GOOGLE_CLOUD_STORAGE
	// HTTP
	// SFTP
	// AMAZON_S3
	// AZURE_BLOBSTORE
	// API
	// AMAZON_SQS
	// PUBSUB
	// AMAZON_KINESIS_FIREHOSE
	// WEBHOOK
	// HTTPS_PUSH_GOOGLE_CLOUD_PUBSUB
	// HTTPS_PUSH_AMAZON_KINESIS_FIREHOSE
	// HTTPS_PUSH_WEBHOOK
	// AZURE_EVENT_HUB
	// GOOGLE_CLOUD_STORAGE_V2
	// AMAZON_S3_V2
	// AMAZON_SQS_V2
	// AZURE_BLOBSTORE_V2
	// GOOGLE_CLOUD_STORAGE_EVENT_DRIVEN Possible values: ["GOOGLE_CLOUD_STORAGE", "HTTP", "SFTP", "AMAZON_S3", "AZURE_BLOBSTORE", "API", "AMAZON_SQS", "PUBSUB", "AMAZON_KINESIS_FIREHOSE", "WEBHOOK", "HTTPS_PUSH_GOOGLE_CLOUD_PUBSUB", "HTTPS_PUSH_AMAZON_KINESIS_FIREHOSE", "HTTPS_PUSH_WEBHOOK", "AZURE_EVENT_HUB", "GOOGLE_CLOUD_STORAGE_V2", "AMAZON_S3_V2", "AMAZON_SQS_V2", "AZURE_BLOBSTORE_V2", "GOOGLE_CLOUD_STORAGE_EVENT_DRIVEN"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#feed_source_type ChronicleFeed#feed_source_type}
	FeedSourceType *string `field:"optional" json:"feedSourceType" yaml:"feedSourceType"`
	// fox_it_stix_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#fox_it_stix_settings ChronicleFeed#fox_it_stix_settings}
	FoxItStixSettings *ChronicleFeedDetailsFoxItStixSettings `field:"optional" json:"foxItStixSettings" yaml:"foxItStixSettings"`
	// gcs_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#gcs_settings ChronicleFeed#gcs_settings}
	GcsSettings *ChronicleFeedDetailsGcsSettings `field:"optional" json:"gcsSettings" yaml:"gcsSettings"`
	// gcs_v2_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#gcs_v2_settings ChronicleFeed#gcs_v2_settings}
	GcsV2Settings *ChronicleFeedDetailsGcsV2Settings `field:"optional" json:"gcsV2Settings" yaml:"gcsV2Settings"`
	// google_cloud_identity_devices_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#google_cloud_identity_devices_settings ChronicleFeed#google_cloud_identity_devices_settings}
	GoogleCloudIdentityDevicesSettings *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettings `field:"optional" json:"googleCloudIdentityDevicesSettings" yaml:"googleCloudIdentityDevicesSettings"`
	// google_cloud_identity_device_users_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#google_cloud_identity_device_users_settings ChronicleFeed#google_cloud_identity_device_users_settings}
	GoogleCloudIdentityDeviceUsersSettings *ChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettings `field:"optional" json:"googleCloudIdentityDeviceUsersSettings" yaml:"googleCloudIdentityDeviceUsersSettings"`
	// google_cloud_storage_event_driven_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#google_cloud_storage_event_driven_settings ChronicleFeed#google_cloud_storage_event_driven_settings}
	GoogleCloudStorageEventDrivenSettings *ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings `field:"optional" json:"googleCloudStorageEventDrivenSettings" yaml:"googleCloudStorageEventDrivenSettings"`
	// http_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#http_settings ChronicleFeed#http_settings}
	HttpSettings *ChronicleFeedDetailsHttpSettings `field:"optional" json:"httpSettings" yaml:"httpSettings"`
	// https_push_amazon_kinesis_firehose_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#https_push_amazon_kinesis_firehose_settings ChronicleFeed#https_push_amazon_kinesis_firehose_settings}
	HttpsPushAmazonKinesisFirehoseSettings *ChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettings `field:"optional" json:"httpsPushAmazonKinesisFirehoseSettings" yaml:"httpsPushAmazonKinesisFirehoseSettings"`
	// https_push_google_cloud_pubsub_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#https_push_google_cloud_pubsub_settings ChronicleFeed#https_push_google_cloud_pubsub_settings}
	HttpsPushGoogleCloudPubsubSettings *ChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettings `field:"optional" json:"httpsPushGoogleCloudPubsubSettings" yaml:"httpsPushGoogleCloudPubsubSettings"`
	// https_push_webhook_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#https_push_webhook_settings ChronicleFeed#https_push_webhook_settings}
	HttpsPushWebhookSettings *ChronicleFeedDetailsHttpsPushWebhookSettings `field:"optional" json:"httpsPushWebhookSettings" yaml:"httpsPushWebhookSettings"`
	// imperva_waf_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#imperva_waf_settings ChronicleFeed#imperva_waf_settings}
	ImpervaWafSettings *ChronicleFeedDetailsImpervaWafSettings `field:"optional" json:"impervaWafSettings" yaml:"impervaWafSettings"`
	// The ingestion metadata labels to apply to all logs ingested through this feed, and the resulting normalized data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#labels ChronicleFeed#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// mandiant_ioc_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#mandiant_ioc_settings ChronicleFeed#mandiant_ioc_settings}
	MandiantIocSettings *ChronicleFeedDetailsMandiantIocSettings `field:"optional" json:"mandiantIocSettings" yaml:"mandiantIocSettings"`
	// microsoft_graph_alert_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#microsoft_graph_alert_settings ChronicleFeed#microsoft_graph_alert_settings}
	MicrosoftGraphAlertSettings *ChronicleFeedDetailsMicrosoftGraphAlertSettings `field:"optional" json:"microsoftGraphAlertSettings" yaml:"microsoftGraphAlertSettings"`
	// microsoft_security_center_alert_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#microsoft_security_center_alert_settings ChronicleFeed#microsoft_security_center_alert_settings}
	MicrosoftSecurityCenterAlertSettings *ChronicleFeedDetailsMicrosoftSecurityCenterAlertSettings `field:"optional" json:"microsoftSecurityCenterAlertSettings" yaml:"microsoftSecurityCenterAlertSettings"`
	// mimecast_mail_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#mimecast_mail_settings ChronicleFeed#mimecast_mail_settings}
	MimecastMailSettings *ChronicleFeedDetailsMimecastMailSettings `field:"optional" json:"mimecastMailSettings" yaml:"mimecastMailSettings"`
	// mimecast_mail_v2_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#mimecast_mail_v2_settings ChronicleFeed#mimecast_mail_v2_settings}
	MimecastMailV2Settings *ChronicleFeedDetailsMimecastMailV2Settings `field:"optional" json:"mimecastMailV2Settings" yaml:"mimecastMailV2Settings"`
	// netskope_alert_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#netskope_alert_settings ChronicleFeed#netskope_alert_settings}
	NetskopeAlertSettings *ChronicleFeedDetailsNetskopeAlertSettings `field:"optional" json:"netskopeAlertSettings" yaml:"netskopeAlertSettings"`
	// netskope_alert_v2_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#netskope_alert_v2_settings ChronicleFeed#netskope_alert_v2_settings}
	NetskopeAlertV2Settings *ChronicleFeedDetailsNetskopeAlertV2Settings `field:"optional" json:"netskopeAlertV2Settings" yaml:"netskopeAlertV2Settings"`
	// office365_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#office365_settings ChronicleFeed#office365_settings}
	Office365Settings *ChronicleFeedDetailsOffice365Settings `field:"optional" json:"office365Settings" yaml:"office365Settings"`
	// okta_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#okta_settings ChronicleFeed#okta_settings}
	OktaSettings *ChronicleFeedDetailsOktaSettings `field:"optional" json:"oktaSettings" yaml:"oktaSettings"`
	// okta_user_context_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#okta_user_context_settings ChronicleFeed#okta_user_context_settings}
	OktaUserContextSettings *ChronicleFeedDetailsOktaUserContextSettings `field:"optional" json:"oktaUserContextSettings" yaml:"oktaUserContextSettings"`
	// pan_ioc_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#pan_ioc_settings ChronicleFeed#pan_ioc_settings}
	PanIocSettings *ChronicleFeedDetailsPanIocSettings `field:"optional" json:"panIocSettings" yaml:"panIocSettings"`
	// pan_prisma_cloud_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#pan_prisma_cloud_settings ChronicleFeed#pan_prisma_cloud_settings}
	PanPrismaCloudSettings *ChronicleFeedDetailsPanPrismaCloudSettings `field:"optional" json:"panPrismaCloudSettings" yaml:"panPrismaCloudSettings"`
	// proofpoint_mail_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#proofpoint_mail_settings ChronicleFeed#proofpoint_mail_settings}
	ProofpointMailSettings *ChronicleFeedDetailsProofpointMailSettings `field:"optional" json:"proofpointMailSettings" yaml:"proofpointMailSettings"`
	// proofpoint_on_demand_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#proofpoint_on_demand_settings ChronicleFeed#proofpoint_on_demand_settings}
	ProofpointOnDemandSettings *ChronicleFeedDetailsProofpointOnDemandSettings `field:"optional" json:"proofpointOnDemandSettings" yaml:"proofpointOnDemandSettings"`
	// pubsub_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#pubsub_settings ChronicleFeed#pubsub_settings}
	PubsubSettings *ChronicleFeedDetailsPubsubSettings `field:"optional" json:"pubsubSettings" yaml:"pubsubSettings"`
	// qualys_scan_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#qualys_scan_settings ChronicleFeed#qualys_scan_settings}
	QualysScanSettings *ChronicleFeedDetailsQualysScanSettings `field:"optional" json:"qualysScanSettings" yaml:"qualysScanSettings"`
	// qualys_vm_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#qualys_vm_settings ChronicleFeed#qualys_vm_settings}
	QualysVmSettings *ChronicleFeedDetailsQualysVmSettings `field:"optional" json:"qualysVmSettings" yaml:"qualysVmSettings"`
	// rapid7_insight_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#rapid7_insight_settings ChronicleFeed#rapid7_insight_settings}
	Rapid7InsightSettings *ChronicleFeedDetailsRapid7InsightSettings `field:"optional" json:"rapid7InsightSettings" yaml:"rapid7InsightSettings"`
	// recorded_future_ioc_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#recorded_future_ioc_settings ChronicleFeed#recorded_future_ioc_settings}
	RecordedFutureIocSettings *ChronicleFeedDetailsRecordedFutureIocSettings `field:"optional" json:"recordedFutureIocSettings" yaml:"recordedFutureIocSettings"`
	// rh_isac_ioc_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#rh_isac_ioc_settings ChronicleFeed#rh_isac_ioc_settings}
	RhIsacIocSettings *ChronicleFeedDetailsRhIsacIocSettings `field:"optional" json:"rhIsacIocSettings" yaml:"rhIsacIocSettings"`
	// salesforce_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#salesforce_settings ChronicleFeed#salesforce_settings}
	SalesforceSettings *ChronicleFeedDetailsSalesforceSettings `field:"optional" json:"salesforceSettings" yaml:"salesforceSettings"`
	// sentinelone_alert_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#sentinelone_alert_settings ChronicleFeed#sentinelone_alert_settings}
	SentineloneAlertSettings *ChronicleFeedDetailsSentineloneAlertSettings `field:"optional" json:"sentineloneAlertSettings" yaml:"sentineloneAlertSettings"`
	// service_now_cmdb_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#service_now_cmdb_settings ChronicleFeed#service_now_cmdb_settings}
	ServiceNowCmdbSettings *ChronicleFeedDetailsServiceNowCmdbSettings `field:"optional" json:"serviceNowCmdbSettings" yaml:"serviceNowCmdbSettings"`
	// sftp_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#sftp_settings ChronicleFeed#sftp_settings}
	SftpSettings *ChronicleFeedDetailsSftpSettings `field:"optional" json:"sftpSettings" yaml:"sftpSettings"`
	// symantec_event_export_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#symantec_event_export_settings ChronicleFeed#symantec_event_export_settings}
	SymantecEventExportSettings *ChronicleFeedDetailsSymantecEventExportSettings `field:"optional" json:"symantecEventExportSettings" yaml:"symantecEventExportSettings"`
	// thinkst_canary_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#thinkst_canary_settings ChronicleFeed#thinkst_canary_settings}
	ThinkstCanarySettings *ChronicleFeedDetailsThinkstCanarySettings `field:"optional" json:"thinkstCanarySettings" yaml:"thinkstCanarySettings"`
	// threat_connect_ioc_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#threat_connect_ioc_settings ChronicleFeed#threat_connect_ioc_settings}
	ThreatConnectIocSettings *ChronicleFeedDetailsThreatConnectIocSettings `field:"optional" json:"threatConnectIocSettings" yaml:"threatConnectIocSettings"`
	// threat_connect_ioc_v3_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#threat_connect_ioc_v3_settings ChronicleFeed#threat_connect_ioc_v3_settings}
	ThreatConnectIocV3Settings *ChronicleFeedDetailsThreatConnectIocV3Settings `field:"optional" json:"threatConnectIocV3Settings" yaml:"threatConnectIocV3Settings"`
	// trellix_hx_alerts_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#trellix_hx_alerts_settings ChronicleFeed#trellix_hx_alerts_settings}
	TrellixHxAlertsSettings *ChronicleFeedDetailsTrellixHxAlertsSettings `field:"optional" json:"trellixHxAlertsSettings" yaml:"trellixHxAlertsSettings"`
	// trellix_hx_bulk_acqs_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#trellix_hx_bulk_acqs_settings ChronicleFeed#trellix_hx_bulk_acqs_settings}
	TrellixHxBulkAcqsSettings *ChronicleFeedDetailsTrellixHxBulkAcqsSettings `field:"optional" json:"trellixHxBulkAcqsSettings" yaml:"trellixHxBulkAcqsSettings"`
	// trellix_hx_hosts_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#trellix_hx_hosts_settings ChronicleFeed#trellix_hx_hosts_settings}
	TrellixHxHostsSettings *ChronicleFeedDetailsTrellixHxHostsSettings `field:"optional" json:"trellixHxHostsSettings" yaml:"trellixHxHostsSettings"`
	// webhook_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#webhook_settings ChronicleFeed#webhook_settings}
	WebhookSettings *ChronicleFeedDetailsWebhookSettings `field:"optional" json:"webhookSettings" yaml:"webhookSettings"`
	// workday_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#workday_settings ChronicleFeed#workday_settings}
	WorkdaySettings *ChronicleFeedDetailsWorkdaySettings `field:"optional" json:"workdaySettings" yaml:"workdaySettings"`
	// workspace_activity_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#workspace_activity_settings ChronicleFeed#workspace_activity_settings}
	WorkspaceActivitySettings *ChronicleFeedDetailsWorkspaceActivitySettings `field:"optional" json:"workspaceActivitySettings" yaml:"workspaceActivitySettings"`
	// workspace_alerts_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#workspace_alerts_settings ChronicleFeed#workspace_alerts_settings}
	WorkspaceAlertsSettings *ChronicleFeedDetailsWorkspaceAlertsSettings `field:"optional" json:"workspaceAlertsSettings" yaml:"workspaceAlertsSettings"`
	// workspace_chrome_os_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#workspace_chrome_os_settings ChronicleFeed#workspace_chrome_os_settings}
	WorkspaceChromeOsSettings *ChronicleFeedDetailsWorkspaceChromeOsSettings `field:"optional" json:"workspaceChromeOsSettings" yaml:"workspaceChromeOsSettings"`
	// workspace_groups_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#workspace_groups_settings ChronicleFeed#workspace_groups_settings}
	WorkspaceGroupsSettings *ChronicleFeedDetailsWorkspaceGroupsSettings `field:"optional" json:"workspaceGroupsSettings" yaml:"workspaceGroupsSettings"`
	// workspace_mobile_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#workspace_mobile_settings ChronicleFeed#workspace_mobile_settings}
	WorkspaceMobileSettings *ChronicleFeedDetailsWorkspaceMobileSettings `field:"optional" json:"workspaceMobileSettings" yaml:"workspaceMobileSettings"`
	// workspace_privileges_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#workspace_privileges_settings ChronicleFeed#workspace_privileges_settings}
	WorkspacePrivilegesSettings *ChronicleFeedDetailsWorkspacePrivilegesSettings `field:"optional" json:"workspacePrivilegesSettings" yaml:"workspacePrivilegesSettings"`
	// workspace_users_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#workspace_users_settings ChronicleFeed#workspace_users_settings}
	WorkspaceUsersSettings *ChronicleFeedDetailsWorkspaceUsersSettings `field:"optional" json:"workspaceUsersSettings" yaml:"workspaceUsersSettings"`
}

