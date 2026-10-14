// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package geminidataanalyticsdataagent


type GeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOptions struct {
	// analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/gemini_data_analytics_data_agent#analysis GeminiDataAnalyticsDataAgent#analysis}
	Analysis *GeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOptionsAnalysis `field:"optional" json:"analysis" yaml:"analysis"`
	// datasource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/gemini_data_analytics_data_agent#datasource GeminiDataAnalyticsDataAgent#datasource}
	Datasource *GeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOptionsDatasource `field:"optional" json:"datasource" yaml:"datasource"`
}

