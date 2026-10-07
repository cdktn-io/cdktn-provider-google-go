// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowgenerator


type DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultError struct {
	// The error message of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dialogflow_generator#message DialogflowGenerator#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Specifies whether the tool call is retryable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dialogflow_generator#retryable DialogflowGenerator#retryable}
	Retryable interface{} `field:"optional" json:"retryable" yaml:"retryable"`
}

