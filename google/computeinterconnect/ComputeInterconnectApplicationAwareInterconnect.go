// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeinterconnect


type ComputeInterconnectApplicationAwareInterconnect struct {
	// bandwidth_percentage_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/compute_interconnect#bandwidth_percentage_policy ComputeInterconnect#bandwidth_percentage_policy}
	BandwidthPercentagePolicy *ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicy `field:"optional" json:"bandwidthPercentagePolicy" yaml:"bandwidthPercentagePolicy"`
	// A description for the AAI profile on this interconnect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/compute_interconnect#profile_description ComputeInterconnect#profile_description}
	ProfileDescription *string `field:"optional" json:"profileDescription" yaml:"profileDescription"`
	// shape_average_percentage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/compute_interconnect#shape_average_percentage ComputeInterconnect#shape_average_percentage}
	ShapeAveragePercentage interface{} `field:"optional" json:"shapeAveragePercentage" yaml:"shapeAveragePercentage"`
	// strict_priority_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/compute_interconnect#strict_priority_policy ComputeInterconnect#strict_priority_policy}
	StrictPriorityPolicy *ComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicy `field:"optional" json:"strictPriorityPolicy" yaml:"strictPriorityPolicy"`
}

