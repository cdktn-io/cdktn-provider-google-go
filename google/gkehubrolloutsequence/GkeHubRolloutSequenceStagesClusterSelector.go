// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gkehubrolloutsequence


type GkeHubRolloutSequenceStagesClusterSelector struct {
	// The label selector must be a valid CEL (Common Expression Language) expression which evaluates resource.labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/gke_hub_rollout_sequence#label_selector GkeHubRolloutSequence#label_selector}
	LabelSelector *string `field:"required" json:"labelSelector" yaml:"labelSelector"`
}

