// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryservice


type AgentRegistryServiceAgentSpec struct {
	// The type of the Agent spec content. Possible values: ["NO_SPEC", "A2A_AGENT_CARD"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/agent_registry_service#type AgentRegistryService#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The content of the Agent spec in the JSON format.
	//
	// This payload is validated against the schema for the specified type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/agent_registry_service#content AgentRegistryService#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
}

