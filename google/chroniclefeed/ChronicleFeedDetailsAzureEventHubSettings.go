// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsAzureEventHubSettings struct {
	// Event hub consumer group to read from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#consumer_group ChronicleFeed#consumer_group}
	ConsumerGroup *string `field:"required" json:"consumerGroup" yaml:"consumerGroup"`
	// Event hub connection string for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#event_hub_connection_string ChronicleFeed#event_hub_connection_string}
	EventHubConnectionString *string `field:"required" json:"eventHubConnectionString" yaml:"eventHubConnectionString"`
	// Event hub to read from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#name ChronicleFeed#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// SAS token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#azure_sas_token ChronicleFeed#azure_sas_token}
	AzureSasToken *string `field:"optional" json:"azureSasToken" yaml:"azureSasToken"`
	// Blob store connection string for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#azure_storage_connection_string ChronicleFeed#azure_storage_connection_string}
	AzureStorageConnectionString *string `field:"optional" json:"azureStorageConnectionString" yaml:"azureStorageConnectionString"`
	// Blob storage container name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#azure_storage_container ChronicleFeed#azure_storage_container}
	AzureStorageContainer *string `field:"optional" json:"azureStorageContainer" yaml:"azureStorageContainer"`
}

