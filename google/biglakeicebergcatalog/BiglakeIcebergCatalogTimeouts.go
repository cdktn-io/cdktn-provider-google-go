// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package biglakeicebergcatalog


type BiglakeIcebergCatalogTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/biglake_iceberg_catalog#create BiglakeIcebergCatalog#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/biglake_iceberg_catalog#delete BiglakeIcebergCatalog#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/biglake_iceberg_catalog#update BiglakeIcebergCatalog#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

