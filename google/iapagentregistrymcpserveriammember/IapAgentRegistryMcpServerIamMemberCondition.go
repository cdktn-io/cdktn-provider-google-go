// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iapagentregistrymcpserveriammember


type IapAgentRegistryMcpServerIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/iap_agent_registry_mcp_server_iam_member#expression IapAgentRegistryMcpServerIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/iap_agent_registry_mcp_server_iam_member#title IapAgentRegistryMcpServerIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/iap_agent_registry_mcp_server_iam_member#description IapAgentRegistryMcpServerIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

