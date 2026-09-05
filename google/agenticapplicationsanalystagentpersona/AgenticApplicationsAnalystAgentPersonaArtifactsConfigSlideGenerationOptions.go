// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agenticapplicationsanalystagentpersona


type AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptions struct {
	// Format for slide export. Possible values: PDF PNG PPTX GOOGLE_SLIDES.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona#export_format AgenticApplicationsAnalystAgentPersona#export_format}
	ExportFormat *string `field:"optional" json:"exportFormat" yaml:"exportFormat"`
	// slide_examples block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona#slide_examples AgenticApplicationsAnalystAgentPersona#slide_examples}
	SlideExamples interface{} `field:"optional" json:"slideExamples" yaml:"slideExamples"`
}

