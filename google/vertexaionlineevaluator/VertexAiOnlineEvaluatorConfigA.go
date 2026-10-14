// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaionlineevaluator


type VertexAiOnlineEvaluatorConfigA struct {
	// The maximum number of evaluations to perform per run. If set to 0, the number is unbounded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/vertex_ai_online_evaluator#max_evaluated_samples_per_run VertexAiOnlineEvaluator#max_evaluated_samples_per_run}
	MaxEvaluatedSamplesPerRun *string `field:"optional" json:"maxEvaluatedSamplesPerRun" yaml:"maxEvaluatedSamplesPerRun"`
	// random_sampling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/vertex_ai_online_evaluator#random_sampling VertexAiOnlineEvaluator#random_sampling}
	RandomSampling *VertexAiOnlineEvaluatorConfigRandomSampling `field:"optional" json:"randomSampling" yaml:"randomSampling"`
}

