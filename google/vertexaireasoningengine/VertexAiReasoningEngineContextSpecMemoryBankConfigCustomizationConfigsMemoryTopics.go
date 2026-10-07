// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopics struct {
	// custom_memory_topic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#custom_memory_topic VertexAiReasoningEngine#custom_memory_topic}
	CustomMemoryTopic *VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsCustomMemoryTopic `field:"optional" json:"customMemoryTopic" yaml:"customMemoryTopic"`
	// managed_memory_topic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#managed_memory_topic VertexAiReasoningEngine#managed_memory_topic}
	ManagedMemoryTopic *VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopic `field:"optional" json:"managedMemoryTopic" yaml:"managedMemoryTopic"`
}

