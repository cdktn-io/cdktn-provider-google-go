// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataformrepository


type DataformRepositoryWorkspaceCompilationOverrides struct {
	// The default database (Google Cloud project ID).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/dataform_repository#default_database DataformRepository#default_database}
	DefaultDatabase *string `field:"optional" json:"defaultDatabase" yaml:"defaultDatabase"`
	// The suffix that should be appended to all schema (BigQuery dataset ID) names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/dataform_repository#schema_suffix DataformRepository#schema_suffix}
	SchemaSuffix *string `field:"optional" json:"schemaSuffix" yaml:"schemaSuffix"`
	// The prefix that should be prepended to all table names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/dataform_repository#table_prefix DataformRepository#table_prefix}
	TablePrefix *string `field:"optional" json:"tablePrefix" yaml:"tablePrefix"`
}

