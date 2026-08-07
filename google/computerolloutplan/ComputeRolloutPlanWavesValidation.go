// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computerolloutplan


type ComputeRolloutPlanWavesValidation struct {
	// The type of the validation.
	//
	// Possible values:
	// "manual": The system waits for an end-user approval API before progressing to the next wave.
	// "time": The system waits for a user specified duration before progressing to the next wave.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/compute_rollout_plan#type ComputeRolloutPlan#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// time_based_validation_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/compute_rollout_plan#time_based_validation_metadata ComputeRolloutPlan#time_based_validation_metadata}
	TimeBasedValidationMetadata *ComputeRolloutPlanWavesValidationTimeBasedValidationMetadata `field:"optional" json:"timeBasedValidationMetadata" yaml:"timeBasedValidationMetadata"`
}

