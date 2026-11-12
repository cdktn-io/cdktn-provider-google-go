// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalineageopenlineagejob


type DataLineageOpenLineageJobInputSymlink struct {
	// Name of the symlink target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#name DataLineageOpenLineageJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace of the symlink target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#namespace DataLineageOpenLineageJob#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Type of the symlink target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#type DataLineageOpenLineageJob#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

