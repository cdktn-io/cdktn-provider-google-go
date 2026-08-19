// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryservice

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgentRegistryServiceConfig struct {
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
	// The location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/agent_registry_service#location AgentRegistryService#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/agent_registry_service#service_id AgentRegistryService#service_id}
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// agent_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/agent_registry_service#agent_spec AgentRegistryService#agent_spec}
	AgentSpec *AgentRegistryServiceAgentSpec `field:"optional" json:"agentSpec" yaml:"agentSpec"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/agent_registry_service#deletion_policy AgentRegistryService#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// The description of the Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/agent_registry_service#description AgentRegistryService#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User-defined display name for the Service. Can have a maximum length of 63 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/agent_registry_service#display_name AgentRegistryService#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// endpoint_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/agent_registry_service#endpoint_spec AgentRegistryService#endpoint_spec}
	EndpointSpec *AgentRegistryServiceEndpointSpec `field:"optional" json:"endpointSpec" yaml:"endpointSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/agent_registry_service#id AgentRegistryService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// interfaces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/agent_registry_service#interfaces AgentRegistryService#interfaces}
	Interfaces interface{} `field:"optional" json:"interfaces" yaml:"interfaces"`
	// mcp_server_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/agent_registry_service#mcp_server_spec AgentRegistryService#mcp_server_spec}
	McpServerSpec *AgentRegistryServiceMcpServerSpec `field:"optional" json:"mcpServerSpec" yaml:"mcpServerSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/agent_registry_service#project AgentRegistryService#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/agent_registry_service#timeouts AgentRegistryService#timeouts}
	Timeouts *AgentRegistryServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

