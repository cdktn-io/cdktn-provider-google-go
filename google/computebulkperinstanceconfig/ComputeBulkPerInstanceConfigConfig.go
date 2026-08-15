// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computebulkperinstanceconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ComputeBulkPerInstanceConfigConfig struct {
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
	// The instance group manager this instance config is part of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_bulk_per_instance_config#instance_group_manager ComputeBulkPerInstanceConfig#instance_group_manager}
	InstanceGroupManager *string `field:"required" json:"instanceGroupManager" yaml:"instanceGroupManager"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_bulk_per_instance_config#deletion_policy ComputeBulkPerInstanceConfig#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_bulk_per_instance_config#id ComputeBulkPerInstanceConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// instances block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_bulk_per_instance_config#instances ComputeBulkPerInstanceConfig#instances}
	Instances interface{} `field:"optional" json:"instances" yaml:"instances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_bulk_per_instance_config#project ComputeBulkPerInstanceConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_bulk_per_instance_config#timeouts ComputeBulkPerInstanceConfig#timeouts}
	Timeouts *ComputeBulkPerInstanceConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Zone where the containing instance group manager is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_bulk_per_instance_config#zone ComputeBulkPerInstanceConfig#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

