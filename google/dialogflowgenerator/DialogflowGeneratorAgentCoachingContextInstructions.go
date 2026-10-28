// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowgenerator


type DialogflowGeneratorAgentCoachingContextInstructions struct {
	// Optional. The action that human agent should take.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dialogflow_generator#agent_action DialogflowGenerator#agent_action}
	AgentAction *string `field:"optional" json:"agentAction" yaml:"agentAction"`
	// Optional. The condition of the instruction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dialogflow_generator#condition DialogflowGenerator#condition}
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Optional. The detailed description of this instruction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dialogflow_generator#display_details DialogflowGenerator#display_details}
	DisplayDetails *string `field:"optional" json:"displayDetails" yaml:"displayDetails"`
	// Optional. Display name for the instruction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dialogflow_generator#display_name DialogflowGenerator#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Optional. The action that system should take.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dialogflow_generator#system_action DialogflowGenerator#system_action}
	SystemAction *string `field:"optional" json:"systemAction" yaml:"systemAction"`
	// Optional. The trigger event of the instruction. Possible values: ["END_OF_UTTERANCE", "MANUAL_CALL", "CUSTOMER_MESSAGE", "AGENT_MESSAGE", "TOOL_CALL_COMPLETION"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dialogflow_generator#triggering_event DialogflowGenerator#triggering_event}
	TriggeringEvent *string `field:"optional" json:"triggeringEvent" yaml:"triggeringEvent"`
}

