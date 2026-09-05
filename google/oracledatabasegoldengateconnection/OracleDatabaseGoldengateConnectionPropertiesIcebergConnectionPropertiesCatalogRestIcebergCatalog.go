// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection


type OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogRestIcebergCatalog struct {
	// The REST uri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/oracle_database_goldengate_connection#uri OracleDatabaseGoldengateConnection#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// The content of the configuration file containing additional properties for the REST catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/oracle_database_goldengate_connection#properties OracleDatabaseGoldengateConnection#properties}
	Properties *string `field:"optional" json:"properties" yaml:"properties"`
}

