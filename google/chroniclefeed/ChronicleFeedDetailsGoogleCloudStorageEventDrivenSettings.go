// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings struct {
	// Google Cloud Storage Bucket URI for the feed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/chronicle_feed#bucket_uri ChronicleFeed#bucket_uri}
	BucketUri *string `field:"required" json:"bucketUri" yaml:"bucketUri"`
	// Subscription name for pubsub topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/chronicle_feed#pubsub_subscription ChronicleFeed#pubsub_subscription}
	PubsubSubscription *string `field:"required" json:"pubsubSubscription" yaml:"pubsubSubscription"`
	// Maximum File Age to ingest in days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/chronicle_feed#max_lookback_days ChronicleFeed#max_lookback_days}
	MaxLookbackDays *float64 `field:"optional" json:"maxLookbackDays" yaml:"maxLookbackDays"`
	// Possible values: NEVER ON_SUCCESS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/chronicle_feed#source_deletion_option ChronicleFeed#source_deletion_option}
	SourceDeletionOption *string `field:"optional" json:"sourceDeletionOption" yaml:"sourceDeletionOption"`
}

