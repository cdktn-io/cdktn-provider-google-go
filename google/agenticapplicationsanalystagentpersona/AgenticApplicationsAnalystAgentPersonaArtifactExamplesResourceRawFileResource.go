// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agenticapplicationsanalystagentpersona


type AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceRawFileResource struct {
	// The raw file content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/agentic_applications_analyst_agent_persona#file_content AgenticApplicationsAnalystAgentPersona#file_content}
	FileContent *string `field:"required" json:"fileContent" yaml:"fileContent"`
	// The title of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/agentic_applications_analyst_agent_persona#file_title AgenticApplicationsAnalystAgentPersona#file_title}
	FileTitle *string `field:"required" json:"fileTitle" yaml:"fileTitle"`
	// The mime type of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/agentic_applications_analyst_agent_persona#mime_type AgenticApplicationsAnalystAgentPersona#mime_type}
	MimeType *string `field:"required" json:"mimeType" yaml:"mimeType"`
}

