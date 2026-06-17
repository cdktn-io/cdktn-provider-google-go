// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection


type OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalog struct {
	// The type of Iceberg catalog. Possible values: GLUE HADOOP NESSIE POLARIS REST.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/oracle_database_goldengate_connection#catalog_type OracleDatabaseGoldengateConnection#catalog_type}
	CatalogType *string `field:"required" json:"catalogType" yaml:"catalogType"`
	// glue_iceberg_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/oracle_database_goldengate_connection#glue_iceberg_catalog OracleDatabaseGoldengateConnection#glue_iceberg_catalog}
	GlueIcebergCatalog *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogGlueIcebergCatalog `field:"optional" json:"glueIcebergCatalog" yaml:"glueIcebergCatalog"`
	// nessie_iceberg_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/oracle_database_goldengate_connection#nessie_iceberg_catalog OracleDatabaseGoldengateConnection#nessie_iceberg_catalog}
	NessieIcebergCatalog *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogNessieIcebergCatalog `field:"optional" json:"nessieIcebergCatalog" yaml:"nessieIcebergCatalog"`
	// polaris_iceberg_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/oracle_database_goldengate_connection#polaris_iceberg_catalog OracleDatabaseGoldengateConnection#polaris_iceberg_catalog}
	PolarisIcebergCatalog *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogPolarisIcebergCatalog `field:"optional" json:"polarisIcebergCatalog" yaml:"polarisIcebergCatalog"`
	// rest_iceberg_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/oracle_database_goldengate_connection#rest_iceberg_catalog OracleDatabaseGoldengateConnection#rest_iceberg_catalog}
	RestIcebergCatalog *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogRestIcebergCatalog `field:"optional" json:"restIcebergCatalog" yaml:"restIcebergCatalog"`
}

