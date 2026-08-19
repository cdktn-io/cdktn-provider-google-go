// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsAmazonSqsV2SettingsAuthentication struct {
	// aws_iam_role_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/chronicle_feed#aws_iam_role_auth ChronicleFeed#aws_iam_role_auth}
	AwsIamRoleAuth *ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuth `field:"required" json:"awsIamRoleAuth" yaml:"awsIamRoleAuth"`
	// sqs_v2_access_key_secret_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/chronicle_feed#sqs_v2_access_key_secret_auth ChronicleFeed#sqs_v2_access_key_secret_auth}
	SqsV2AccessKeySecretAuth *ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuth `field:"required" json:"sqsV2AccessKeySecretAuth" yaml:"sqsV2AccessKeySecretAuth"`
}

