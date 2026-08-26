// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agenticapplicationsanalystagentpersona


type AgenticApplicationsAnalystAgentPersonaExternalDataSources struct {
	// Whether this external data source is enabled for the current analysis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/agentic_applications_analyst_agent_persona#enabled AgenticApplicationsAnalystAgentPersona#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// air_quality block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/agentic_applications_analyst_agent_persona#air_quality AgenticApplicationsAnalystAgentPersona#air_quality}
	AirQuality *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQuality `field:"optional" json:"airQuality" yaml:"airQuality"`
	// bureau_labor_statistics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/agentic_applications_analyst_agent_persona#bureau_labor_statistics AgenticApplicationsAnalystAgentPersona#bureau_labor_statistics}
	BureauLaborStatistics *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatistics `field:"optional" json:"bureauLaborStatistics" yaml:"bureauLaborStatistics"`
	// coindesk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/agentic_applications_analyst_agent_persona#coindesk AgenticApplicationsAnalystAgentPersona#coindesk}
	Coindesk *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindesk `field:"optional" json:"coindesk" yaml:"coindesk"`
	// finnhub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/agentic_applications_analyst_agent_persona#finnhub AgenticApplicationsAnalystAgentPersona#finnhub}
	Finnhub *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhub `field:"optional" json:"finnhub" yaml:"finnhub"`
	// fred block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/agentic_applications_analyst_agent_persona#fred AgenticApplicationsAnalystAgentPersona#fred}
	Fred *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFred `field:"optional" json:"fred" yaml:"fred"`
	// sec_edgar block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/agentic_applications_analyst_agent_persona#sec_edgar AgenticApplicationsAnalystAgentPersona#sec_edgar}
	SecEdgar *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgar `field:"optional" json:"secEdgar" yaml:"secEdgar"`
	// treasury_securities_auctions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/agentic_applications_analyst_agent_persona#treasury_securities_auctions AgenticApplicationsAnalystAgentPersona#treasury_securities_auctions}
	TreasurySecuritiesAuctions *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctions `field:"optional" json:"treasurySecuritiesAuctions" yaml:"treasurySecuritiesAuctions"`
	// usda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/agentic_applications_analyst_agent_persona#usda AgenticApplicationsAnalystAgentPersona#usda}
	Usda *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsda `field:"optional" json:"usda" yaml:"usda"`
}

