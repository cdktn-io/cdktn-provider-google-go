// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterFleet struct {
	// The type of the cluster's fleet membership.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/container_cluster#membership_type ContainerCluster#membership_type}
	MembershipType *string `field:"optional" json:"membershipType" yaml:"membershipType"`
	// The Fleet host project of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/container_cluster#project ContainerCluster#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

