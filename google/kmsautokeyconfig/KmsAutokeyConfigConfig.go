// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kmsautokeyconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KmsAutokeyConfigConfig struct {
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
	// The folder for which to retrieve config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/kms_autokey_config#folder KmsAutokeyConfig#folder}
	Folder *string `field:"required" json:"folder" yaml:"folder"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/kms_autokey_config#deletion_policy KmsAutokeyConfig#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/kms_autokey_config#id KmsAutokeyConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The target key project for a given folder where KMS Autokey will provision a CryptoKey for any new KeyHandle the Developer creates.
	//
	// Should have the form
	// 'projects/<project_id_or_number>'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/kms_autokey_config#key_project KmsAutokeyConfig#key_project}
	KeyProject *string `field:"optional" json:"keyProject" yaml:"keyProject"`
	// How Autokey determines which project to use when provisioning CMEK keys. Possible values: ["DEDICATED_KEY_PROJECT", "RESOURCE_PROJECT", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/kms_autokey_config#key_project_resolution_mode KmsAutokeyConfig#key_project_resolution_mode}
	KeyProjectResolutionMode *string `field:"optional" json:"keyProjectResolutionMode" yaml:"keyProjectResolutionMode"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/kms_autokey_config#timeouts KmsAutokeyConfig#timeouts}
	Timeouts *KmsAutokeyConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

