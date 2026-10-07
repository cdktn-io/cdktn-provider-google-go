// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfig struct {
	// agent_to_anywhere_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#agent_to_anywhere_config VertexAiReasoningEngine#agent_to_anywhere_config}
	AgentToAnywhereConfig *VertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigAgentToAnywhereConfig `field:"optional" json:"agentToAnywhereConfig" yaml:"agentToAnywhereConfig"`
	// client_to_agent_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#client_to_agent_config VertexAiReasoningEngine#client_to_agent_config}
	ClientToAgentConfig *VertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigClientToAgentConfig `field:"optional" json:"clientToAgentConfig" yaml:"clientToAgentConfig"`
}

