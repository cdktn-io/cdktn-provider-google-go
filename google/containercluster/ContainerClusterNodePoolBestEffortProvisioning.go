// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolBestEffortProvisioning struct {
	// Whether node pool creation ignores non-fatal errors like stockout to provision as many nodes as possible, eventually bringing up the target number of nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Minimum number of nodes that must be provisioned for the creation to be considered successful.
	//
	// Can only be set when enabled is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/container_cluster#min_provision_nodes ContainerCluster#min_provision_nodes}
	MinProvisionNodes *float64 `field:"optional" json:"minProvisionNodes" yaml:"minProvisionNodes"`
}

