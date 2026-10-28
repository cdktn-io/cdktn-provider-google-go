// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaionlineevaluator


type VertexAiOnlineEvaluatorConfigRandomSampling struct {
	// The percentage of traces to sample for evaluation. Must be an integer between 1 and 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_online_evaluator#percentage VertexAiOnlineEvaluator#percentage}
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
}

