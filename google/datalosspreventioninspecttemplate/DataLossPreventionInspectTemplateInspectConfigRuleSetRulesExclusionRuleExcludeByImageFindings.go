// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioninspecttemplate


type DataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindings struct {
	// info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_loss_prevention_inspect_template#info_types DataLossPreventionInspectTemplate#info_types}
	InfoTypes interface{} `field:"required" json:"infoTypes" yaml:"infoTypes"`
	// image_containment_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_loss_prevention_inspect_template#image_containment_type DataLossPreventionInspectTemplate#image_containment_type}
	ImageContainmentType *DataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsImageContainmentType `field:"optional" json:"imageContainmentType" yaml:"imageContainmentType"`
}

