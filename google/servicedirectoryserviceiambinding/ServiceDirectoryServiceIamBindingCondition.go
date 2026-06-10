// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicedirectoryserviceiambinding


type ServiceDirectoryServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/service_directory_service_iam_binding#expression ServiceDirectoryServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/service_directory_service_iam_binding#title ServiceDirectoryServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/service_directory_service_iam_binding#description ServiceDirectoryServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

