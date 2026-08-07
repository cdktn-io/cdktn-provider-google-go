// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computefirewallpolicyiammember


type ComputeFirewallPolicyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/compute_firewall_policy_iam_member#expression ComputeFirewallPolicyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/compute_firewall_policy_iam_member#title ComputeFirewallPolicyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/compute_firewall_policy_iam_member#description ComputeFirewallPolicyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

