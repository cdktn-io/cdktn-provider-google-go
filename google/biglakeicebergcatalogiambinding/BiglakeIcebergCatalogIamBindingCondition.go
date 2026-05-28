// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package biglakeicebergcatalogiambinding


type BiglakeIcebergCatalogIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.34.0/docs/resources/biglake_iceberg_catalog_iam_binding#expression BiglakeIcebergCatalogIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.34.0/docs/resources/biglake_iceberg_catalog_iam_binding#title BiglakeIcebergCatalogIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.34.0/docs/resources/biglake_iceberg_catalog_iam_binding#description BiglakeIcebergCatalogIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

