// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabaseexadbvmcluster


type OracleDatabaseExadbVmClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_exadb_vm_cluster#create OracleDatabaseExadbVmCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_exadb_vm_cluster#delete OracleDatabaseExadbVmCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_exadb_vm_cluster#update OracleDatabaseExadbVmCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

