// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexairagfile


type VertexAiRagFileRagFileTransformationConfigRagFileChunkingConfigFixedLengthChunking struct {
	// The overlap between chunks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_rag_file#chunk_overlap VertexAiRagFile#chunk_overlap}
	ChunkOverlap *float64 `field:"optional" json:"chunkOverlap" yaml:"chunkOverlap"`
	// The size of the chunks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_rag_file#chunk_size VertexAiRagFile#chunk_size}
	ChunkSize *float64 `field:"optional" json:"chunkSize" yaml:"chunkSize"`
}

