// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigClientToAgentConfig struct {
	// Required. The resource name of the Agent Gateway to use for inbound traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#agent_gateway VertexAiReasoningEngine#agent_gateway}
	AgentGateway *string `field:"required" json:"agentGateway" yaml:"agentGateway"`
}

