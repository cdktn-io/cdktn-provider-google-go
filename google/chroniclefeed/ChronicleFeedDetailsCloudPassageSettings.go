// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsCloudPassageSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsCloudPassageSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Event types filter for the events API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#event_types ChronicleFeed#event_types}
	EventTypes *[]*string `field:"optional" json:"eventTypes" yaml:"eventTypes"`
}

