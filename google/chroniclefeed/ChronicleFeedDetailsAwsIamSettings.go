// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsAwsIamSettings struct {
	// Supported AWS IAM api type. Possible values: USERS ROLES GROUPS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/chronicle_feed#api_type ChronicleFeed#api_type}
	ApiType *string `field:"optional" json:"apiType" yaml:"apiType"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsAwsIamSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
}

