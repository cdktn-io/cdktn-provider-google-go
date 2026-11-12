// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataflowconfigstoresetting


type DataflowConfigStoreSettingValue struct {
	// Represents a boolean value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataflow_config_store_setting#bool_value DataflowConfigStoreSetting#bool_value}
	BoolValue interface{} `field:"optional" json:"boolValue" yaml:"boolValue"`
	// Represents a string value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataflow_config_store_setting#string_value DataflowConfigStoreSetting#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

