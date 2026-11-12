// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline


type VertexAiTrainingPipelineModelToUploadContainerSpecStartupProbe struct {
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#exec VertexAiTrainingPipeline#exec}
	Exec *VertexAiTrainingPipelineModelToUploadContainerSpecStartupProbeExec `field:"optional" json:"exec" yaml:"exec"`
	// How often (in seconds) to perform the probe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#period_seconds VertexAiTrainingPipeline#period_seconds}
	PeriodSeconds *float64 `field:"optional" json:"periodSeconds" yaml:"periodSeconds"`
	// Number of seconds after which the probe times out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#timeout_seconds VertexAiTrainingPipeline#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

