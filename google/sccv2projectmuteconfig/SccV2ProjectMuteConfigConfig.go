// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sccv2projectmuteconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SccV2ProjectMuteConfigConfig struct {
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
	// An expression that defines the filter to apply across create/update events of findings.
	//
	// While creating a filter string, be mindful of
	// the scope in which the mute configuration is being created. E.g.,
	// If a filter contains project = X but is created under the
	// project = Y scope, it might not match any findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/scc_v2_project_mute_config#filter SccV2ProjectMuteConfig#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Unique identifier provided by the client within the parent scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/scc_v2_project_mute_config#mute_config_id SccV2ProjectMuteConfig#mute_config_id}
	MuteConfigId *string `field:"required" json:"muteConfigId" yaml:"muteConfigId"`
	// The type of the mute config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/scc_v2_project_mute_config#type SccV2ProjectMuteConfig#type}
	Type *string `field:"required" json:"type" yaml:"type"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/scc_v2_project_mute_config#deletion_policy SccV2ProjectMuteConfig#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// A description of the mute config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/scc_v2_project_mute_config#description SccV2ProjectMuteConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/scc_v2_project_mute_config#id SccV2ProjectMuteConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// location Id is provided by project. If not provided, Use global as default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/scc_v2_project_mute_config#location SccV2ProjectMuteConfig#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/scc_v2_project_mute_config#project SccV2ProjectMuteConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/scc_v2_project_mute_config#timeouts SccV2ProjectMuteConfig#timeouts}
	Timeouts *SccV2ProjectMuteConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

