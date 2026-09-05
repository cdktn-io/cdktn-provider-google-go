// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection


type OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties struct {
	// The content of the private key file (PEM file) corresponding to the API key of the fingerprint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/oracle_database_goldengate_connection#private_key_file OracleDatabaseGoldengateConnection#private_key_file}
	PrivateKeyFile *string `field:"optional" json:"privateKeyFile" yaml:"privateKeyFile"`
	// The passphrase of the private key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/oracle_database_goldengate_connection#private_key_passphrase_secret OracleDatabaseGoldengateConnection#private_key_passphrase_secret}
	PrivateKeyPassphraseSecret *string `field:"optional" json:"privateKeyPassphraseSecret" yaml:"privateKeyPassphraseSecret"`
	// The fingerprint of the API Key of the user specified by the userId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/oracle_database_goldengate_connection#public_key_fingerprint OracleDatabaseGoldengateConnection#public_key_fingerprint}
	PublicKeyFingerprint *string `field:"optional" json:"publicKeyFingerprint" yaml:"publicKeyFingerprint"`
	// The name of the region. e.g.: us-ashburn-1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/oracle_database_goldengate_connection#region OracleDatabaseGoldengateConnection#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The technology type of OracleNosqlConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/oracle_database_goldengate_connection#technology_type OracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
	// The OCID of the OCI tenancy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/oracle_database_goldengate_connection#tenancy_id OracleDatabaseGoldengateConnection#tenancy_id}
	TenancyId *string `field:"optional" json:"tenancyId" yaml:"tenancyId"`
	// Specifies that the user intends to authenticate to the instance using a resource principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/oracle_database_goldengate_connection#use_resource_principal OracleDatabaseGoldengateConnection#use_resource_principal}
	UseResourcePrincipal interface{} `field:"optional" json:"useResourcePrincipal" yaml:"useResourcePrincipal"`
	// The OCID of the OCI user who will access the Oracle NoSQL database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/oracle_database_goldengate_connection#user_id OracleDatabaseGoldengateConnection#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

