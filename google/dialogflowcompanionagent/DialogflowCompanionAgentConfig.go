// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowcompanionagent

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DialogflowCompanionAgentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID to use for the companion agent, which will become the final component of the companion agent's resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#companion_agent_id DialogflowCompanionAgent#companion_agent_id}
	CompanionAgentId *string `field:"required" json:"companionAgentId" yaml:"companionAgentId"`
	// The location of the companion agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#location DialogflowCompanionAgent#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// ces_tool_specs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#ces_tool_specs DialogflowCompanionAgent#ces_tool_specs}
	CesToolSpecs interface{} `field:"optional" json:"cesToolSpecs" yaml:"cesToolSpecs"`
	// Whether Terraform will be prevented from destroying the instance.
	//
	// Defaults to "DELETE".
	// When a 'terraform destroy' or 'terraform apply' would delete the instance,
	// the command will fail if this field is set to "PREVENT" in Terraform state.
	// When set to "ABANDON", the command will remove the resource from Terraform
	// management without updating or deleting the resource in the API.
	// When set to "DELETE", deleting the resource is allowed.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#deletion_policy DialogflowCompanionAgent#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Description for the companion agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#description DialogflowCompanionAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Display name for the companion agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#display_name DialogflowCompanionAgent#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#id DialogflowCompanionAgent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#project DialogflowCompanionAgent#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// skill_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#skill_configs DialogflowCompanionAgent#skill_configs}
	SkillConfigs interface{} `field:"optional" json:"skillConfigs" yaml:"skillConfigs"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#timeouts DialogflowCompanionAgent#timeouts}
	Timeouts *DialogflowCompanionAgentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// toolset_tools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#toolset_tools DialogflowCompanionAgent#toolset_tools}
	ToolsetTools interface{} `field:"optional" json:"toolsetTools" yaml:"toolsetTools"`
}

