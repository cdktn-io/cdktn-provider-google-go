// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsAudioTranscriptionWords struct {
	// Transcript of the word.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#word VertexAiReasoningEngine#word}
	Word *string `field:"required" json:"word" yaml:"word"`
	// End offset in time of the word relative to the start of the audio.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#end_offset VertexAiReasoningEngine#end_offset}
	EndOffset *string `field:"optional" json:"endOffset" yaml:"endOffset"`
	// Start offset in time of the word relative to the start of the audio.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#start_offset VertexAiReasoningEngine#start_offset}
	StartOffset *string `field:"optional" json:"startOffset" yaml:"startOffset"`
}

