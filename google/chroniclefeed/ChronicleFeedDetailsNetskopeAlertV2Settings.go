// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsNetskopeAlertV2Settings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsNetskopeAlertV2SettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Content Category.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/chronicle_feed#content_category ChronicleFeed#content_category}
	ContentCategory *string `field:"optional" json:"contentCategory" yaml:"contentCategory"`
	// Content type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/chronicle_feed#content_types ChronicleFeed#content_types}
	ContentTypes *[]*string `field:"optional" json:"contentTypes" yaml:"contentTypes"`
	// API Hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/chronicle_feed#hostname ChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
}

