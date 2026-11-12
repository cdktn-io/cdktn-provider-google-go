// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline


type VertexAiTrainingPipelineInputDataConfigPredefinedSplit struct {
	// The key is a name of one of the Dataset's data columns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#key VertexAiTrainingPipeline#key}
	Key *string `field:"required" json:"key" yaml:"key"`
}

