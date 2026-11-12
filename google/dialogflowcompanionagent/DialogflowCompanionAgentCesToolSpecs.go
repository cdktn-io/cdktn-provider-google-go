// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowcompanionagent


type DialogflowCompanionAgentCesToolSpecs struct {
	// The resource name of the CES tool. Format: 'projects/{project}/locations/{location}/apps/{app}/tools/{tool}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#ces_tool DialogflowCompanionAgent#ces_tool}
	CesTool *string `field:"optional" json:"cesTool" yaml:"cesTool"`
	// Indicates whether the tool requires human confirmation. Possible values: ["REQUIRED", "NOT_REQUIRED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#confirmation_requirement DialogflowCompanionAgent#confirmation_requirement}
	ConfirmationRequirement *string `field:"optional" json:"confirmationRequirement" yaml:"confirmationRequirement"`
	// parameter_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#parameter_overrides DialogflowCompanionAgent#parameter_overrides}
	ParameterOverrides interface{} `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
	// Indicates whether the tool is enabled in proactive mode.
	//
	// All tools are enabled in proactive mode by default. To disable a tool in proactive mode, set this field to 'false'.
	// NOTE: At least one of 'proactive_enabled' or 'reactive_enabled' should be 'true' else the tool will be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#proactive_enabled DialogflowCompanionAgent#proactive_enabled}
	ProactiveEnabled interface{} `field:"optional" json:"proactiveEnabled" yaml:"proactiveEnabled"`
	// Indicates whether the tool is enabled in reactive mode.
	//
	// All tools are disabled in reactive mode by default. To enable a tool in reactive mode, set this field to 'true'.
	// NOTE: At least one of 'reactive_enabled' or 'proactive_enabled' should be 'true' else the tool will be ignored. CompanionAgent config will throw an error if any unsupported tool is enabled in reactive mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#reactive_enabled DialogflowCompanionAgent#reactive_enabled}
	ReactiveEnabled interface{} `field:"optional" json:"reactiveEnabled" yaml:"reactiveEnabled"`
}

