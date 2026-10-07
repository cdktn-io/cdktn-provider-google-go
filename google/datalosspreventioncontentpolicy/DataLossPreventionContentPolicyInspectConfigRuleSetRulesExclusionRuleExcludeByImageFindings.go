// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindings struct {
	// info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#info_types DataLossPreventionContentPolicy#info_types}
	InfoTypes interface{} `field:"required" json:"infoTypes" yaml:"infoTypes"`
	// image_containment_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#image_containment_type DataLossPreventionContentPolicy#image_containment_type}
	ImageContainmentType *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsImageContainmentType `field:"optional" json:"imageContainmentType" yaml:"imageContainmentType"`
}

