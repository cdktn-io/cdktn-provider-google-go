// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsAmazonSqsSettingsAuthentication struct {
	// additional_s3_access_key_secret_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/chronicle_feed#additional_s3_access_key_secret_auth ChronicleFeed#additional_s3_access_key_secret_auth}
	AdditionalS3AccessKeySecretAuth *ChronicleFeedDetailsAmazonSqsSettingsAuthenticationAdditionalS3AccessKeySecretAuth `field:"optional" json:"additionalS3AccessKeySecretAuth" yaml:"additionalS3AccessKeySecretAuth"`
	// sqs_access_key_secret_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/chronicle_feed#sqs_access_key_secret_auth ChronicleFeed#sqs_access_key_secret_auth}
	SqsAccessKeySecretAuth *ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuth `field:"optional" json:"sqsAccessKeySecretAuth" yaml:"sqsAccessKeySecretAuth"`
}

