// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaievaluationitem


type VertexAiEvaluationItemEvaluationRequestRubricsRubrics struct {
	// content block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#content VertexAiEvaluationItem#content}
	Content *VertexAiEvaluationItemEvaluationRequestRubricsRubricsContent `field:"required" json:"content" yaml:"content"`
	// The relative importance of this rubric. Possible values: ["HIGH", "MEDIUM", "LOW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#importance VertexAiEvaluationItem#importance}
	Importance *string `field:"optional" json:"importance" yaml:"importance"`
	// Unique identifier for the rubric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#rubric_id VertexAiEvaluationItem#rubric_id}
	RubricId *string `field:"optional" json:"rubricId" yaml:"rubricId"`
	// A type designator for the rubric, e.g. 'SUMMARIZATION_QUALITY'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#type VertexAiEvaluationItem#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

