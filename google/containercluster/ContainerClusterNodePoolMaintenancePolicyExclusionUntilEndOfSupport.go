// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolMaintenancePolicyExclusionUntilEndOfSupport struct {
	// Whether to enable the maintenance exclusion until the end of support for this NodePool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

