// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection


type OracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionProperties struct {
	// The content of the service account key file containing the credentials required to use Google Pub/Sub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#service_account_key_file OracleDatabaseGoldengateConnection#service_account_key_file}
	ServiceAccountKeyFile *string `field:"optional" json:"serviceAccountKeyFile" yaml:"serviceAccountKeyFile"`
	// The technology type of GooglePubsubConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#technology_type OracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
}

