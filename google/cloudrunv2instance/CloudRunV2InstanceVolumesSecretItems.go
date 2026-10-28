// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2instance


type CloudRunV2InstanceVolumesSecretItems struct {
	// The relative path of the secret in the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/cloud_run_v2_instance#path CloudRunV2Instance#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Integer octal mode bits to use on this file, must be a value between 01 and 0777 (octal).
	//
	// If 0 or not set, the Volume's default mode will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/cloud_run_v2_instance#mode CloudRunV2Instance#mode}
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
	// The Cloud Secret Manager secret version.
	//
	// Can be 'latest' for the latest value or an integer for a specific version
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/cloud_run_v2_instance#version CloudRunV2Instance#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

