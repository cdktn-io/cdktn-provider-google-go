// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolNodeConfigSandboxConfig struct {
	// Type of the sandbox to use for the node (e.g. 'GVISOR').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/container_cluster#type ContainerCluster#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

