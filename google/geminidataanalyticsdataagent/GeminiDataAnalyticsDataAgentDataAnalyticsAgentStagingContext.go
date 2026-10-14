// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package geminidataanalyticsdataagent


type GeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContext struct {
	// datasource_references block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/gemini_data_analytics_data_agent#datasource_references GeminiDataAnalyticsDataAgent#datasource_references}
	DatasourceReferences *GeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextDatasourceReferences `field:"required" json:"datasourceReferences" yaml:"datasourceReferences"`
	// example_queries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/gemini_data_analytics_data_agent#example_queries GeminiDataAnalyticsDataAgent#example_queries}
	ExampleQueries interface{} `field:"optional" json:"exampleQueries" yaml:"exampleQueries"`
	// glossary_terms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/gemini_data_analytics_data_agent#glossary_terms GeminiDataAnalyticsDataAgent#glossary_terms}
	GlossaryTerms interface{} `field:"optional" json:"glossaryTerms" yaml:"glossaryTerms"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/gemini_data_analytics_data_agent#options GeminiDataAnalyticsDataAgent#options}
	Options *GeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOptions `field:"optional" json:"options" yaml:"options"`
	// schema_relationships block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/gemini_data_analytics_data_agent#schema_relationships GeminiDataAnalyticsDataAgent#schema_relationships}
	SchemaRelationships interface{} `field:"optional" json:"schemaRelationships" yaml:"schemaRelationships"`
	// Optional. The basic entry point for data owners creating domain knowledge for the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/gemini_data_analytics_data_agent#system_instruction GeminiDataAnalyticsDataAgent#system_instruction}
	SystemInstruction *string `field:"optional" json:"systemInstruction" yaml:"systemInstruction"`
}

