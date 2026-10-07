// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexairagcorpus


type VertexAiRagCorpusVectorDbConfigRagManagedDb struct {
	// ann block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_rag_corpus#ann VertexAiRagCorpus#ann}
	Ann *VertexAiRagCorpusVectorDbConfigRagManagedDbAnn `field:"optional" json:"ann" yaml:"ann"`
	// knn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_rag_corpus#knn VertexAiRagCorpus#knn}
	Knn *VertexAiRagCorpusVectorDbConfigRagManagedDbKnn `field:"optional" json:"knn" yaml:"knn"`
}

