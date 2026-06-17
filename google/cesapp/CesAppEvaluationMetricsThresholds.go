// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesapp


type CesAppEvaluationMetricsThresholds struct {
	// golden_evaluation_metrics_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/ces_app#golden_evaluation_metrics_thresholds CesApp#golden_evaluation_metrics_thresholds}
	GoldenEvaluationMetricsThresholds *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholds `field:"optional" json:"goldenEvaluationMetricsThresholds" yaml:"goldenEvaluationMetricsThresholds"`
}

