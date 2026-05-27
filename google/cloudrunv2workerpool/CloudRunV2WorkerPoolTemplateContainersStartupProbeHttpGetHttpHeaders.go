// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2workerpool


type CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetHttpHeaders struct {
	// Required. The header field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.33.0/docs/resources/cloud_run_v2_worker_pool#port CloudRunV2WorkerPool#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Optional. The header field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.33.0/docs/resources/cloud_run_v2_worker_pool#value CloudRunV2WorkerPool#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

