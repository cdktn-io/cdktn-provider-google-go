// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computenetworkfirewallpolicyiambinding


type ComputeNetworkFirewallPolicyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_network_firewall_policy_iam_binding#expression ComputeNetworkFirewallPolicyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_network_firewall_policy_iam_binding#title ComputeNetworkFirewallPolicyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_network_firewall_policy_iam_binding#description ComputeNetworkFirewallPolicyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

