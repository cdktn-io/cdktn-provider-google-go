// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsTrellixHxHostsSettingsAuthentication struct {
	// msso block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#msso ChronicleFeed#msso}
	Msso *ChronicleFeedDetailsTrellixHxHostsSettingsAuthenticationMsso `field:"optional" json:"msso" yaml:"msso"`
	// trellix_iam block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#trellix_iam ChronicleFeed#trellix_iam}
	TrellixIam *ChronicleFeedDetailsTrellixHxHostsSettingsAuthenticationTrellixIam `field:"optional" json:"trellixIam" yaml:"trellixIam"`
}

