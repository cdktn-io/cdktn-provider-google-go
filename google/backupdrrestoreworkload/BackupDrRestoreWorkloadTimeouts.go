// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupdrrestoreworkload


type BackupDrRestoreWorkloadTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/backup_dr_restore_workload#create BackupDrRestoreWorkload#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/backup_dr_restore_workload#delete BackupDrRestoreWorkload#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

