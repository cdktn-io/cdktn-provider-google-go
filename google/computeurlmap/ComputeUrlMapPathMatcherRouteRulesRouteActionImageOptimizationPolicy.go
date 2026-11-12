// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeurlmap


type ComputeUrlMapPathMatcherRouteRulesRouteActionImageOptimizationPolicy struct {
	// Controls whether presets or control parameters specified as URL query parameters are used to optimize the image.
	//
	// If not specified, the default
	// value is DISABLED. Possible values: ["DISABLED", "ENABLED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/compute_url_map#query_parameter_interpretation ComputeUrlMap#query_parameter_interpretation}
	QueryParameterInterpretation *string `field:"optional" json:"queryParameterInterpretation" yaml:"queryParameterInterpretation"`
}

