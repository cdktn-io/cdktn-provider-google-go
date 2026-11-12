// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan


type DataplexDatascanDataQualitySpecRulesDebugQueries struct {
	// Specifies the SQL statement to be executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataplex_datascan#sql_statement DataplexDatascan#sql_statement}
	SqlStatement *string `field:"required" json:"sqlStatement" yaml:"sqlStatement"`
	// Specifies the description of the debug query. The maximum length is 1,024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataplex_datascan#description DataplexDatascan#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

