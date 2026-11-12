// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline


type VertexAiTrainingPipelineInputDataConfigFilterSplit struct {
	// A filter on DataItems of the Dataset for test set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#test_filter VertexAiTrainingPipeline#test_filter}
	TestFilter *string `field:"required" json:"testFilter" yaml:"testFilter"`
	// A filter on DataItems of the Dataset for training set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#training_filter VertexAiTrainingPipeline#training_filter}
	TrainingFilter *string `field:"required" json:"trainingFilter" yaml:"trainingFilter"`
	// A filter on DataItems of the Dataset for validation set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#validation_filter VertexAiTrainingPipeline#validation_filter}
	ValidationFilter *string `field:"required" json:"validationFilter" yaml:"validationFilter"`
}

