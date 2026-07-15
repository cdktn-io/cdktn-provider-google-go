// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolNodeDrainConfig struct {
	// The duration of the grace termination period for node drain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/container_cluster#grace_termination_duration ContainerCluster#grace_termination_duration}
	GraceTerminationDuration *string `field:"optional" json:"graceTerminationDuration" yaml:"graceTerminationDuration"`
	// The duration of the PDB timeout period for node drain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/container_cluster#pdb_timeout_duration ContainerCluster#pdb_timeout_duration}
	PdbTimeoutDuration *string `field:"optional" json:"pdbTimeoutDuration" yaml:"pdbTimeoutDuration"`
	// Whether to respect PodDisruptionBudget policy during node pool deletion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/container_cluster#respect_pdb_during_node_pool_deletion ContainerCluster#respect_pdb_during_node_pool_deletion}
	RespectPdbDuringNodePoolDeletion interface{} `field:"optional" json:"respectPdbDuringNodePoolDeletion" yaml:"respectPdbDuringNodePoolDeletion"`
}

