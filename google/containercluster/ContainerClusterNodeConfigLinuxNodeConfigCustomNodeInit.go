// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodeConfigLinuxNodeConfigCustomNodeInit struct {
	// init_script block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/container_cluster#init_script ContainerCluster#init_script}
	InitScript *ContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScript `field:"optional" json:"initScript" yaml:"initScript"`
}

