// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesGeneratedMemories struct {
	// Represents the fact to generate a memory from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#fact VertexAiReasoningEngine#fact}
	Fact *string `field:"required" json:"fact" yaml:"fact"`
	// topics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#topics VertexAiReasoningEngine#topics}
	Topics interface{} `field:"optional" json:"topics" yaml:"topics"`
}

