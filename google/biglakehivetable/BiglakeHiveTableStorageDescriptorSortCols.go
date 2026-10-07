// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package biglakehivetable


type BiglakeHiveTableStorageDescriptorSortCols struct {
	// The column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/biglake_hive_table#col BiglakeHiveTable#col}
	Col *string `field:"required" json:"col" yaml:"col"`
	// Sort order: 1 for Ascending, 0 for Descending.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/biglake_hive_table#order BiglakeHiveTable#order}
	Order *float64 `field:"required" json:"order" yaml:"order"`
}

