// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2instance


type CloudRunV2InstanceContainersEnvValueSourceSecretKeyRef struct {
	// The name of the secret in Cloud Secret Manager.
	//
	// Format: {secretName} if the secret is in the same project. projects/{project}/secrets/{secretName} if the secret is in a different project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/cloud_run_v2_instance#secret CloudRunV2Instance#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
	// The Cloud Secret Manager secret version.
	//
	// Can be 'latest' for the latest value or an integer for a specific version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/cloud_run_v2_instance#version CloudRunV2Instance#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

