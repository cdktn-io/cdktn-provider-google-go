// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamorganizationaccesspolicy


type IamOrganizationAccessPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/iam_organization_access_policy#create IamOrganizationAccessPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/iam_organization_access_policy#delete IamOrganizationAccessPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/iam_organization_access_policy#update IamOrganizationAccessPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

