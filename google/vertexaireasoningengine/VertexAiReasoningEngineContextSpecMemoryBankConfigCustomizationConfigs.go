// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigs struct {
	// consolidation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#consolidation_config VertexAiReasoningEngine#consolidation_config}
	ConsolidationConfig *VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsConsolidationConfig `field:"optional" json:"consolidationConfig" yaml:"consolidationConfig"`
	// Indicates whether natural language memory generation should be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#disable_natural_language_memories VertexAiReasoningEngine#disable_natural_language_memories}
	DisableNaturalLanguageMemories interface{} `field:"optional" json:"disableNaturalLanguageMemories" yaml:"disableNaturalLanguageMemories"`
	// Optional. Generate memories in the third person if set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#enable_third_person_memories VertexAiReasoningEngine#enable_third_person_memories}
	EnableThirdPersonMemories interface{} `field:"optional" json:"enableThirdPersonMemories" yaml:"enableThirdPersonMemories"`
	// generate_memories_examples block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#generate_memories_examples VertexAiReasoningEngine#generate_memories_examples}
	GenerateMemoriesExamples interface{} `field:"optional" json:"generateMemoriesExamples" yaml:"generateMemoriesExamples"`
	// memory_topics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#memory_topics VertexAiReasoningEngine#memory_topics}
	MemoryTopics interface{} `field:"optional" json:"memoryTopics" yaml:"memoryTopics"`
	// Optional. List of scope keys that this customization config applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#scope_keys VertexAiReasoningEngine#scope_keys}
	ScopeKeys *[]*string `field:"optional" json:"scopeKeys" yaml:"scopeKeys"`
}

