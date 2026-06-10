// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkconnectivityhubiammember


type NetworkConnectivityHubIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/network_connectivity_hub_iam_member#expression NetworkConnectivityHubIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/network_connectivity_hub_iam_member#title NetworkConnectivityHubIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/network_connectivity_hub_iam_member#description NetworkConnectivityHubIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

