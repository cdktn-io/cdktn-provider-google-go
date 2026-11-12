// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalineageopenlineagejob


type DataLineageOpenLineageJobInput struct {
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
	Catalog *DataLineageOpenLineageJobInputCatalog `field:"optional" json:"catalog" yaml:"catalog"`
	// symlink block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#symlink DataLineageOpenLineageJob#symlink}
	Symlink interface{} `field:"optional" json:"symlink" yaml:"symlink"`
}

