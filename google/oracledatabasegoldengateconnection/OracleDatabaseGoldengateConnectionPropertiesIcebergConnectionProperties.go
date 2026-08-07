// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection


type OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionProperties struct {
	// catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/oracle_database_goldengate_connection#catalog OracleDatabaseGoldengateConnection#catalog}
	Catalog *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalog `field:"required" json:"catalog" yaml:"catalog"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/oracle_database_goldengate_connection#storage OracleDatabaseGoldengateConnection#storage}
	Storage *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorage `field:"required" json:"storage" yaml:"storage"`
	// The technology type of Iceberg connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/oracle_database_goldengate_connection#technology_type OracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"required" json:"technologyType" yaml:"technologyType"`
}

