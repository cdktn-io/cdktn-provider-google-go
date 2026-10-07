// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesapp


type CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholds struct {
	// expectation_level_metrics_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#expectation_level_metrics_thresholds CesApp#expectation_level_metrics_thresholds}
	ExpectationLevelMetricsThresholds *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholds `field:"optional" json:"expectationLevelMetricsThresholds" yaml:"expectationLevelMetricsThresholds"`
	// tool_matching_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#tool_matching_settings CesApp#tool_matching_settings}
	ToolMatchingSettings *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettings `field:"optional" json:"toolMatchingSettings" yaml:"toolMatchingSettings"`
	// turn_level_metrics_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#turn_level_metrics_thresholds CesApp#turn_level_metrics_thresholds}
	TurnLevelMetricsThresholds *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholds `field:"optional" json:"turnLevelMetricsThresholds" yaml:"turnLevelMetricsThresholds"`
}

