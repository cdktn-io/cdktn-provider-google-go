// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsAmazonS3Settings struct {
	// S3 URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#s3_uri ChronicleFeed#s3_uri}
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Possible values: SOURCE_DELETION_NEVER SOURCE_DELETION_ON_SUCCESS SOURCE_DELETION_ON_SUCCESS_FILES_ONLY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#source_deletion_option ChronicleFeed#source_deletion_option}
	SourceDeletionOption *string `field:"required" json:"sourceDeletionOption" yaml:"sourceDeletionOption"`
	// Possible values: FILES FOLDERS FOLDERS_RECURSIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#source_type ChronicleFeed#source_type}
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsAmazonS3SettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
}

