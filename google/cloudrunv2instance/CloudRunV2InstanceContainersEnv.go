// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2instance


type CloudRunV2InstanceContainersEnv struct {
	// Name of the environment variable. Must be a C_IDENTIFIER, and may not exceed 32768 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/cloud_run_v2_instance#name CloudRunV2Instance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Literal value of the environment variable.
	//
	// Defaults to "" and the maximum allowed length is 32768 characters. Variable references are not supported in Cloud Run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/cloud_run_v2_instance#value CloudRunV2Instance#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// value_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/cloud_run_v2_instance#value_source CloudRunV2Instance#value_source}
	ValueSource *CloudRunV2InstanceContainersEnvValueSource `field:"optional" json:"valueSource" yaml:"valueSource"`
}

