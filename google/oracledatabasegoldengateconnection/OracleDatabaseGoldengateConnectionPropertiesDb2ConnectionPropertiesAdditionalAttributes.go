// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection


type OracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesAdditionalAttributes struct {
	// The name of the property entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/oracle_database_goldengate_connection#key OracleDatabaseGoldengateConnection#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the property entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/oracle_database_goldengate_connection#value OracleDatabaseGoldengateConnection#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

