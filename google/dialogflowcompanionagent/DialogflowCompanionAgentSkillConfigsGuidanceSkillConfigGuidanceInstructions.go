// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowcompanionagent


type DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructions struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#actions DialogflowCompanionAgent#actions}
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The condition of the instruction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#condition DialogflowCompanionAgent#condition}
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Whether to disable suggested reply generation for this instruction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#disable_suggested_reply DialogflowCompanionAgent#disable_suggested_reply}
	DisableSuggestedReply interface{} `field:"optional" json:"disableSuggestedReply" yaml:"disableSuggestedReply"`
	// The detailed description of this instruction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#display_details DialogflowCompanionAgent#display_details}
	DisplayDetails *string `field:"optional" json:"displayDetails" yaml:"displayDetails"`
	// Display name for the instruction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#display_name DialogflowCompanionAgent#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Event that triggers the guidance instruction. Possible values: ["END_OF_UTTERANCE", "CUSTOMER_MESSAGE", "AGENT_MESSAGE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#trigger_event DialogflowCompanionAgent#trigger_event}
	TriggerEvent *string `field:"optional" json:"triggerEvent" yaml:"triggerEvent"`
}

