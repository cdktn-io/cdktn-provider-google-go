// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsTrellixHxAlertsSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsTrellixHxAlertsSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Trellix HX Device URL.
	//
	// This must be a valid URL with an http or https scheme. It has no default.
	// Usually a device URL is in the form of either:
	// https://xxx.trellix.com/hx/id//
	// - or -
	// https://htapdeviceproxy.md.mandiant.net/dphb/hx//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/chronicle_feed#endpoint ChronicleFeed#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

