// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package geminidataanalyticsdataagent


type GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueries struct {
	// looker_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#looker_query GeminiDataAnalyticsDataAgent#looker_query}
	LookerQuery *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesLookerQuery `field:"optional" json:"lookerQuery" yaml:"lookerQuery"`
	// Optional. Natural language questions that a user might ask.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#natural_language_questions GeminiDataAnalyticsDataAgent#natural_language_questions}
	NaturalLanguageQuestions *[]*string `field:"optional" json:"naturalLanguageQuestions" yaml:"naturalLanguageQuestions"`
}

