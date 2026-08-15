// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeregionnetworkfirewallpolicyiammember


type ComputeRegionNetworkFirewallPolicyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_region_network_firewall_policy_iam_member#expression ComputeRegionNetworkFirewallPolicyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_region_network_firewall_policy_iam_member#title ComputeRegionNetworkFirewallPolicyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_region_network_firewall_policy_iam_member#description ComputeRegionNetworkFirewallPolicyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

