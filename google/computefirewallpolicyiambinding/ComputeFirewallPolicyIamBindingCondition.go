// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computefirewallpolicyiambinding


type ComputeFirewallPolicyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/compute_firewall_policy_iam_binding#expression ComputeFirewallPolicyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/compute_firewall_policy_iam_binding#title ComputeFirewallPolicyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/compute_firewall_policy_iam_binding#description ComputeFirewallPolicyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

