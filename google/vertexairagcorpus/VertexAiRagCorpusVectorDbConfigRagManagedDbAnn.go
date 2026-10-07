// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexairagcorpus


type VertexAiRagCorpusVectorDbConfigRagManagedDbAnn struct {
	// Number of leaf nodes in the tree-based structure. Default value is 500.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_rag_corpus#leaf_count VertexAiRagCorpus#leaf_count}
	LeafCount *float64 `field:"optional" json:"leafCount" yaml:"leafCount"`
	// The depth of the tree-based structure. Only depth values of 2 and 3 are supported. Default value is 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_rag_corpus#tree_depth VertexAiRagCorpus#tree_depth}
	TreeDepth *float64 `field:"optional" json:"treeDepth" yaml:"treeDepth"`
}

