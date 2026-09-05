// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agenticapplicationsanalystagentpersona


type AgenticApplicationsAnalystAgentPersonaTables struct {
	// The name of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona#name AgenticApplicationsAnalystAgentPersona#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona#columns AgenticApplicationsAnalystAgentPersona#columns}
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// The description of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona#description AgenticApplicationsAnalystAgentPersona#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

