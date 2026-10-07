// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowgenerator


type DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfo struct {
	// tool_call block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dialogflow_generator#tool_call DialogflowGenerator#tool_call}
	ToolCall *DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCall `field:"required" json:"toolCall" yaml:"toolCall"`
	// tool_call_result block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dialogflow_generator#tool_call_result DialogflowGenerator#tool_call_result}
	ToolCallResult *DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResult `field:"required" json:"toolCallResult" yaml:"toolCallResult"`
}

