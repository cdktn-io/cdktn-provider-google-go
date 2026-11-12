// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline


type VertexAiTrainingPipelineModelToUploadExplanationSpec struct {
	// Metadata describing the Model's input and output for explanation. This field uses JSON data as a string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#metadata VertexAiTrainingPipeline#metadata}
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// Parameters that configure explaining of the Model's predictions. This field uses JSON data as a string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#parameters VertexAiTrainingPipeline#parameters}
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

