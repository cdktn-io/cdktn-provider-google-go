// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaievaluationitem


type VertexAiEvaluationItemEvaluationRequestPromptUserScenario struct {
	// The plan for the conversation, used to drive the multi-turn agent run and generate the simulated agent evaluation dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#conversation_plan VertexAiEvaluationItem#conversation_plan}
	ConversationPlan *string `field:"required" json:"conversationPlan" yaml:"conversationPlan"`
	// The prompt that starts the conversation between the simulated user and the agent under test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#starting_prompt VertexAiEvaluationItem#starting_prompt}
	StartingPrompt *string `field:"required" json:"startingPrompt" yaml:"startingPrompt"`
}

