// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolNodeConfigTaintConfig struct {
	// Architecture taint behavior. Controls, how we apply taints based on the node architecture.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/container_cluster#architecture_taint_behavior ContainerCluster#architecture_taint_behavior}
	ArchitectureTaintBehavior *string `field:"required" json:"architectureTaintBehavior" yaml:"architectureTaintBehavior"`
}

