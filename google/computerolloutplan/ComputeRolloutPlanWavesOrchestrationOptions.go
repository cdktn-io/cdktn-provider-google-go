// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computerolloutplan


type ComputeRolloutPlanWavesOrchestrationOptions struct {
	// delays block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/compute_rollout_plan#delays ComputeRolloutPlan#delays}
	Delays interface{} `field:"optional" json:"delays" yaml:"delays"`
	// Maximum number of locations to be orchestrated in parallel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/compute_rollout_plan#max_concurrent_locations ComputeRolloutPlan#max_concurrent_locations}
	MaxConcurrentLocations *float64 `field:"optional" json:"maxConcurrentLocations" yaml:"maxConcurrentLocations"`
	// Maximum number of resources to be orchestrated per location in parallel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/compute_rollout_plan#max_concurrent_resources_per_location ComputeRolloutPlan#max_concurrent_resources_per_location}
	MaxConcurrentResourcesPerLocation *float64 `field:"optional" json:"maxConcurrentResourcesPerLocation" yaml:"maxConcurrentResourcesPerLocation"`
}

