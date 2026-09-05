// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsMimecastMailV2Settings struct {
	// auth_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#auth_credentials ChronicleFeed#auth_credentials}
	AuthCredentials *ChronicleFeedDetailsMimecastMailV2SettingsAuthCredentials `field:"optional" json:"authCredentials" yaml:"authCredentials"`
}

