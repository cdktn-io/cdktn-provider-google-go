// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package geminidataanalyticsdataagent


type GeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextDatasourceReferencesLookerExploreReferences struct {
	// Required. Name of the LookML Explore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#explore GeminiDataAnalyticsDataAgent#explore}
	Explore *string `field:"required" json:"explore" yaml:"explore"`
	// Required. Name of the LookML model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#lookml_model GeminiDataAnalyticsDataAgent#lookml_model}
	LookmlModel *string `field:"required" json:"lookmlModel" yaml:"lookmlModel"`
	// Optional. The base URL of the Looker instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#looker_instance_uri GeminiDataAnalyticsDataAgent#looker_instance_uri}
	LookerInstanceUri *string `field:"optional" json:"lookerInstanceUri" yaml:"lookerInstanceUri"`
	// private_looker_instance_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#private_looker_instance_info GeminiDataAnalyticsDataAgent#private_looker_instance_info}
	PrivateLookerInstanceInfo *GeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextDatasourceReferencesLookerExploreReferencesPrivateLookerInstanceInfo `field:"optional" json:"privateLookerInstanceInfo" yaml:"privateLookerInstanceInfo"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#schema GeminiDataAnalyticsDataAgent#schema}
	Schema *GeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextDatasourceReferencesLookerExploreReferencesSchema `field:"optional" json:"schema" yaml:"schema"`
}

