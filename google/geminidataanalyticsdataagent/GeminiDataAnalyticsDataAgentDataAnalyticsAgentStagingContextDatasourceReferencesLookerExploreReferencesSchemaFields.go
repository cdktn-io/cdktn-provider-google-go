// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package geminidataanalyticsdataagent


type GeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextDatasourceReferencesLookerExploreReferencesSchemaFields struct {
	// Optional. Field category.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#category GeminiDataAnalyticsDataAgent#category}
	Category *string `field:"optional" json:"category" yaml:"category"`
	// Optional. A brief description of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#description GeminiDataAnalyticsDataAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional. The mode of the field (e.g., NULLABLE, REPEATED).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#mode GeminiDataAnalyticsDataAgent#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Optional. The name of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#name GeminiDataAnalyticsDataAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional. Recursive property for nested schema structures, encoded as a JSON string representing a list of Field objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#subfields GeminiDataAnalyticsDataAgent#subfields}
	Subfields *string `field:"optional" json:"subfields" yaml:"subfields"`
	// Optional. A list of alternative names or synonyms that can be used to refer to this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#synonyms GeminiDataAnalyticsDataAgent#synonyms}
	Synonyms *[]*string `field:"optional" json:"synonyms" yaml:"synonyms"`
	// Optional. A list of tags or keywords associated with the field, used for categorization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#tags GeminiDataAnalyticsDataAgent#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Optional. The type of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#type GeminiDataAnalyticsDataAgent#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Optional. Value format of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#value_format GeminiDataAnalyticsDataAgent#value_format}
	ValueFormat *string `field:"optional" json:"valueFormat" yaml:"valueFormat"`
}

