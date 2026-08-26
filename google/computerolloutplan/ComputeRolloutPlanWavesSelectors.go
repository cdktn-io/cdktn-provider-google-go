// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computerolloutplan


type ComputeRolloutPlanWavesSelectors struct {
	// location_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/compute_rollout_plan#location_selector ComputeRolloutPlan#location_selector}
	LocationSelector *ComputeRolloutPlanWavesSelectorsLocationSelector `field:"optional" json:"locationSelector" yaml:"locationSelector"`
	// resource_hierarchy_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/compute_rollout_plan#resource_hierarchy_selector ComputeRolloutPlan#resource_hierarchy_selector}
	ResourceHierarchySelector *ComputeRolloutPlanWavesSelectorsResourceHierarchySelector `field:"optional" json:"resourceHierarchySelector" yaml:"resourceHierarchySelector"`
}

