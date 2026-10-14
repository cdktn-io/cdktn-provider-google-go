// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package geminidataanalyticsdataagent


type GeminiDataAnalyticsDataAgentDataAnalyticsAgent struct {
	// published_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/gemini_data_analytics_data_agent#published_context GeminiDataAnalyticsDataAgent#published_context}
	PublishedContext *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContext `field:"optional" json:"publishedContext" yaml:"publishedContext"`
	// staging_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/gemini_data_analytics_data_agent#staging_context GeminiDataAnalyticsDataAgent#staging_context}
	StagingContext *GeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContext `field:"optional" json:"stagingContext" yaml:"stagingContext"`
}

