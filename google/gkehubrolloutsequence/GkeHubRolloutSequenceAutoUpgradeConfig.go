// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gkehubrolloutsequence


type GkeHubRolloutSequenceAutoUpgradeConfig struct {
	// rollout_creation_scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/gke_hub_rollout_sequence#rollout_creation_scope GkeHubRolloutSequence#rollout_creation_scope}
	RolloutCreationScope *GkeHubRolloutSequenceAutoUpgradeConfigRolloutCreationScope `field:"optional" json:"rolloutCreationScope" yaml:"rolloutCreationScope"`
}

