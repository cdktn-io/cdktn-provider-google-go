// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowgenerator


type DialogflowGeneratorAgentCoachingContext struct {
	// instructions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dialogflow_generator#instructions DialogflowGenerator#instructions}
	Instructions interface{} `field:"optional" json:"instructions" yaml:"instructions"`
	// Optional. Output language code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dialogflow_generator#output_language_code DialogflowGenerator#output_language_code}
	OutputLanguageCode *string `field:"optional" json:"outputLanguageCode" yaml:"outputLanguageCode"`
	// Optional. The overarching guidance for the agent coaching. This should be set only for v1.5 and later versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dialogflow_generator#overarching_guidance DialogflowGenerator#overarching_guidance}
	OverarchingGuidance *string `field:"optional" json:"overarchingGuidance" yaml:"overarchingGuidance"`
	// Optional. Version of the feature. If not set, default to latest version. Current candidates are ["2.5"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dialogflow_generator#version DialogflowGenerator#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

