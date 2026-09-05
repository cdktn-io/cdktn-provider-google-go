// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamprojectaccesspolicy


type IamProjectAccessPolicyDetailsRulesConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/iam_project_access_policy#service IamProjectAccessPolicy#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/iam_project_access_policy#expression IamProjectAccessPolicy#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

