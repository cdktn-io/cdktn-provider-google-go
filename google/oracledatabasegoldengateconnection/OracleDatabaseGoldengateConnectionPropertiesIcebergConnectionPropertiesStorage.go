// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection


type OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorage struct {
	// The type of Iceberg storage. Possible values: AMAZON_S3 GOOGLE_CLOUD_STORAGE AZURE_DATA_LAKE_STORAGE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/oracle_database_goldengate_connection#storage_type OracleDatabaseGoldengateConnection#storage_type}
	StorageType *string `field:"required" json:"storageType" yaml:"storageType"`
	// amazon_s3_iceberg_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/oracle_database_goldengate_connection#amazon_s3_iceberg_storage OracleDatabaseGoldengateConnection#amazon_s3_iceberg_storage}
	AmazonS3IcebergStorage *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAmazonS3IcebergStorage `field:"optional" json:"amazonS3IcebergStorage" yaml:"amazonS3IcebergStorage"`
	// azure_data_lake_storage_iceberg_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/oracle_database_goldengate_connection#azure_data_lake_storage_iceberg_storage OracleDatabaseGoldengateConnection#azure_data_lake_storage_iceberg_storage}
	AzureDataLakeStorageIcebergStorage *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorage `field:"optional" json:"azureDataLakeStorageIcebergStorage" yaml:"azureDataLakeStorageIcebergStorage"`
	// google_cloud_storage_iceberg_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/oracle_database_goldengate_connection#google_cloud_storage_iceberg_storage OracleDatabaseGoldengateConnection#google_cloud_storage_iceberg_storage}
	GoogleCloudStorageIcebergStorage *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageGoogleCloudStorageIcebergStorage `field:"optional" json:"googleCloudStorageIcebergStorage" yaml:"googleCloudStorageIcebergStorage"`
}

