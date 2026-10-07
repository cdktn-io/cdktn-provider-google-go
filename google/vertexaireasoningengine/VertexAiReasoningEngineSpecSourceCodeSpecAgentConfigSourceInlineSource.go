// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceInlineSource struct {
	// Required. Input only. The application source code archive, provided as a compressed tarball (.tar.gz) file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#source_archive VertexAiReasoningEngine#source_archive}
	SourceArchive *string `field:"required" json:"sourceArchive" yaml:"sourceArchive"`
}

