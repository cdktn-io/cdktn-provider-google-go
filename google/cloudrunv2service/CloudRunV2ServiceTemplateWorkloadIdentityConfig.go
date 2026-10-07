// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2service


type CloudRunV2ServiceTemplateWorkloadIdentityConfig struct {
	// The Revision's SPIFFE workload identity. Enables provisioning of SPIFFE workload certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/cloud_run_v2_service#identity CloudRunV2Service#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// Controls whether an instance receives a MWLID certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/cloud_run_v2_service#identity_certificate_enabled CloudRunV2Service#identity_certificate_enabled}
	IdentityCertificateEnabled interface{} `field:"optional" json:"identityCertificateEnabled" yaml:"identityCertificateEnabled"`
	// The type of identity to use. Possible values: ["IDENTITY_TYPE_SERVICE_ACCOUNT", "IDENTITY_TYPE_WORKLOAD_IDENTITY", "IDENTITY_TYPE_AGENT_IDENTITY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/cloud_run_v2_service#identity_type CloudRunV2Service#identity_type}
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
}

