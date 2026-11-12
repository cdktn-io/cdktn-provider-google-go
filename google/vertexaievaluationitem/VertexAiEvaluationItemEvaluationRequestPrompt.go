// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaievaluationitem


type VertexAiEvaluationItemEvaluationRequestPrompt struct {
	// Represents the complete execution trace of a multi-turn conversation, which can involve single or multiple agents, as a JSON string.
	//
	// This
	// serves as the input context for agent scraping.
	// At most one of 'text', 'value', 'prompt_template_data', 'agent_data', or 'user_scenario' can be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#agent_data VertexAiEvaluationItem#agent_data}
	AgentData *string `field:"optional" json:"agentData" yaml:"agentData"`
	// prompt_template_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#prompt_template_data VertexAiEvaluationItem#prompt_template_data}
	PromptTemplateData *VertexAiEvaluationItemEvaluationRequestPromptPromptTemplateData `field:"optional" json:"promptTemplateData" yaml:"promptTemplateData"`
	// Text prompt. At most one of 'text', 'value', 'prompt_template_data', 'agent_data', or 'user_scenario' can be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#text VertexAiEvaluationItem#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
	// user_scenario block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#user_scenario VertexAiEvaluationItem#user_scenario}
	UserScenario *VertexAiEvaluationItemEvaluationRequestPromptUserScenario `field:"optional" json:"userScenario" yaml:"userScenario"`
	// Fields and values that can be used to populate the prompt template, as a JSON string.
	//
	// At most one of 'text', 'value', 'prompt_template_data', 'agent_data', or 'user_scenario' can be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#value VertexAiEvaluationItem#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

