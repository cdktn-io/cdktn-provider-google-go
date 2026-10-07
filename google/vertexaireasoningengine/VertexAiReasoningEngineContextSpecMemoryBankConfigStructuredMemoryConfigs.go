// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineContextSpecMemoryBankConfigStructuredMemoryConfigs struct {
	// schema_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#schema_configs VertexAiReasoningEngine#schema_configs}
	SchemaConfigs interface{} `field:"optional" json:"schemaConfigs" yaml:"schemaConfigs"`
	// Optional. List of scope keys that this structured memory config applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#scope_keys VertexAiReasoningEngine#scope_keys}
	ScopeKeys *[]*string `field:"optional" json:"scopeKeys" yaml:"scopeKeys"`
}

