// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupdrautoprotectionpolicy


type BackupDrAutoProtectionPolicyCriteria struct {
	// matching_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/backup_dr_auto_protection_policy#matching_conditions BackupDrAutoProtectionPolicy#matching_conditions}
	MatchingConditions interface{} `field:"required" json:"matchingConditions" yaml:"matchingConditions"`
}

