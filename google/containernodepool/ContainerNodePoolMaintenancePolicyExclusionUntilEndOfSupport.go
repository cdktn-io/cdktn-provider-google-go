// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolMaintenancePolicyExclusionUntilEndOfSupport struct {
	// Whether to enable the maintenance exclusion until the end of support for this NodePool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/container_node_pool#enabled ContainerNodePool#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

