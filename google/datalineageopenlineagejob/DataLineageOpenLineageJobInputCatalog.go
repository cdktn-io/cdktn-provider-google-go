// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalineageopenlineagejob


type DataLineageOpenLineageJobInputCatalog struct {
	// Catalog framework.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#framework DataLineageOpenLineageJob#framework}
	Framework *string `field:"required" json:"framework" yaml:"framework"`
	// Catalog entity name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#name DataLineageOpenLineageJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Catalog entity type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#type DataLineageOpenLineageJob#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

