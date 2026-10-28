// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package geminidataanalyticsdataagent


type GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextDatasourceReferences struct {
	// bq block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#bq GeminiDataAnalyticsDataAgent#bq}
	Bq *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextDatasourceReferencesBq `field:"optional" json:"bq" yaml:"bq"`
	// looker block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#looker GeminiDataAnalyticsDataAgent#looker}
	Looker *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextDatasourceReferencesLooker `field:"optional" json:"looker" yaml:"looker"`
	// studio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#studio GeminiDataAnalyticsDataAgent#studio}
	Studio *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextDatasourceReferencesStudio `field:"optional" json:"studio" yaml:"studio"`
}

