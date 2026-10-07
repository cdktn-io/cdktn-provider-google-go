// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoType struct {
	// google_drive_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#google_drive_label DataLossPreventionContentPolicy#google_drive_label}
	GoogleDriveLabel *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabel `field:"optional" json:"googleDriveLabel" yaml:"googleDriveLabel"`
	// sensitivity_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#sensitivity_label DataLossPreventionContentPolicy#sensitivity_label}
	SensitivityLabel *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeSensitivityLabel `field:"optional" json:"sensitivityLabel" yaml:"sensitivityLabel"`
}

