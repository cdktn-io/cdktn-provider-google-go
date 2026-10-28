// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package geminidataanalyticsdataagent


type GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextDatasourceReferencesLookerExploreReferencesPrivateLookerInstanceInfo struct {
	// Optional. The Looker instance ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#looker_instance_id GeminiDataAnalyticsDataAgent#looker_instance_id}
	LookerInstanceId *string `field:"optional" json:"lookerInstanceId" yaml:"lookerInstanceId"`
	// Optional. The service directory name of the Looker instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#service_directory_name GeminiDataAnalyticsDataAgent#service_directory_name}
	ServiceDirectoryName *string `field:"optional" json:"serviceDirectoryName" yaml:"serviceDirectoryName"`
}

