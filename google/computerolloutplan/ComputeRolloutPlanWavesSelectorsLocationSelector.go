// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computerolloutplan


type ComputeRolloutPlanWavesSelectorsLocationSelector struct {
	// Example: "us-central1-a".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_rollout_plan#included_locations ComputeRolloutPlan#included_locations}
	IncludedLocations *[]*string `field:"optional" json:"includedLocations" yaml:"includedLocations"`
}

