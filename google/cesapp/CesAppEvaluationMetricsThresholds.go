// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesapp


type CesAppEvaluationMetricsThresholds struct {
	// golden_evaluation_metrics_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#golden_evaluation_metrics_thresholds CesApp#golden_evaluation_metrics_thresholds}
	GoldenEvaluationMetricsThresholds *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholds `field:"optional" json:"goldenEvaluationMetricsThresholds" yaml:"goldenEvaluationMetricsThresholds"`
	// The hallucination metric behavior for golden evaluations. Possible values: ["DISABLED", "ENABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#golden_hallucination_metric_behavior CesApp#golden_hallucination_metric_behavior}
	GoldenHallucinationMetricBehavior *string `field:"optional" json:"goldenHallucinationMetricBehavior" yaml:"goldenHallucinationMetricBehavior"`
	// The hallucination metric behavior for scenario evaluations. Possible values: ["DISABLED", "ENABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#scenario_hallucination_metric_behavior CesApp#scenario_hallucination_metric_behavior}
	ScenarioHallucinationMetricBehavior *string `field:"optional" json:"scenarioHallucinationMetricBehavior" yaml:"scenarioHallucinationMetricBehavior"`
}

