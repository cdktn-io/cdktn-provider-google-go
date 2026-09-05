// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iapagentregistryendpointiambinding


type IapAgentRegistryEndpointIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/iap_agent_registry_endpoint_iam_binding#expression IapAgentRegistryEndpointIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/iap_agent_registry_endpoint_iam_binding#title IapAgentRegistryEndpointIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/iap_agent_registry_endpoint_iam_binding#description IapAgentRegistryEndpointIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

