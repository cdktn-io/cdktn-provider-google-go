// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsSpeechMetadata struct {
	// Identifies which speaker is speaking this turn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_reasoning_engine#speaker VertexAiReasoningEngine#speaker}
	Speaker *string `field:"optional" json:"speaker" yaml:"speaker"`
	// Natural language description of the vocal style (e.g., "cheerful").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_reasoning_engine#style VertexAiReasoningEngine#style}
	Style *string `field:"optional" json:"style" yaml:"style"`
}

