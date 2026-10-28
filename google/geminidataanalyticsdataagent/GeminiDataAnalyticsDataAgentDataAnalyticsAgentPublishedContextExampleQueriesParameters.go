// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package geminidataanalyticsdataagent


type GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextExampleQueriesParameters struct {
	// Required. The data type of the parameter, e.g. "STRING", "INT64", "DATE".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#data_type GeminiDataAnalyticsDataAgent#data_type}
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// Required. The name of the parameter reference in the SQL query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#name GeminiDataAnalyticsDataAgent#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional.
	//
	// The description of the parameter that can be used by the LLM to extract the parameter value from the user question.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#description GeminiDataAnalyticsDataAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

