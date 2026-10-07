// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesapp


type CesAppLoggingSettingsMetricAnalysisSettings struct {
	// Whether to collect conversation data for llm analysis metrics.
	//
	// If true,
	// conversation data will not be collected for llm analysis metrics;
	// otherwise, conversation data will be collected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#llm_metrics_opted_out CesApp#llm_metrics_opted_out}
	LlmMetricsOptedOut interface{} `field:"optional" json:"llmMetricsOptedOut" yaml:"llmMetricsOptedOut"`
}

