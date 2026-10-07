// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agenticapplicationsanalystagentpersona


type AgenticApplicationsAnalystAgentPersonaWebSearchConfig struct {
	// Whether web search grounding is disabled for the analyst agent.
	//
	// Defaults to false if not specified (i.e. web search grounding is enabled).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/agentic_applications_analyst_agent_persona#disabled AgenticApplicationsAnalystAgentPersona#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// List of domains to be excluded from Google Search / Enterprise Web Search grounding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/agentic_applications_analyst_agent_persona#excluded_domains AgenticApplicationsAnalystAgentPersona#excluded_domains}
	ExcludedDomains *[]*string `field:"optional" json:"excludedDomains" yaml:"excludedDomains"`
}

