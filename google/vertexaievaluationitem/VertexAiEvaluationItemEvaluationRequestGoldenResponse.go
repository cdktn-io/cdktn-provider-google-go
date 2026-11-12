// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaievaluationitem


type VertexAiEvaluationItemEvaluationRequestGoldenResponse struct {
	// The name of the candidate that produced the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#candidate VertexAiEvaluationItem#candidate}
	Candidate *string `field:"required" json:"candidate" yaml:"candidate"`
	// Represents the complete execution trace of a multi-turn conversation, which can involve single or multiple agents, as a JSON string.
	//
	// This
	// field is used to provide the full output of an agent's run, including
	// all turns and events, for direct evaluation.
	// At most one of 'text', 'value', or 'agent_data' can be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#agent_data VertexAiEvaluationItem#agent_data}
	AgentData *string `field:"optional" json:"agentData" yaml:"agentData"`
	// Text response. At most one of 'text', 'value', or 'agent_data' can be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#text VertexAiEvaluationItem#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
	// Fields and values that can be used to populate the response template, as a JSON string.
	//
	// At most one of 'text', 'value', or 'agent_data' can be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#value VertexAiEvaluationItem#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

