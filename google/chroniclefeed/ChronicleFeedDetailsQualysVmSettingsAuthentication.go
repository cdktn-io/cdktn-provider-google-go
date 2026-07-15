// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsQualysVmSettingsAuthentication struct {
	// Secret of the account identified by user_name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#secret ChronicleFeed#secret}
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
	// Username of an identity used for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_feed#user ChronicleFeed#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

