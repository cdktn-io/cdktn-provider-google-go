// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computerolloutplan


type ComputeRolloutPlanWaves struct {
	// selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/compute_rollout_plan#selectors ComputeRolloutPlan#selectors}
	Selectors interface{} `field:"required" json:"selectors" yaml:"selectors"`
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/compute_rollout_plan#validation ComputeRolloutPlan#validation}
	Validation *ComputeRolloutPlanWavesValidation `field:"required" json:"validation" yaml:"validation"`
	// The display name of this wave of the rollout plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/compute_rollout_plan#display_name ComputeRolloutPlan#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// orchestration_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/compute_rollout_plan#orchestration_options ComputeRolloutPlan#orchestration_options}
	OrchestrationOptions *ComputeRolloutPlanWavesOrchestrationOptions `field:"optional" json:"orchestrationOptions" yaml:"orchestrationOptions"`
}

