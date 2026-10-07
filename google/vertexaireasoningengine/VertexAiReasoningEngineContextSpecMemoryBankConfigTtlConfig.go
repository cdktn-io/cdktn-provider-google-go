// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfig struct {
	// The default TTL duration of the memories in the Memory Bank.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#default_ttl VertexAiReasoningEngine#default_ttl}
	DefaultTtl *string `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// granular_ttl_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#granular_ttl_config VertexAiReasoningEngine#granular_ttl_config}
	GranularTtlConfig *VertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfigGranularTtlConfig `field:"optional" json:"granularTtlConfig" yaml:"granularTtlConfig"`
	// The default TTL duration of the memory revisions in the Memory Bank.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#memory_revision_default_ttl VertexAiReasoningEngine#memory_revision_default_ttl}
	MemoryRevisionDefaultTtl *string `field:"optional" json:"memoryRevisionDefaultTtl" yaml:"memoryRevisionDefaultTtl"`
}

