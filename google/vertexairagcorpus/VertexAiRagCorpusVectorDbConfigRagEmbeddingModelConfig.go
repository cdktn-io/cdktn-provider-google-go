// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexairagcorpus


type VertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfig struct {
	// vertex_prediction_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_rag_corpus#vertex_prediction_endpoint VertexAiRagCorpus#vertex_prediction_endpoint}
	VertexPredictionEndpoint *VertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfigVertexPredictionEndpoint `field:"optional" json:"vertexPredictionEndpoint" yaml:"vertexPredictionEndpoint"`
}

