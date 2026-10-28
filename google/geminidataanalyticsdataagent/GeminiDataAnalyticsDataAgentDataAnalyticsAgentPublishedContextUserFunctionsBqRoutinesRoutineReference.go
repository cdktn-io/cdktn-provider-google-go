// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package geminidataanalyticsdataagent


type GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextUserFunctionsBqRoutinesRoutineReference struct {
	// Optional. The dataset ID of the routine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#dataset_id GeminiDataAnalyticsDataAgent#dataset_id}
	DatasetId *string `field:"optional" json:"datasetId" yaml:"datasetId"`
	// Optional. The project ID of the routine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#project_id GeminiDataAnalyticsDataAgent#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Optional. The routine ID of the routine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#routine_id GeminiDataAnalyticsDataAgent#routine_id}
	RoutineId *string `field:"optional" json:"routineId" yaml:"routineId"`
}

