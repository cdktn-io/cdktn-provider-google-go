// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsWorkspacePrivilegesSettingsAuthenticationClaims struct {
	// Audience.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#audience ChronicleFeed#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// Issuer. Usually the client_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#issuer ChronicleFeed#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Subject. Usually the email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#subject ChronicleFeed#subject}
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}

