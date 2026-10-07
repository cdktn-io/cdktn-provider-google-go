// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineContextSpecMemoryBankConfig struct {
	// customization_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#customization_configs VertexAiReasoningEngine#customization_configs}
	CustomizationConfigs interface{} `field:"optional" json:"customizationConfigs" yaml:"customizationConfigs"`
	// If true, no memory revisions will be created for any requests to the Memory Bank.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#disable_memory_revisions VertexAiReasoningEngine#disable_memory_revisions}
	DisableMemoryRevisions interface{} `field:"optional" json:"disableMemoryRevisions" yaml:"disableMemoryRevisions"`
	// generation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#generation_config VertexAiReasoningEngine#generation_config}
	GenerationConfig *VertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfig `field:"optional" json:"generationConfig" yaml:"generationConfig"`
	// similarity_search_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#similarity_search_config VertexAiReasoningEngine#similarity_search_config}
	SimilaritySearchConfig *VertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfig `field:"optional" json:"similaritySearchConfig" yaml:"similaritySearchConfig"`
	// structured_memory_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#structured_memory_configs VertexAiReasoningEngine#structured_memory_configs}
	StructuredMemoryConfigs interface{} `field:"optional" json:"structuredMemoryConfigs" yaml:"structuredMemoryConfigs"`
	// ttl_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#ttl_config VertexAiReasoningEngine#ttl_config}
	TtlConfig *VertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfig `field:"optional" json:"ttlConfig" yaml:"ttlConfig"`
}

