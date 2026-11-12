// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowcompanionagent


type DialogflowCompanionAgentSkillConfigs struct {
	// guidance_skill_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#guidance_skill_config DialogflowCompanionAgent#guidance_skill_config}
	GuidanceSkillConfig *DialogflowCompanionAgentSkillConfigsGuidanceSkillConfig `field:"optional" json:"guidanceSkillConfig" yaml:"guidanceSkillConfig"`
	// The event that should trigger the skill. Possible values: ["END_OF_UTTERANCE", "CUSTOMER_MESSAGE", "AGENT_MESSAGE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#skill_triggering_event DialogflowCompanionAgent#skill_triggering_event}
	SkillTriggeringEvent *string `field:"optional" json:"skillTriggeringEvent" yaml:"skillTriggeringEvent"`
}

