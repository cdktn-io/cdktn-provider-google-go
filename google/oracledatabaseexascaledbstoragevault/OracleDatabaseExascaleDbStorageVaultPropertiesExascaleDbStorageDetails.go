// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabaseexascaledbstoragevault


type OracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetails struct {
	// The total storage allocation for the ExascaleDbStorageVault, in gigabytes (GB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/oracle_database_exascale_db_storage_vault#total_size_gbs OracleDatabaseExascaleDbStorageVault#total_size_gbs}
	TotalSizeGbs *float64 `field:"required" json:"totalSizeGbs" yaml:"totalSizeGbs"`
}

