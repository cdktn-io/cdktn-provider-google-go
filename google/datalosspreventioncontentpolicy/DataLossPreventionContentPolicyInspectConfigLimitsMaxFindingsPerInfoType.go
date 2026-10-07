// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyInspectConfigLimitsMaxFindingsPerInfoType struct {
	// Max findings limit for the given infoType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#max_findings DataLossPreventionContentPolicy#max_findings}
	MaxFindings *float64 `field:"required" json:"maxFindings" yaml:"maxFindings"`
	// info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#info_type DataLossPreventionContentPolicy#info_type}
	InfoType *DataLossPreventionContentPolicyInspectConfigLimitsMaxFindingsPerInfoTypeInfoType `field:"optional" json:"infoType" yaml:"infoType"`
}

