// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsSftpSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsSftpSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Possible values: SOURCE_DELETION_NEVER SOURCE_DELETION_ON_SUCCESS SOURCE_DELETION_ON_SUCCESS_FILES_ONLY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/chronicle_feed#source_deletion_option ChronicleFeed#source_deletion_option}
	SourceDeletionOption *string `field:"optional" json:"sourceDeletionOption" yaml:"sourceDeletionOption"`
	// Possible values: FILES FOLDERS FOLDERS_RECURSIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/chronicle_feed#source_type ChronicleFeed#source_type}
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// SFTP URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/chronicle_feed#uri ChronicleFeed#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

