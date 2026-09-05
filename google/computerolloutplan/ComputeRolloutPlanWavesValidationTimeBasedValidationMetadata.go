// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computerolloutplan


type ComputeRolloutPlanWavesValidationTimeBasedValidationMetadata struct {
	// The duration that the system waits in between waves.
	//
	// This wait starts
	// after all changes in the wave are rolled out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_rollout_plan#wait_duration ComputeRolloutPlan#wait_duration}
	WaitDuration *string `field:"optional" json:"waitDuration" yaml:"waitDuration"`
}

