// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection


type OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAmazonS3IcebergStorage struct {
	// The access key ID of Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/oracle_database_goldengate_connection#access_key_id OracleDatabaseGoldengateConnection#access_key_id}
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// The bucket of Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/oracle_database_goldengate_connection#bucket OracleDatabaseGoldengateConnection#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The region of Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/oracle_database_goldengate_connection#region OracleDatabaseGoldengateConnection#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// The scheme type of Amazon S3. Possible values: S3 S3A.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/oracle_database_goldengate_connection#scheme_type OracleDatabaseGoldengateConnection#scheme_type}
	SchemeType *string `field:"required" json:"schemeType" yaml:"schemeType"`
	// The endpoint of Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/oracle_database_goldengate_connection#endpoint OracleDatabaseGoldengateConnection#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The secret access key of Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/oracle_database_goldengate_connection#secret_access_key_secret OracleDatabaseGoldengateConnection#secret_access_key_secret}
	SecretAccessKeySecret *string `field:"optional" json:"secretAccessKeySecret" yaml:"secretAccessKeySecret"`
}

