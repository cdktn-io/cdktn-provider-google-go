// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaionlineevaluator


type VertexAiOnlineEvaluatorCloudObservabilitySessionScope struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/vertex_ai_online_evaluator#filter VertexAiOnlineEvaluator#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// inactivity_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.7.0/docs/resources/vertex_ai_online_evaluator#inactivity_trigger VertexAiOnlineEvaluator#inactivity_trigger}
	InactivityTrigger *VertexAiOnlineEvaluatorCloudObservabilitySessionScopeInactivityTrigger `field:"optional" json:"inactivityTrigger" yaml:"inactivityTrigger"`
}

