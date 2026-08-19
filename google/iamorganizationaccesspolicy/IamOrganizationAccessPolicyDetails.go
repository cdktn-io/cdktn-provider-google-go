// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamorganizationaccesspolicy


type IamOrganizationAccessPolicyDetails struct {
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/iam_organization_access_policy#rules IamOrganizationAccessPolicy#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

