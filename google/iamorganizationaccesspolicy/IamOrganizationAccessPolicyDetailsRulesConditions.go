// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamorganizationaccesspolicy


type IamOrganizationAccessPolicyDetailsRulesConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/iam_organization_access_policy#service IamOrganizationAccessPolicy#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/iam_organization_access_policy#expression IamOrganizationAccessPolicy#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

