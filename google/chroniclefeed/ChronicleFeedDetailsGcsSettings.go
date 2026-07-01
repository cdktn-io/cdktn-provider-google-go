// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsGcsSettings struct {
	// Bucket URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/chronicle_feed#bucket_uri ChronicleFeed#bucket_uri}
	BucketUri *string `field:"optional" json:"bucketUri" yaml:"bucketUri"`
	// Possible values: SOURCE_DELETION_NEVER SOURCE_DELETION_ON_SUCCESS SOURCE_DELETION_ON_SUCCESS_FILES_ONLY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/chronicle_feed#source_deletion_option ChronicleFeed#source_deletion_option}
	SourceDeletionOption *string `field:"optional" json:"sourceDeletionOption" yaml:"sourceDeletionOption"`
	// Possible values: FILES FOLDERS FOLDERS_RECURSIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/chronicle_feed#source_type ChronicleFeed#source_type}
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

