// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsHttpsPushWebhookSettings struct {
	// Delimiter to split on for the feed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#split_delimiter ChronicleFeed#split_delimiter}
	SplitDelimiter *string `field:"optional" json:"splitDelimiter" yaml:"splitDelimiter"`
}

