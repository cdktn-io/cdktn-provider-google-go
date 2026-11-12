// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaievaluationitem


type VertexAiEvaluationItemEvaluationRequest struct {
	// candidate_responses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#candidate_responses VertexAiEvaluationItem#candidate_responses}
	CandidateResponses interface{} `field:"optional" json:"candidateResponses" yaml:"candidateResponses"`
	// golden_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#golden_response VertexAiEvaluationItem#golden_response}
	GoldenResponse *VertexAiEvaluationItemEvaluationRequestGoldenResponse `field:"optional" json:"goldenResponse" yaml:"goldenResponse"`
	// prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#prompt VertexAiEvaluationItem#prompt}
	Prompt *VertexAiEvaluationItemEvaluationRequestPrompt `field:"optional" json:"prompt" yaml:"prompt"`
	// rubrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#rubrics VertexAiEvaluationItem#rubrics}
	Rubrics interface{} `field:"optional" json:"rubrics" yaml:"rubrics"`
}

