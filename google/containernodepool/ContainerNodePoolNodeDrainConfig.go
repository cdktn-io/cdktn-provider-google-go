// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNodeDrainConfig struct {
	// The duration of the grace termination period for node drain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/container_node_pool#grace_termination_duration ContainerNodePool#grace_termination_duration}
	GraceTerminationDuration *string `field:"optional" json:"graceTerminationDuration" yaml:"graceTerminationDuration"`
	// The duration of the PDB timeout period for node drain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/container_node_pool#pdb_timeout_duration ContainerNodePool#pdb_timeout_duration}
	PdbTimeoutDuration *string `field:"optional" json:"pdbTimeoutDuration" yaml:"pdbTimeoutDuration"`
	// Whether to respect PodDisruptionBudget policy during node pool deletion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/container_node_pool#respect_pdb_during_node_pool_deletion ContainerNodePool#respect_pdb_during_node_pool_deletion}
	RespectPdbDuringNodePoolDeletion interface{} `field:"optional" json:"respectPdbDuringNodePoolDeletion" yaml:"respectPdbDuringNodePoolDeletion"`
}

