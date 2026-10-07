// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamples struct {
	// conversation_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#conversation_source VertexAiReasoningEngine#conversation_source}
	ConversationSource *VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSource `field:"optional" json:"conversationSource" yaml:"conversationSource"`
	// generated_memories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#generated_memories VertexAiReasoningEngine#generated_memories}
	GeneratedMemories interface{} `field:"optional" json:"generatedMemories" yaml:"generatedMemories"`
}

