// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iapagentregistryagentiambinding

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IapAgentRegistryAgentIamBindingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/iap_agent_registry_agent_iam_binding#agent_id IapAgentRegistryAgentIamBinding#agent_id}.
	AgentId *string `field:"required" json:"agentId" yaml:"agentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/iap_agent_registry_agent_iam_binding#members IapAgentRegistryAgentIamBinding#members}.
	Members *[]*string `field:"required" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/iap_agent_registry_agent_iam_binding#role IapAgentRegistryAgentIamBinding#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/iap_agent_registry_agent_iam_binding#condition IapAgentRegistryAgentIamBinding#condition}
	Condition *IapAgentRegistryAgentIamBindingCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/iap_agent_registry_agent_iam_binding#id IapAgentRegistryAgentIamBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/iap_agent_registry_agent_iam_binding#location IapAgentRegistryAgentIamBinding#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/iap_agent_registry_agent_iam_binding#project IapAgentRegistryAgentIamBinding#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

