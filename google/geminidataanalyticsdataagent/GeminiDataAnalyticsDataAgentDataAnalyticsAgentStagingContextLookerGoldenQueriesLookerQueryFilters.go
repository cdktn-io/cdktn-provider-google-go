// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package geminidataanalyticsdataagent


type GeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextLookerGoldenQueriesLookerQueryFilters struct {
	// Required. The field to filter on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#field GeminiDataAnalyticsDataAgent#field}
	Field *string `field:"required" json:"field" yaml:"field"`
	// Optional. The value for the field to filter on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#value GeminiDataAnalyticsDataAgent#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

