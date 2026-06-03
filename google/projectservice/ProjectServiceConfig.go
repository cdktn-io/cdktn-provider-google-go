// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package projectservice

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ProjectServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/project_service#service ProjectService#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/project_service#deletion_policy ProjectService#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/project_service#disable_dependent_services ProjectService#disable_dependent_services}.
	DisableDependentServices interface{} `field:"optional" json:"disableDependentServices" yaml:"disableDependentServices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/project_service#disable_on_destroy ProjectService#disable_on_destroy}.
	DisableOnDestroy interface{} `field:"optional" json:"disableOnDestroy" yaml:"disableOnDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/project_service#id ProjectService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/project_service#project ProjectService#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/project_service#timeouts ProjectService#timeouts}
	Timeouts *ProjectServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

