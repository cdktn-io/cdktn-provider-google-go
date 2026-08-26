// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iapagentregistryagentiammember


type IapAgentRegistryAgentIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/iap_agent_registry_agent_iam_member#expression IapAgentRegistryAgentIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/iap_agent_registry_agent_iam_member#title IapAgentRegistryAgentIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/iap_agent_registry_agent_iam_member#description IapAgentRegistryAgentIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

