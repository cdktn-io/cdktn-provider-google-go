// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computenetworkfirewallpolicyiammember


type ComputeNetworkFirewallPolicyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/compute_network_firewall_policy_iam_member#expression ComputeNetworkFirewallPolicyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/compute_network_firewall_policy_iam_member#title ComputeNetworkFirewallPolicyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/compute_network_firewall_policy_iam_member#description ComputeNetworkFirewallPolicyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

