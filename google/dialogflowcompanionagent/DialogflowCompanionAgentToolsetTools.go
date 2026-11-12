// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowcompanionagent


type DialogflowCompanionAgentToolsetTools struct {
	// The name of the toolset to retrieve the schema for. Format: 'projects/{project}/locations/{location}/apps/{app}/toolsets/{toolset}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#toolset DialogflowCompanionAgent#toolset}
	Toolset *string `field:"required" json:"toolset" yaml:"toolset"`
	// Indicates whether the tool requires human confirmation. Possible values: ["REQUIRED", "NOT_REQUIRED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#confirmation_requirement DialogflowCompanionAgent#confirmation_requirement}
	ConfirmationRequirement *string `field:"optional" json:"confirmationRequirement" yaml:"confirmationRequirement"`
	// The 'operationId' field of the OpenAPI endpoint. The 'operationId' must be present in the toolset's definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#operation_id DialogflowCompanionAgent#operation_id}
	OperationId *string `field:"optional" json:"operationId" yaml:"operationId"`
	// parameter_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#parameter_overrides DialogflowCompanionAgent#parameter_overrides}
	ParameterOverrides interface{} `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
	// Indicates whether the toolset tool is enabled in proactive mode.
	//
	// NOTE: At least one of 'proactive_enabled' or 'reactive_enabled' should be 'true' else the toolset tool will be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#proactive_enabled DialogflowCompanionAgent#proactive_enabled}
	ProactiveEnabled interface{} `field:"optional" json:"proactiveEnabled" yaml:"proactiveEnabled"`
	// Indicates whether the toolset tool is enabled in reactive mode.
	//
	// NOTE: At least one of 'reactive_enabled' or 'proactive_enabled' should be 'true' else the toolset tool will be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#reactive_enabled DialogflowCompanionAgent#reactive_enabled}
	ReactiveEnabled interface{} `field:"optional" json:"reactiveEnabled" yaml:"reactiveEnabled"`
}

