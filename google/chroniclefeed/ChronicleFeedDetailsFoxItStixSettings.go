// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsFoxItStixSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsFoxItStixSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Collection available at the poll service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#collection ChronicleFeed#collection}
	Collection *string `field:"optional" json:"collection" yaml:"collection"`
	// TAXII poll service URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#poll_service_uri ChronicleFeed#poll_service_uri}
	PollServiceUri *string `field:"optional" json:"pollServiceUri" yaml:"pollServiceUri"`
	// ssl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#ssl ChronicleFeed#ssl}
	Ssl *ChronicleFeedDetailsFoxItStixSettingsSsl `field:"optional" json:"ssl" yaml:"ssl"`
}

