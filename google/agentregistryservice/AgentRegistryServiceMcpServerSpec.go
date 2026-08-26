// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryservice


type AgentRegistryServiceMcpServerSpec struct {
	// The type of the MCP Server spec content. Possible values: ["NO_SPEC", "TOOL_SPEC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/agent_registry_service#type AgentRegistryService#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The content of the MCP Server spec. This payload is validated against the schema for the specified type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/agent_registry_service#content AgentRegistryService#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
}

