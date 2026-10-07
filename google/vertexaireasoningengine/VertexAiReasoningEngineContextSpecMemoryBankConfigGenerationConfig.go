// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfig struct {
	// The model used to generate memories. Format: projects/{project}/locations/{location}/publishers/google/models/{model}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#model VertexAiReasoningEngine#model}
	Model *string `field:"required" json:"model" yaml:"model"`
	// generation_trigger_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#generation_trigger_config VertexAiReasoningEngine#generation_trigger_config}
	GenerationTriggerConfig *VertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfigGenerationTriggerConfig `field:"optional" json:"generationTriggerConfig" yaml:"generationTriggerConfig"`
}

