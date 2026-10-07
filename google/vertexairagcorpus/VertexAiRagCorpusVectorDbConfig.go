// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexairagcorpus


type VertexAiRagCorpusVectorDbConfig struct {
	// api_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_rag_corpus#api_auth VertexAiRagCorpus#api_auth}
	ApiAuth *VertexAiRagCorpusVectorDbConfigApiAuth `field:"optional" json:"apiAuth" yaml:"apiAuth"`
	// pinecone block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_rag_corpus#pinecone VertexAiRagCorpus#pinecone}
	Pinecone *VertexAiRagCorpusVectorDbConfigPinecone `field:"optional" json:"pinecone" yaml:"pinecone"`
	// rag_embedding_model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_rag_corpus#rag_embedding_model_config VertexAiRagCorpus#rag_embedding_model_config}
	RagEmbeddingModelConfig *VertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfig `field:"optional" json:"ragEmbeddingModelConfig" yaml:"ragEmbeddingModelConfig"`
	// rag_managed_db block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_rag_corpus#rag_managed_db VertexAiRagCorpus#rag_managed_db}
	RagManagedDb *VertexAiRagCorpusVectorDbConfigRagManagedDb `field:"optional" json:"ragManagedDb" yaml:"ragManagedDb"`
	// vertex_vector_search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_rag_corpus#vertex_vector_search VertexAiRagCorpus#vertex_vector_search}
	VertexVectorSearch *VertexAiRagCorpusVectorDbConfigVertexVectorSearch `field:"optional" json:"vertexVectorSearch" yaml:"vertexVectorSearch"`
}

