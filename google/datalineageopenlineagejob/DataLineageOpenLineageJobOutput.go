// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalineageopenlineagejob


type DataLineageOpenLineageJobOutput struct {
	// Name of the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#name DataLineageOpenLineageJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace of the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#namespace DataLineageOpenLineageJob#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#catalog DataLineageOpenLineageJob#catalog}
	Catalog *DataLineageOpenLineageJobOutputCatalog `field:"optional" json:"catalog" yaml:"catalog"`
	// column_lineage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#column_lineage DataLineageOpenLineageJob#column_lineage}
	ColumnLineage *DataLineageOpenLineageJobOutputColumnLineage `field:"optional" json:"columnLineage" yaml:"columnLineage"`
	// symlink block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#symlink DataLineageOpenLineageJob#symlink}
	Symlink interface{} `field:"optional" json:"symlink" yaml:"symlink"`
}

