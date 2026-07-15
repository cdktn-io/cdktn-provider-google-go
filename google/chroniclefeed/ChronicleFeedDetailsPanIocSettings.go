// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsPanIocSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsPanIocSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// PAN IOC feed name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#feed ChronicleFeed#feed}
	Feed *string `field:"optional" json:"feed" yaml:"feed"`
	// PAN IOC feed ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#feed_id ChronicleFeed#feed_id}
	FeedId *string `field:"optional" json:"feedId" yaml:"feedId"`
}

