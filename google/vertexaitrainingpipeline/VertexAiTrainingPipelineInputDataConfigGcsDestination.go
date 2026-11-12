// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline


type VertexAiTrainingPipelineInputDataConfigGcsDestination struct {
	// The Cloud Storage URI prefix to write to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#output_uri_prefix VertexAiTrainingPipeline#output_uri_prefix}
	OutputUriPrefix *string `field:"required" json:"outputUriPrefix" yaml:"outputUriPrefix"`
}

