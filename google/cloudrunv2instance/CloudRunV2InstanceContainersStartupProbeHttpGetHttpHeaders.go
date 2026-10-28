// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2instance


type CloudRunV2InstanceContainersStartupProbeHttpGetHttpHeaders struct {
	// Required. The header field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/cloud_run_v2_instance#name CloudRunV2Instance#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional. The header field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/cloud_run_v2_instance#value CloudRunV2Instance#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

