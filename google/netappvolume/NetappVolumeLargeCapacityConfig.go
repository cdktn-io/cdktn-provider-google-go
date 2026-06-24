// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package netappvolume


type NetappVolumeLargeCapacityConfig struct {
	// The number of internal constituents (e.g., FlexVols) for this large volume. The minimum number of constituents is 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/netapp_volume#constituent_count NetappVolume#constituent_count}
	ConstituentCount *float64 `field:"optional" json:"constituentCount" yaml:"constituentCount"`
}

