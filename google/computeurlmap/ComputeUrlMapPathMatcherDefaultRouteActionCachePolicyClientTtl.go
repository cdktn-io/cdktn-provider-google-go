// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeurlmap


type ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyClientTtl struct {
	// Span of time at a resolution of a second. Must be from 0 to 315,576,000,000 inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_url_map#seconds ComputeUrlMap#seconds}
	Seconds *string `field:"required" json:"seconds" yaml:"seconds"`
	// Span of time that's a fraction of a second at nanosecond resolution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_url_map#nanos ComputeUrlMap#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
}

