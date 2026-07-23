// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package migrationcenterassetsexportjob


type MigrationCenterAssetsExportJobSignedUriDestination struct {
	// The file format to export. Possible values: CSV XLSX.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/migration_center_assets_export_job#file_format MigrationCenterAssetsExportJob#file_format}
	FileFormat *string `field:"required" json:"fileFormat" yaml:"fileFormat"`
}

