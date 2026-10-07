// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfig struct {
	// The model used to generate embeddings to lookup similar memories. Format: projects/{project}/locations/{location}/publishers/google/models/{model}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#embedding_model VertexAiReasoningEngine#embedding_model}
	EmbeddingModel *string `field:"required" json:"embeddingModel" yaml:"embeddingModel"`
}

