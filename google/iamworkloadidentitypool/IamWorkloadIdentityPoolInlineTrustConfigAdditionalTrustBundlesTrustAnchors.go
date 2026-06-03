// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamworkloadidentitypool


type IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesTrustAnchors struct {
	// PEM certificate of the PKI used for validation. Must only contain one ca certificate(either root or intermediate cert).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/iam_workload_identity_pool#pem_certificate IamWorkloadIdentityPool#pem_certificate}
	PemCertificate *string `field:"required" json:"pemCertificate" yaml:"pemCertificate"`
}

