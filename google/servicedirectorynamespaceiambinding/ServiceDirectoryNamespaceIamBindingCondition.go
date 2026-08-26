// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicedirectorynamespaceiambinding


type ServiceDirectoryNamespaceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/service_directory_namespace_iam_binding#expression ServiceDirectoryNamespaceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/service_directory_namespace_iam_binding#title ServiceDirectoryNamespaceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/service_directory_namespace_iam_binding#description ServiceDirectoryNamespaceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

