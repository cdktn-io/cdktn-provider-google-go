// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection


type OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageGoogleCloudStorageIcebergStorage struct {
	// The bucket of Google Cloud Storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/oracle_database_goldengate_connection#bucket OracleDatabaseGoldengateConnection#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The project ID of Google Cloud Storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/oracle_database_goldengate_connection#project_id OracleDatabaseGoldengateConnection#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The service account key file of Google Cloud Storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/oracle_database_goldengate_connection#service_account_key_file OracleDatabaseGoldengateConnection#service_account_key_file}
	ServiceAccountKeyFile *string `field:"optional" json:"serviceAccountKeyFile" yaml:"serviceAccountKeyFile"`
}

