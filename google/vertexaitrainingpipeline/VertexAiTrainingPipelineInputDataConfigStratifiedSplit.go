// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline


type VertexAiTrainingPipelineInputDataConfigStratifiedSplit struct {
	// The key is a name of one of the Dataset's data columns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#key VertexAiTrainingPipeline#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The fraction of the input data that is to be used to evaluate the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#test_fraction VertexAiTrainingPipeline#test_fraction}
	TestFraction *float64 `field:"optional" json:"testFraction" yaml:"testFraction"`
	// The fraction of the input data that is to be used to train the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#training_fraction VertexAiTrainingPipeline#training_fraction}
	TrainingFraction *float64 `field:"optional" json:"trainingFraction" yaml:"trainingFraction"`
	// The fraction of the input data that is to be used to validate the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#validation_fraction VertexAiTrainingPipeline#validation_fraction}
	ValidationFraction *float64 `field:"optional" json:"validationFraction" yaml:"validationFraction"`
}

