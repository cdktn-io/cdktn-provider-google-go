// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package geminidataanalyticsdataagent


type GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextUserFunctionsBqRoutines struct {
	// Optional. User override or addition to description, to tell the agent when to use the UDF.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#description GeminiDataAnalyticsDataAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// routine_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/gemini_data_analytics_data_agent#routine_reference GeminiDataAnalyticsDataAgent#routine_reference}
	RoutineReference *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextUserFunctionsBqRoutinesRoutineReference `field:"optional" json:"routineReference" yaml:"routineReference"`
}

