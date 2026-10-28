// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agenticapplicationsanalystagentpersona


type AgenticApplicationsAnalystAgentPersonaArtifactsConfig struct {
	// audit_trail_export_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/agentic_applications_analyst_agent_persona#audit_trail_export_options AgenticApplicationsAnalystAgentPersona#audit_trail_export_options}
	AuditTrailExportOptions *AgenticApplicationsAnalystAgentPersonaArtifactsConfigAuditTrailExportOptions `field:"optional" json:"auditTrailExportOptions" yaml:"auditTrailExportOptions"`
	// document_generation_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/agentic_applications_analyst_agent_persona#document_generation_options AgenticApplicationsAnalystAgentPersona#document_generation_options}
	DocumentGenerationOptions *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptions `field:"optional" json:"documentGenerationOptions" yaml:"documentGenerationOptions"`
	// methodology_export_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/agentic_applications_analyst_agent_persona#methodology_export_options AgenticApplicationsAnalystAgentPersona#methodology_export_options}
	MethodologyExportOptions *AgenticApplicationsAnalystAgentPersonaArtifactsConfigMethodologyExportOptions `field:"optional" json:"methodologyExportOptions" yaml:"methodologyExportOptions"`
	// slide_generation_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/agentic_applications_analyst_agent_persona#slide_generation_options AgenticApplicationsAnalystAgentPersona#slide_generation_options}
	SlideGenerationOptions *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptions `field:"optional" json:"slideGenerationOptions" yaml:"slideGenerationOptions"`
	// visualization_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/agentic_applications_analyst_agent_persona#visualization_options AgenticApplicationsAnalystAgentPersona#visualization_options}
	VisualizationOptions *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptions `field:"optional" json:"visualizationOptions" yaml:"visualizationOptions"`
}

