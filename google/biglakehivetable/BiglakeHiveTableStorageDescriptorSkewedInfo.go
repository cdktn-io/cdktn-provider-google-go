// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package biglakehivetable


type BiglakeHiveTableStorageDescriptorSkewedInfo struct {
	// The column names that are skewed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/biglake_hive_table#skewed_col_names BiglakeHiveTable#skewed_col_names}
	SkewedColNames *[]*string `field:"required" json:"skewedColNames" yaml:"skewedColNames"`
	// skewed_col_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/biglake_hive_table#skewed_col_values BiglakeHiveTable#skewed_col_values}
	SkewedColValues interface{} `field:"required" json:"skewedColValues" yaml:"skewedColValues"`
	// skewed_key_values_locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/biglake_hive_table#skewed_key_values_locations BiglakeHiveTable#skewed_key_values_locations}
	SkewedKeyValuesLocations interface{} `field:"required" json:"skewedKeyValuesLocations" yaml:"skewedKeyValuesLocations"`
}

