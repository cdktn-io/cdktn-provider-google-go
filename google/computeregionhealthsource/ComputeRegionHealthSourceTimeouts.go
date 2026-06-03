// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeregionhealthsource


type ComputeRegionHealthSourceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/compute_region_health_source#create ComputeRegionHealthSource#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/compute_region_health_source#delete ComputeRegionHealthSource#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/compute_region_health_source#update ComputeRegionHealthSource#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

