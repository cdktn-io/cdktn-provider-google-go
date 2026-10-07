// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agenticapplicationsanalystagentpersona


type AgenticApplicationsAnalystAgentPersonaArtifactsConfigMethodologyExportOptions struct {
	// If true, append the detailed methodology to the final response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/agentic_applications_analyst_agent_persona#append_methodology AgenticApplicationsAnalystAgentPersona#append_methodology}
	AppendMethodology interface{} `field:"optional" json:"appendMethodology" yaml:"appendMethodology"`
	// Format for methodology export. Possible values: MARKDOWN HTML PDF.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/agentic_applications_analyst_agent_persona#export_format AgenticApplicationsAnalystAgentPersona#export_format}
	ExportFormat *string `field:"optional" json:"exportFormat" yaml:"exportFormat"`
	// If true, export the detailed methodology as a separate artifact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/agentic_applications_analyst_agent_persona#export_methodology_artifact AgenticApplicationsAnalystAgentPersona#export_methodology_artifact}
	ExportMethodologyArtifact interface{} `field:"optional" json:"exportMethodologyArtifact" yaml:"exportMethodologyArtifact"`
}

