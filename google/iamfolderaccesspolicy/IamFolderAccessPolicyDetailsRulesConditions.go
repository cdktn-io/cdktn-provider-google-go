// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamfolderaccesspolicy


type IamFolderAccessPolicyDetailsRulesConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/iam_folder_access_policy#service IamFolderAccessPolicy#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/iam_folder_access_policy#expression IamFolderAccessPolicy#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

