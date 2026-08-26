// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeinstancefromtemplate


type ComputeInstanceFromTemplateWorkloadIdentityConfig struct {
	// Identity SPIFFE id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/compute_instance_from_template#identity ComputeInstanceFromTemplate#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// Specifies whether identity certificates are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/compute_instance_from_template#identity_certificate_enabled ComputeInstanceFromTemplate#identity_certificate_enabled}
	IdentityCertificateEnabled interface{} `field:"optional" json:"identityCertificateEnabled" yaml:"identityCertificateEnabled"`
}

