// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexairagcorpus


type VertexAiRagCorpusVectorDbConfigVertexVectorSearch struct {
	// The resource name of the Index. Format: projects/{project}/locations/{location}/indexes/{index}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_rag_corpus#index VertexAiRagCorpus#index}
	Index *string `field:"required" json:"index" yaml:"index"`
	// The resource name of the Index Endpoint. Format: projects/{project}/locations/{location}/indexEndpoints/{index_endpoint}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_rag_corpus#index_endpoint VertexAiRagCorpus#index_endpoint}
	IndexEndpoint *string `field:"required" json:"indexEndpoint" yaml:"indexEndpoint"`
}

