// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaionlineevaluator


type VertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilter struct {
	// duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_online_evaluator#duration VertexAiOnlineEvaluator#duration}
	Duration *VertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterDuration `field:"optional" json:"duration" yaml:"duration"`
	// model_call_errors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_online_evaluator#model_call_errors VertexAiOnlineEvaluator#model_call_errors}
	ModelCallErrors *VertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCallErrors `field:"optional" json:"modelCallErrors" yaml:"modelCallErrors"`
	// model_calls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_online_evaluator#model_calls VertexAiOnlineEvaluator#model_calls}
	ModelCalls *VertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCalls `field:"optional" json:"modelCalls" yaml:"modelCalls"`
	// tool_call_errors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_online_evaluator#tool_call_errors VertexAiOnlineEvaluator#tool_call_errors}
	ToolCallErrors *VertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCallErrors `field:"optional" json:"toolCallErrors" yaml:"toolCallErrors"`
	// tool_calls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_online_evaluator#tool_calls VertexAiOnlineEvaluator#tool_calls}
	ToolCalls *VertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCalls `field:"optional" json:"toolCalls" yaml:"toolCalls"`
	// total_token_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_online_evaluator#total_token_usage VertexAiOnlineEvaluator#total_token_usage}
	TotalTokenUsage *VertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterTotalTokenUsage `field:"optional" json:"totalTokenUsage" yaml:"totalTokenUsage"`
	// user_turns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_online_evaluator#user_turns VertexAiOnlineEvaluator#user_turns}
	UserTurns *VertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterUserTurns `field:"optional" json:"userTurns" yaml:"userTurns"`
}

