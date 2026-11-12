// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline


type VertexAiTrainingPipelineInputDataConfigBigqueryDestination struct {
	// The BigQuery table to which the output is to be written.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#output_uri VertexAiTrainingPipeline#output_uri}
	OutputUri *string `field:"required" json:"outputUri" yaml:"outputUri"`
}

