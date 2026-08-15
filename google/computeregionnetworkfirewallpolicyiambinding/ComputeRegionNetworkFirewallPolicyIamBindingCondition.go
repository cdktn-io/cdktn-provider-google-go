// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeregionnetworkfirewallpolicyiambinding


type ComputeRegionNetworkFirewallPolicyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_region_network_firewall_policy_iam_binding#expression ComputeRegionNetworkFirewallPolicyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_region_network_firewall_policy_iam_binding#title ComputeRegionNetworkFirewallPolicyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_region_network_firewall_policy_iam_binding#description ComputeRegionNetworkFirewallPolicyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

