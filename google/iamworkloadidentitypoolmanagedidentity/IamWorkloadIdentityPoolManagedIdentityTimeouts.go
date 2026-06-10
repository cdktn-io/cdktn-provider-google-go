// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamworkloadidentitypoolmanagedidentity


type IamWorkloadIdentityPoolManagedIdentityTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/iam_workload_identity_pool_managed_identity#create IamWorkloadIdentityPoolManagedIdentity#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/iam_workload_identity_pool_managed_identity#delete IamWorkloadIdentityPoolManagedIdentity#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/iam_workload_identity_pool_managed_identity#update IamWorkloadIdentityPoolManagedIdentity#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

