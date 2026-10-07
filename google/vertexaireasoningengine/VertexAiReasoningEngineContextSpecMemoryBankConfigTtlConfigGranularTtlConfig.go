// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfigGranularTtlConfig struct {
	// The TTL duration for memories uploaded via CreateMemory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#create_ttl VertexAiReasoningEngine#create_ttl}
	CreateTtl *string `field:"optional" json:"createTtl" yaml:"createTtl"`
	// The TTL duration for memories newly generated via GenerateMemories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#generate_created_ttl VertexAiReasoningEngine#generate_created_ttl}
	GenerateCreatedTtl *string `field:"optional" json:"generateCreatedTtl" yaml:"generateCreatedTtl"`
	// The TTL duration for memories updated via GenerateMemories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#generate_updated_ttl VertexAiReasoningEngine#generate_updated_ttl}
	GenerateUpdatedTtl *string `field:"optional" json:"generateUpdatedTtl" yaml:"generateUpdatedTtl"`
}

