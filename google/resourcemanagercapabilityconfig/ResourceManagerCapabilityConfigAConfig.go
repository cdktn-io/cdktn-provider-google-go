// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resourcemanagercapabilityconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ResourceManagerCapabilityConfigAConfig struct {
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
	// User-specified identifier of the capability config. Must be 6 to 30 characters, and contain only lowercase letters, numbers, and hyphens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/resource_manager_capability_config#capability_config_id ResourceManagerCapabilityConfigA#capability_config_id}
	CapabilityConfigId *string `field:"required" json:"capabilityConfigId" yaml:"capabilityConfigId"`
	// The parent resource in which to create the capability config. Format: 'folders/{folder_id}', 'organizations/{organization_id}', or 'projects/{project_number}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/resource_manager_capability_config#parent ResourceManagerCapabilityConfigA#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// The capabilities enabled for the resource and its sub-tree. Possible values: "AGENT_MANAGEMENT".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/resource_manager_capability_config#types ResourceManagerCapabilityConfigA#types}
	Types *[]*string `field:"required" json:"types" yaml:"types"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/resource_manager_capability_config#deletion_policy ResourceManagerCapabilityConfigA#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// User-defined name for the capability config. Must be between 4 and 30 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/resource_manager_capability_config#display_name ResourceManagerCapabilityConfigA#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/resource_manager_capability_config#id ResourceManagerCapabilityConfigA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The management project for the capability config.
	//
	// If unspecified, a project will be created automatically.
	// Must be specified for project-scoped capability config.
	// Format: 'projects/{project_number}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/resource_manager_capability_config#management_project ResourceManagerCapabilityConfigA#management_project}
	ManagementProject *string `field:"optional" json:"managementProject" yaml:"managementProject"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/resource_manager_capability_config#timeouts ResourceManagerCapabilityConfigA#timeouts}
	Timeouts *ResourceManagerCapabilityConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

