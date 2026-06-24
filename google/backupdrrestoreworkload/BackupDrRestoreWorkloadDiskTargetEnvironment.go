// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupdrrestoreworkload


type BackupDrRestoreWorkloadDiskTargetEnvironment struct {
	// Required. Target project for the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/backup_dr_restore_workload#project BackupDrRestoreWorkload#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Required. Target zone for the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/backup_dr_restore_workload#zone BackupDrRestoreWorkload#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// If true, use the BackupDR P4SA credentials for same-project restores. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/backup_dr_restore_workload#use_project_service_account BackupDrRestoreWorkload#use_project_service_account}
	UseProjectServiceAccount interface{} `field:"optional" json:"useProjectServiceAccount" yaml:"useProjectServiceAccount"`
}

