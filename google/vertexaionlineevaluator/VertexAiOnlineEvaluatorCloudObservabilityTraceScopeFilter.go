// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaionlineevaluator


type VertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilter struct {
	// duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/vertex_ai_online_evaluator#duration VertexAiOnlineEvaluator#duration}
	Duration *VertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterDuration `field:"optional" json:"duration" yaml:"duration"`
	// total_token_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/vertex_ai_online_evaluator#total_token_usage VertexAiOnlineEvaluator#total_token_usage}
	TotalTokenUsage *VertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterTotalTokenUsage `field:"optional" json:"totalTokenUsage" yaml:"totalTokenUsage"`
}

