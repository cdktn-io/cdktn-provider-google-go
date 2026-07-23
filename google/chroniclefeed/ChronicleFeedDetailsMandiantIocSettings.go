// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsMandiantIocSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsMandiantIocSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// time since when to start fetching the IOCs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#start_time ChronicleFeed#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

