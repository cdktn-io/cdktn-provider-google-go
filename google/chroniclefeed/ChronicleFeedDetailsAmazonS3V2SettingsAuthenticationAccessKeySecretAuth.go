// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAccessKeySecretAuth struct {
	// Access Key ID for an AWS account (a 20-character, alphanumeric string).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/chronicle_feed#access_key_id ChronicleFeed#access_key_id}
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// Secret Access Key for an AWS account (a 40-character string).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/chronicle_feed#secret_access_key ChronicleFeed#secret_access_key}
	SecretAccessKey *string `field:"required" json:"secretAccessKey" yaml:"secretAccessKey"`
}

