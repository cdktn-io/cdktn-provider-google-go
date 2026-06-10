// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package biglakeicebergtable


type BiglakeIcebergTableTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/biglake_iceberg_table#create BiglakeIcebergTable#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/biglake_iceberg_table#delete BiglakeIcebergTable#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/biglake_iceberg_table#update BiglakeIcebergTable#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

