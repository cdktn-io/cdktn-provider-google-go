// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package geminigibqobservabilitysetting


type GeminiGibqObservabilitySettingConversationalAnalyticsSetting struct {
	// Whether to enable feedback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/gemini_gibq_observability_setting#feedback_enabled GeminiGibqObservabilitySetting#feedback_enabled}
	FeedbackEnabled interface{} `field:"optional" json:"feedbackEnabled" yaml:"feedbackEnabled"`
	// Whether to enable logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/gemini_gibq_observability_setting#logging_enabled GeminiGibqObservabilitySetting#logging_enabled}
	LoggingEnabled interface{} `field:"optional" json:"loggingEnabled" yaml:"loggingEnabled"`
	// Whether to enable metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/gemini_gibq_observability_setting#metrics_enabled GeminiGibqObservabilitySetting#metrics_enabled}
	MetricsEnabled interface{} `field:"optional" json:"metricsEnabled" yaml:"metricsEnabled"`
	// Whether to enable traces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/gemini_gibq_observability_setting#traces_enabled GeminiGibqObservabilitySetting#traces_enabled}
	TracesEnabled interface{} `field:"optional" json:"tracesEnabled" yaml:"tracesEnabled"`
}

