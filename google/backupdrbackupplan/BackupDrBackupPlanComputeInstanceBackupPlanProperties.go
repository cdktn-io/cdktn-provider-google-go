// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupdrbackupplan


type BackupDrBackupPlanComputeInstanceBackupPlanProperties struct {
	// Indicates whether to perform a guest flush operation before taking a compute instance backup.
	//
	// When set to true, the system will attempt
	// to ensure application-consistent backups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.33.0/docs/resources/backup_dr_backup_plan#guest_flush BackupDrBackupPlan#guest_flush}
	GuestFlush interface{} `field:"required" json:"guestFlush" yaml:"guestFlush"`
}

