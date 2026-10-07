// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexairagcorpus


type VertexAiRagCorpusVectorDbConfigPinecone struct {
	// Pinecone index name. This value cannot be changed after it's set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_rag_corpus#index_name VertexAiRagCorpus#index_name}
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
}

