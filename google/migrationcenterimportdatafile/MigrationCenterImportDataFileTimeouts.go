// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package migrationcenterimportdatafile


type MigrationCenterImportDataFileTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/migration_center_import_data_file#create MigrationCenterImportDataFile#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/migration_center_import_data_file#delete MigrationCenterImportDataFile#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

