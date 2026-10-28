// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexairagfile


type VertexAiRagFileRagFileTransformationConfig struct {
	// rag_file_chunking_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_rag_file#rag_file_chunking_config VertexAiRagFile#rag_file_chunking_config}
	RagFileChunkingConfig *VertexAiRagFileRagFileTransformationConfigRagFileChunkingConfig `field:"optional" json:"ragFileChunkingConfig" yaml:"ragFileChunkingConfig"`
}

