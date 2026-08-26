// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package migrationcenterassetsexportjob


type MigrationCenterAssetsExportJobCondition struct {
	// Assets filter, supports the same syntax as asset listing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/migration_center_assets_export_job#filter MigrationCenterAssetsExportJob#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
}

