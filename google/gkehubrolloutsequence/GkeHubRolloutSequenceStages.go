// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gkehubrolloutsequence


type GkeHubRolloutSequenceStages struct {
	// List of Fleet projects to select the clusters from. Expected format: projects/{project}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/gke_hub_rollout_sequence#fleet_projects GkeHubRolloutSequence#fleet_projects}
	FleetProjects *[]*string `field:"required" json:"fleetProjects" yaml:"fleetProjects"`
	// cluster_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/gke_hub_rollout_sequence#cluster_selector GkeHubRolloutSequence#cluster_selector}
	ClusterSelector *GkeHubRolloutSequenceStagesClusterSelector `field:"optional" json:"clusterSelector" yaml:"clusterSelector"`
	// Soak time after upgrading all the clusters in the stage. Has to be specified in seconds, minutes, hours or days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/gke_hub_rollout_sequence#soak_duration GkeHubRolloutSequence#soak_duration}
	SoakDuration *string `field:"optional" json:"soakDuration" yaml:"soakDuration"`
}

