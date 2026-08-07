// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupdrrestoreworkload


type BackupDrRestoreWorkloadComputeInstanceTargetEnvironment struct {
	// Required. Target project for the Compute Engine instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/backup_dr_restore_workload#project BackupDrRestoreWorkload#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Required. The zone of the Compute Engine instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/backup_dr_restore_workload#zone BackupDrRestoreWorkload#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// If true, use the BackupDR P4SA credentials for same-project restores. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/backup_dr_restore_workload#use_project_service_account BackupDrRestoreWorkload#use_project_service_account}
	UseProjectServiceAccount interface{} `field:"optional" json:"useProjectServiceAccount" yaml:"useProjectServiceAccount"`
}

