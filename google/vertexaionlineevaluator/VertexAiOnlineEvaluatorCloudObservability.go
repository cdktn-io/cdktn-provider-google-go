// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaionlineevaluator


type VertexAiOnlineEvaluatorCloudObservability struct {
	// open_telemetry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_online_evaluator#open_telemetry VertexAiOnlineEvaluator#open_telemetry}
	OpenTelemetry *VertexAiOnlineEvaluatorCloudObservabilityOpenTelemetry `field:"required" json:"openTelemetry" yaml:"openTelemetry"`
	// Optional log view that will be used to query logs. If empty, the project's default view ('projects/{{project_id}}') will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_online_evaluator#log_view VertexAiOnlineEvaluator#log_view}
	LogView *string `field:"optional" json:"logView" yaml:"logView"`
	// session_scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_online_evaluator#session_scope VertexAiOnlineEvaluator#session_scope}
	SessionScope *VertexAiOnlineEvaluatorCloudObservabilitySessionScope `field:"optional" json:"sessionScope" yaml:"sessionScope"`
	// trace_scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_online_evaluator#trace_scope VertexAiOnlineEvaluator#trace_scope}
	TraceScope *VertexAiOnlineEvaluatorCloudObservabilityTraceScope `field:"optional" json:"traceScope" yaml:"traceScope"`
	// Optional trace view that will be used to query traces.
	//
	// If empty, the
	// '_AllSpans' view from the '_Trace' US bucket will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_online_evaluator#trace_view VertexAiOnlineEvaluator#trace_view}
	TraceView *string `field:"optional" json:"traceView" yaml:"traceView"`
}

