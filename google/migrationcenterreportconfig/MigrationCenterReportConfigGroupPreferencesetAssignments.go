// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package migrationcenterreportconfig


type MigrationCenterReportConfigGroupPreferencesetAssignments struct {
	// Name of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/migration_center_report_config#group MigrationCenterReportConfigA#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// Name of the Preference Set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/migration_center_report_config#preference_set MigrationCenterReportConfigA#preference_set}
	PreferenceSet *string `field:"required" json:"preferenceSet" yaml:"preferenceSet"`
}

