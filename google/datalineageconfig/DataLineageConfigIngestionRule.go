// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalineageconfig


type DataLineageConfigIngestionRule struct {
	// integration_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/data_lineage_config#integration_selector DataLineageConfig#integration_selector}
	IntegrationSelector *DataLineageConfigIngestionRuleIntegrationSelector `field:"required" json:"integrationSelector" yaml:"integrationSelector"`
	// lineage_enablement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/data_lineage_config#lineage_enablement DataLineageConfig#lineage_enablement}
	LineageEnablement *DataLineageConfigIngestionRuleLineageEnablement `field:"required" json:"lineageEnablement" yaml:"lineageEnablement"`
}

