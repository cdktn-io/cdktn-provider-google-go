// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsAzureBlobStoreV2SettingsAuthentication struct {
	// Access Key also known as shared key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#access_key ChronicleFeed#access_key}
	AccessKey *string `field:"required" json:"accessKey" yaml:"accessKey"`
	// azure_v2_workload_identity_federation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#azure_v2_workload_identity_federation ChronicleFeed#azure_v2_workload_identity_federation}
	AzureV2WorkloadIdentityFederation *ChronicleFeedDetailsAzureBlobStoreV2SettingsAuthenticationAzureV2WorkloadIdentityFederation `field:"required" json:"azureV2WorkloadIdentityFederation" yaml:"azureV2WorkloadIdentityFederation"`
	// SAS Token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#sas_token ChronicleFeed#sas_token}
	SasToken *string `field:"required" json:"sasToken" yaml:"sasToken"`
}

