// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalineageconfig


type DataLineageConfigIngestionRuleIntegrationSelector struct {
	// Integration to which the rule applies. Possible values: ["DATAPROC", "LOOKER_CORE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/data_lineage_config#integration DataLineageConfig#integration}
	Integration *string `field:"required" json:"integration" yaml:"integration"`
}

