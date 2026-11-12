// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupdrautoprotectionpolicy


type BackupDrAutoProtectionPolicyCriteriaMatchingConditions struct {
	// label_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/backup_dr_auto_protection_policy#label_condition BackupDrAutoProtectionPolicy#label_condition}
	LabelCondition *BackupDrAutoProtectionPolicyCriteriaMatchingConditionsLabelCondition `field:"optional" json:"labelCondition" yaml:"labelCondition"`
}

