// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataflowconfigstoresetting

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataflowConfigStoreSettingConfig struct {
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
	// The ID of the config store setting (e.g. 'default_launcher_machine_type').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataflow_config_store_setting#config_store_setting_id DataflowConfigStoreSetting#config_store_setting_id}
	ConfigStoreSettingId *string `field:"required" json:"configStoreSettingId" yaml:"configStoreSettingId"`
	// The Dataflow region for the setting (e.g. 'us-central1').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataflow_config_store_setting#location DataflowConfigStoreSetting#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The parent resource in which to create the config store setting.
	//
	// Must be in one of the following formats:
	// * 'projects/{project_number}' (numeric project number is required)
	// * 'folders/{folder_id}'
	// * 'organizations/{organization_id}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataflow_config_store_setting#parent DataflowConfigStoreSetting#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataflow_config_store_setting#value DataflowConfigStoreSetting#value}
	Value *DataflowConfigStoreSettingValue `field:"required" json:"value" yaml:"value"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataflow_config_store_setting#deletion_policy DataflowConfigStoreSetting#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataflow_config_store_setting#id DataflowConfigStoreSetting#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataflow_config_store_setting#timeouts DataflowConfigStoreSetting#timeouts}
	Timeouts *DataflowConfigStoreSettingTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

