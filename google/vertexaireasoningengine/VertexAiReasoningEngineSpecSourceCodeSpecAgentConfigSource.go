// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSource struct {
	// adk_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#adk_config VertexAiReasoningEngine#adk_config}
	AdkConfig *VertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceAdkConfig `field:"optional" json:"adkConfig" yaml:"adkConfig"`
	// inline_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#inline_source VertexAiReasoningEngine#inline_source}
	InlineSource *VertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceInlineSource `field:"optional" json:"inlineSource" yaml:"inlineSource"`
}

