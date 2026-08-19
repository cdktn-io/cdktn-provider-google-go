// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolAutoConfigLinuxNodeConfigNodeKernelModuleLoading struct {
	// The policy for kernel module loading.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/container_cluster#policy ContainerCluster#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

