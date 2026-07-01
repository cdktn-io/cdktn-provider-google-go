// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection


type OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorage struct {
	// The account of Azure Data Lake Storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/oracle_database_goldengate_connection#azure_account OracleDatabaseGoldengateConnection#azure_account}
	AzureAccount *string `field:"required" json:"azureAccount" yaml:"azureAccount"`
	// The container of Azure Data Lake Storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/oracle_database_goldengate_connection#container OracleDatabaseGoldengateConnection#container}
	Container *string `field:"required" json:"container" yaml:"container"`
	// The account key of Azure Data Lake Storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/oracle_database_goldengate_connection#account_key_secret OracleDatabaseGoldengateConnection#account_key_secret}
	AccountKeySecret *string `field:"optional" json:"accountKeySecret" yaml:"accountKeySecret"`
	// The endpoint of Azure Data Lake Storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/oracle_database_goldengate_connection#endpoint OracleDatabaseGoldengateConnection#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

