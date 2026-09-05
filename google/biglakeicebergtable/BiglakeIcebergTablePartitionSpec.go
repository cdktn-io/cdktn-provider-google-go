// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package biglakeicebergtable


type BiglakeIcebergTablePartitionSpec struct {
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/biglake_iceberg_table#fields BiglakeIcebergTable#fields}
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
}

