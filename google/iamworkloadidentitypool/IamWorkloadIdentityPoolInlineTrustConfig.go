// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamworkloadidentitypool


type IamWorkloadIdentityPoolInlineTrustConfig struct {
	// additional_trust_bundles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/iam_workload_identity_pool#additional_trust_bundles IamWorkloadIdentityPool#additional_trust_bundles}
	AdditionalTrustBundles interface{} `field:"optional" json:"additionalTrustBundles" yaml:"additionalTrustBundles"`
}

