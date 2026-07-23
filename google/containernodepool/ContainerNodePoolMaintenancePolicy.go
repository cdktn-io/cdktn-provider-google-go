// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolMaintenancePolicy struct {
	// exclusion_until_end_of_support block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/container_node_pool#exclusion_until_end_of_support ContainerNodePool#exclusion_until_end_of_support}
	ExclusionUntilEndOfSupport interface{} `field:"optional" json:"exclusionUntilEndOfSupport" yaml:"exclusionUntilEndOfSupport"`
}

