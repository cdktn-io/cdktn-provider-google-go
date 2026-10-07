// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowgenerator


type DialogflowGeneratorSummarizationContextFewShotExamplesOutput struct {
	// summary_suggestion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dialogflow_generator#summary_suggestion DialogflowGenerator#summary_suggestion}
	SummarySuggestion *DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestion `field:"optional" json:"summarySuggestion" yaml:"summarySuggestion"`
	// tool_call_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dialogflow_generator#tool_call_info DialogflowGenerator#tool_call_info}
	ToolCallInfo interface{} `field:"optional" json:"toolCallInfo" yaml:"toolCallInfo"`
}

