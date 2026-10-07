// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsFileData struct {
	// The URI of the file in Google Cloud Storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#file_uri VertexAiReasoningEngine#file_uri}
	FileUri *string `field:"required" json:"fileUri" yaml:"fileUri"`
	// The IANA standard MIME type of the source data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#mime_type VertexAiReasoningEngine#mime_type}
	MimeType *string `field:"required" json:"mimeType" yaml:"mimeType"`
}

