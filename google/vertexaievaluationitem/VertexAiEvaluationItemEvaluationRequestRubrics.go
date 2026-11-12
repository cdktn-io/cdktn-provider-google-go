// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaievaluationitem


type VertexAiEvaluationItemEvaluationRequestRubrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#name VertexAiEvaluationItem#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Human-readable name for the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#display_name VertexAiEvaluationItem#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Unique identifier for the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#group_id VertexAiEvaluationItem#group_id}
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// rubrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#rubrics VertexAiEvaluationItem#rubrics}
	Rubrics interface{} `field:"optional" json:"rubrics" yaml:"rubrics"`
}

