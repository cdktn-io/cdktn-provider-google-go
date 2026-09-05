// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsAzureBlobStoreSettingsAuthentication struct {
	// SAS Token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#sas_token ChronicleFeed#sas_token}
	SasToken *string `field:"optional" json:"sasToken" yaml:"sasToken"`
	// Shared Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#shared_key ChronicleFeed#shared_key}
	SharedKey *string `field:"optional" json:"sharedKey" yaml:"sharedKey"`
}

