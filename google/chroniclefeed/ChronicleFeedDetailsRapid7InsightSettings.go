// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsRapid7InsightSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsRapid7InsightSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Rapid7 API endpoint. Should be "vulnerabilities" or "assets".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#endpoint ChronicleFeed#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// API Hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#hostname ChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
}

