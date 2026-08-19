// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vectorsearchcollection


type VectorSearchCollectionVectorSchemaDenseVector struct {
	// Dimensionality of the vector field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vector_search_collection#dimensions VectorSearchCollection#dimensions}
	Dimensions *float64 `field:"optional" json:"dimensions" yaml:"dimensions"`
	// vertex_embedding_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vector_search_collection#vertex_embedding_config VectorSearchCollection#vertex_embedding_config}
	VertexEmbeddingConfig *VectorSearchCollectionVectorSchemaDenseVectorVertexEmbeddingConfig `field:"optional" json:"vertexEmbeddingConfig" yaml:"vertexEmbeddingConfig"`
}

