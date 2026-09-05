// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsCrowdstrikeAlertsSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsCrowdstrikeAlertsSettingsAuthentication `field:"required" json:"authentication" yaml:"authentication"`
	// API Hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#hostname ChronicleFeed#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Ingestion Type. Possible values: BRING_ALL_ALERTS BRING_ONLY_NEW_ALERTS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#ingestion_type ChronicleFeed#ingestion_type}
	IngestionType *string `field:"optional" json:"ingestionType" yaml:"ingestionType"`
}

