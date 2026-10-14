// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaionlineevaluator


type VertexAiOnlineEvaluatorCloudObservabilityOpenTelemetry struct {
	// Defines which version of the OTel Semantic Convention the data follows. Can be "1.39.0" or newer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/vertex_ai_online_evaluator#semconv_version VertexAiOnlineEvaluator#semconv_version}
	SemconvVersion *string `field:"required" json:"semconvVersion" yaml:"semconvVersion"`
}

