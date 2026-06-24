// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsDuoAuthSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsDuoAuthSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// API Hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/chronicle_feed#hostname ChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
}

