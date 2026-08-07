// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsAmazonS3V2SettingsAuthentication struct {
	// access_key_secret_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/chronicle_feed#access_key_secret_auth ChronicleFeed#access_key_secret_auth}
	AccessKeySecretAuth *ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAccessKeySecretAuth `field:"optional" json:"accessKeySecretAuth" yaml:"accessKeySecretAuth"`
	// aws_iam_role_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/chronicle_feed#aws_iam_role_auth ChronicleFeed#aws_iam_role_auth}
	AwsIamRoleAuth *ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuth `field:"optional" json:"awsIamRoleAuth" yaml:"awsIamRoleAuth"`
}

