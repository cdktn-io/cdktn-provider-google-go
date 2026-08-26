// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsAzureBlobStoreV2Settings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsAzureBlobStoreV2SettingsAuthentication `field:"required" json:"authentication" yaml:"authentication"`
	// Azure URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#azure_uri ChronicleFeed#azure_uri}
	AzureUri *string `field:"required" json:"azureUri" yaml:"azureUri"`
	// Maximum File Age to ingest in days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#max_lookback_days ChronicleFeed#max_lookback_days}
	MaxLookbackDays *float64 `field:"optional" json:"maxLookbackDays" yaml:"maxLookbackDays"`
	// Possible values: NEVER ON_SUCCESS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#source_deletion_option ChronicleFeed#source_deletion_option}
	SourceDeletionOption *string `field:"optional" json:"sourceDeletionOption" yaml:"sourceDeletionOption"`
}

