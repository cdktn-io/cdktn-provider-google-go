// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineContextSpec struct {
	// memory_bank_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#memory_bank_config VertexAiReasoningEngine#memory_bank_config}
	MemoryBankConfig *VertexAiReasoningEngineContextSpecMemoryBankConfig `field:"optional" json:"memoryBankConfig" yaml:"memoryBankConfig"`
}

