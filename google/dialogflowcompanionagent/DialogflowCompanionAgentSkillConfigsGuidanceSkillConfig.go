// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowcompanionagent


type DialogflowCompanionAgentSkillConfigsGuidanceSkillConfig struct {
	// guidance_instructions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#guidance_instructions DialogflowCompanionAgent#guidance_instructions}
	GuidanceInstructions interface{} `field:"optional" json:"guidanceInstructions" yaml:"guidanceInstructions"`
	// This is specific additional guidance that can be configured by the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#overarching_guidance DialogflowCompanionAgent#overarching_guidance}
	OverarchingGuidance *string `field:"optional" json:"overarchingGuidance" yaml:"overarchingGuidance"`
}

