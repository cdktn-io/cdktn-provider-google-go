// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsPanPrismaCloudSettingsAuthentication struct {
	// Password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#password ChronicleFeed#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#user ChronicleFeed#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

