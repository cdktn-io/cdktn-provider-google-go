// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContent struct {
	// parts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#parts VertexAiReasoningEngine#parts}
	Parts interface{} `field:"required" json:"parts" yaml:"parts"`
	// The producer of the content.
	//
	// Must be either 'user' or 'model'. If not set, the service will default to 'user'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#role VertexAiReasoningEngine#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

